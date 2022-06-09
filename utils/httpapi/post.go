/**
 * @Time    :2022/2/28 18:18
 * @Author  :ZhangXiaoyu
 */

package httpapi

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	datautils "github.com/melf-xyzh/gin-start/utils/data"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
	"unsafe"
)

type PostRequest struct{}

const (
	CONTENT_TYPE_JSON       = "application/json;charset=UTF-8"
	CONTENT_TYPE_URLENCODED = "application/x-www-form-urlencoded"
	CONTENT_TYPE_FROM_DATA  = "multipart/form-data"
)

var Post = new(PostRequest)

// PostJson
/**
 * @Description: POST请求（application/json）
 * @receiver p
 * @param urlPath
 * @param bodyMap
 * @return data
 * @return err
 */
func (p *PostRequest) PostJson(urlPath string, bodyMap map[string]interface{}) (data string, err error) {
	var bytesData []byte
	bytesData, err = json.Marshal(bodyMap)
	if err != nil {
		return "", err
	}
	var res *http.Response
	res, err = http.Post(urlPath, CONTENT_TYPE_JSON, bytes.NewReader(bytesData))
	if err != nil {
		return "", err
	}
	// 最后关闭res.Body文件
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			return
		}
	}(res.Body)

	// 使用ioutil.ReadAll将res.Body中的数据读取出来,并使用body接收
	var body []byte
	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	// byte数组直接转成string，优化内存
	str := (*string)(unsafe.Pointer(&body))
	return *str, nil
}

// PostFormData
/**
 * @Description: POST请求（multipart/form-data）
 * @receiver p
 * @param urlPath
 * @param bodyMap
 * @return data
 * @return err
 */
func (p *PostRequest) PostFormData(urlPath string, bodyMap map[string]interface{}) (data string, err error) {
	dataVal := url.Values{}
	// 拼接参数
	if bodyMap != nil {
		for k, v := range bodyMap {
			dataVal.Set(k, datautils.ToString(v))
		}
	}

	var res *http.Response
	res, err = http.PostForm(urlPath, dataVal)
	if err != nil {
		return
	}

	// 最后关闭res.Body文件
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			return
		}
	}(res.Body)
	var body []byte
	// 使用ioutil.ReadAll将res.Body中的数据读取出来,并使用body接收
	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	// byte数组直接转成string，优化内存
	str := (*string)(unsafe.Pointer(&body))
	return *str, nil
}

// PostUrlencoded
/**
 * @Description: POST请求（浏览器默认表单格式）（application/x-www-form-urlencoded）
 * @receiver p
 * @param urlPath http接口地址
 * @param bodyMap 参数map
 * @return data
 * @return err
 */
func (p *PostRequest) PostUrlencoded(urlPath string, bodyMap map[string]interface{}) (data string, err error) {
	dataVal := url.Values{}
	// 拼接参数
	if bodyMap != nil {
		for k, v := range bodyMap {
			dataVal.Set(k, datautils.ToString(v))
		}
	}

	var res *http.Response
	res, err = http.Post(urlPath, CONTENT_TYPE_URLENCODED, strings.NewReader(dataVal.Encode()))
	if err != nil {
		fmt.Println(err)
	}

	// 最后关闭res.Body文件
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			return
		}
	}(res.Body)
	var body []byte
	// 使用ioutil.ReadAll将res.Body中的数据读取出来,并使用body接收
	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	// byte数组直接转成string，优化内存
	str := (*string)(unsafe.Pointer(&body))
	return *str, nil
}

// PostPro
/**
 * @Description: Post请求（专业版）（支持自定义header、cookie）
 * @receiver p
 * @param urlPath http接口地址
 * @param contentType Content-Type
 * @param paramMap 参数map
 * @param headerMap 请求头Map
 * @param cookies cookie列表
 * @return data 返回的数据
 * @return err 错误
 */
func (p *PostRequest) PostPro(urlPath, contentType string, bodyMap map[string]interface{}, headerMap map[string]interface{}, cookies []*http.Cookie) (data string, err error) {
	dataVal := url.Values{}
	// 拼接参数
	if bodyMap != nil {
		for k, v := range bodyMap {
			dataVal.Set(k, datautils.ToString(v))
		}
	}

	// 解决 x509: certificate signed by unknown authority
	// 通过设置tls.Config的InsecureSkipVerify为true，client将不再对服务端的证书进行校验。
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := http.Client{
		Timeout:   time.Second * 20,
		Transport: transport,
	}
	var req *http.Request
	req, err = http.NewRequest(http.MethodPost, urlPath, strings.NewReader(dataVal.Encode()))
	// 添加请求头
	if headerMap != nil {
		for k, v := range headerMap {
			req.Header.Set(k, datautils.ToString(v))
		}
	}
	req.Header.Set("Content-Type", contentType)
	// 添加cookie
	if cookies != nil {
		for _, cookie := range cookies {
			req.AddCookie(cookie)
		}
	}
	// 发送请求
	var res *http.Response
	res, err = client.Do(req)
	if err != nil {
		return "", err
	}

	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			return
		}
	}(res.Body)

	var body []byte
	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	// byte数组直接转成string，优化内存
	str := (*string)(unsafe.Pointer(&body))
	return *str, nil
}
