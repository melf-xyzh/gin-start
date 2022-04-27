/**
 * @Time    :2022/4/22 17:10
 * @Author  :ZhangXiaoyu
 */

package middleware

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/melf-xyzh/gin-start/global"
	"github.com/mssola/user_agent"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
	"time"
)

func autoCreateTables() {
	if global.DB != nil {
		err := global.DB.AutoMigrate(
			Record{},
		)
		if err != nil {
			log.Println(err)
		}
	}
}

// AccessRecord 访问记录中间件
func AccessRecord() gin.HandlerFunc {
	autoCreateTables()
	return func(c *gin.Context) {
		userAgent := c.Request.UserAgent()
		ua := user_agent.New(c.Request.UserAgent())
		rex := regexp.MustCompile(`\(([^)]+)\)`)
		params := rex.FindAllStringSubmatch(userAgent, -1)

		param := strings.Replace(params[0][0], ")", "", 1)
		uaInfo := strings.Split(param, ";")

		engineName, engineVersion := ua.Engine()
		browserName, browserVersion := ua.Browser()
		now := time.Now()
		c.Next()

		c.ContentType()

		body, _ := ioutil.ReadAll(c.Request.Body)
		data := string(body)
		data = strings.Replace(data, "\n", "", -1)
		dataLs := strings.Split(data, "\r")

		// Query
		query := c.Request.URL.RawQuery

		// Body
		if len(dataLs) > 0 {
			var build strings.Builder
			for i, v := range dataLs {
				if strings.Contains(v, "Content-Disposition") {
					build.WriteString(v)
					if !strings.Contains(v, "filename") && len(dataLs) >= i+2 {
						build.WriteString("; value=")
						build.WriteString(dataLs[i+2])
					}
					build.WriteString("\n")
				}
			}
			data = build.String()
		}

		// Headers
		header := c.Request.Header
		var headerStr string
		if len(header) > 0 {
			var build strings.Builder
			for k, v := range header {
				build.WriteString(k)
				build.WriteString(":")
				for i, v0 := range v {
					build.WriteString(v0)
					if i != len(v)-1 {
						build.WriteString(",")
					}
				}
				build.WriteString(";\n")
			}
			headerStr = build.String()
		}

		writer := ResponseWriter{ResponseWriter: c.Writer, body: &bytes.Buffer{}}
		c.Writer = writer
		fmt.Println(writer.body.String())

		record := Record{
			Ip:             c.ClientIP(),
			Method:         c.Request.Method,
			Path:           c.Request.URL.Path,
			Status:         c.Writer.Status(),
			Latency:        time.Now().Sub(now).Milliseconds(),
			UserAgent:      userAgent,
			Error:          c.Errors.ByType(gin.ErrorTypePrivate).String(),
			Body:           data,
			Query:          query,
			Header:         headerStr,
			Response:       writer.body.String(),
			Location:       "",
			Province:       "",
			City:           "",
			District:       "",
			Isp:            "",
			Platform:       ua.Platform(),
			OS:             ua.OS(),
			Engine:         engineName + "/" + engineVersion,
			BrowserName:    browserName,
			BrowserVersion: browserVersion,
			Brand:          "",
			ProductModel:   strings.TrimSpace(uaInfo[2]),
			ContentType:    c.ContentType(),
		}
		record.ID = global.CreateId()
		record.CreateTime = global.CreateTime()
		global.DB.Create(&record)
	}
}

type ResponseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

type Record struct {
	global.Model
	Ip             string `json:"ip"                     form:"ip"                   gorm:"column:ip;comment:ip;type:varchar(40);"`
	Method         string `json:"method"                 form:"method"               gorm:"column:method;comment:请求方法;type:varchar(10);"`
	Path           string `json:"path"                   form:"path"                 gorm:"column:path;comment:请求路径;type:text;"`
	Status         int    `json:"status"                 form:"status"               gorm:"column:status;comment:状态码;type:varchar(4);"`
	Latency        int64  `json:"latency"                form:"latency"              gorm:"column:latency;comment:延迟;type:varchar(5);"`
	UserAgent      string `json:"userAgent"              form:"userAgent"            gorm:"column:user_agent;comment:User-Agent;type:text;"`
	Error          string `json:"error"                  form:"error"                gorm:"column:error;comment:错误;type:varchar(255);"`
	Body           string `json:"body"                   form:"body"                 gorm:"column:body;comment:Body;type:text;"`
	Query          string `json:"query"                  form:"query"                gorm:"column:query;comment:Query;type:text;"`
	Header         string `json:"header"                 form:"header"               gorm:"column:header;comment:Header;type:text;"`
	Response       string `json:"response"               form:"response"             gorm:"column:response;comment:相应;type:text;"`
	Location       string `json:"location"               form:"location"             gorm:"column:location;comment:位置;type:varchar(30);"`
	Province       string `json:"province"               form:"province"             gorm:"column:province;comment:省;type:varchar(20);"`
	City           string `json:"city"                   form:"city"                 gorm:"column:city;comment:市;type:varchar(20);"`
	District       string `json:"district"               form:"district"             gorm:"column:district;comment:区;type:varchar(20);"`
	Isp            string `json:"isp"                    form:"isp"                  gorm:"column:isp;comment:运营商;type:varchar(20);"`
	Platform       string `json:"platform"               form:"platform"             gorm:"column:platform;comment:平台;type:varchar(50);"`
	OS             string `json:"os"                     form:"os"                   gorm:"column:os;comment:系统;type:varchar(50);"`
	Engine         string `json:"engine"                 form:"engine"               gorm:"column:engine;comment:浏览器引擎;type:varchar(100);"`
	BrowserName    string `json:"browserName"            form:"browserName"          gorm:"column:browser_name;comment:浏览器;type:varchar(200);"`
	BrowserVersion string `json:"browserVersion"         form:"browserVersion"       gorm:"column:browser_version;comment:浏览器版本;type:varchar(50);"`
	Brand          string `json:"brand"                  form:"brand"                gorm:"column:brand;comment:品牌;type:varchar(255);"`
	ProductModel   string `json:"productModel"           form:"productModel"         gorm:"column:product_model;comment:型号;type:varchar(255);"`
	ContentType    string `json:"contentType"            form:"contentType"          gorm:"column:content_type;comment:内容类型;type:varchar(255);"`
}

// TableName 自定义表名
func (Record) TableName() string {
	return "record"
}
