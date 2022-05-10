/**
 * @Time    :2022/4/22 17:10
 * @Author  :ZhangXiaoyu
 */

package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/melf-xyzh/gin-start/global"
	"github.com/melf-xyzh/gin-start/utils/ip"
	"github.com/mssola/user_agent"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
	"time"
)

// 自动建表
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

// parseUserAgent
/**
 *  @Description: 解析userAgent
 *  @param c
 *  @param record
 */
func parseUserAgent(c *gin.Context, record *Record) {
	// 获取userAgent
	userAgent := c.Request.UserAgent()
	ua := user_agent.New(c.Request.UserAgent())
	rex := regexp.MustCompile(`\(([^)]+)\)`)
	params := rex.FindAllStringSubmatch(userAgent, -1)
	param := strings.Replace(params[0][0], ")", "", 1)
	uaInfo := strings.Split(param, ";")
	engineName, engineVersion := ua.Engine()
	browserName, browserVersion := ua.Browser()

	record.UserAgent = userAgent
	record.Platform = ua.Platform()
	record.OS = ua.OS()
	record.Engine = engineName + "/" + engineVersion
	record.BrowserName = browserName
	record.BrowserVersion = browserVersion
	record.Brand = ""
	record.ProductModel = strings.TrimSpace(uaInfo[2])
}

// parseBody
/**
 *  @Description: 获取Body
 *  @param c
 *  @param record
 */
func parseBody(c *gin.Context, record *Record) {
	// 获取Request.Body
	body, _ := ioutil.ReadAll(c.Request.Body)
	// 将其转为String
	data := string(body)
	// 替换字符串及字符串分隔
	data = strings.Replace(data, "\n", "", -1)
	dataLs := strings.Split(data, "\r")
	record.Body = data
	if len(dataLs) == 0 {
		return
	}
	// 拼接请求参数
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
	return
}

func parseIp(c *gin.Context, record *Record) {
	record.Ip = c.ClientIP()
	data, err := ip.GetGPSByIpAmap(record.Ip)
	if err != nil {
		return
	}
	if data["country"] != nil {
		record.Country = data["country"].(string)
	}
	if data["location"] != nil {
		record.Country = data["location"].(string)
	}
	if data["province"] != nil {
		record.Country = data["province"].(string)
	}
	if data["city"] != nil {
		record.Country = data["city"].(string)
	}
}

// getHeader
/**
 *  @Description: 获取header
 *  @param c
 *  @return headerStr
 */
func getHeader(c *gin.Context) (headerStr string) {
	// 获取Herader
	header := c.Request.Header
	if len(header) == 0 {
		return
	}
	// 拼接Herder
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
	return
}

// AccessRecord 访问记录中间件
func AccessRecord() gin.HandlerFunc {
	autoCreateTables()
	return func(c *gin.Context) {
		record := &Record{}
		record.Method = c.Request.Method
		record.Path = c.Request.URL.Path
		// ContentType
		record.ContentType = c.ContentType()
		// 解析userAgent
		parseUserAgent(c, record)
		// 解析Body
		parseBody(c, record)
		// IP
		parseIp(c, record)
		// 保存Query参数
		record.Query = c.Request.URL.RawQuery
		// 保存Header
		record.Header = getHeader(c)
		// ResponseWriter.Write查看响应
		writer := responseBodyWriter{
			ResponseWriter: c.Writer,
			body:           &bytes.Buffer{},
		}
		c.Writer = writer
		now := time.Now()
		c.Next()
		record.Response = writer.body.String()
		record.Error = c.Errors.ByType(gin.ErrorTypePrivate).String()
		// 延时，毫秒
		record.Latency = time.Now().Sub(now).Milliseconds()
		// 请求状态码
		record.Status = c.Writer.Status()
		// 保存基础信息
		record.ID = global.CreateId()
		record.CreateTime = global.CreateTime()
		global.DB.Create(&record)
	}
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
	Response       string `json:"response"               form:"response"             gorm:"column:response;comment:响应;type:text;"`
	Location       string `json:"location"               form:"location"             gorm:"column:location;comment:位置;type:varchar(30);"`
	Country        string `json:"country"                form:"country"              gorm:"column:country;comment:国家;type:varchar(255)"`
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

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}
