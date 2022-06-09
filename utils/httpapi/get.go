/**
 * @Time    :2022/2/28 16:22
 * @Author  :ZhangXiaoyu
 */

package httpapi

import (
	"crypto/tls"
	"errors"
	datautils "github.com/melf-xyzh/gin-start/utils/data"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
	"unsafe"
)

type GetRequest struct{}

var Get = new(GetRequest)

// Get
/**
 * @Description: 标准的get请求
 * @receiver h
 * @param urlPath http接口地址
 * @return data
 * @return err
 */
func (g *GetRequest) Get(urlPath string) (data string, err error) {
	var res *http.Response
	res, err = http.Get(urlPath)
	if err != nil {
		return data, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			return
		}
	}(res.Body)
	// 读取标准输出
	body, errRead := ioutil.ReadAll(res.Body)
	if res.StatusCode != 200 || errRead != nil {
		err = errors.New("请求失败")
		return data, err
	}
	// byte数组直接转成string，优化内存
	str := (*string)(unsafe.Pointer(&body))
	return *str, nil
}

// getUrlPathWithParams
/**
 * @Description: 拼接url和params
 * @receiver h
 * @param urlPath
 * @param paramMap
 * @return urlPathNew
 * @return err
 */
func (g GetRequest) getUrlPathWithParams(urlPath string, paramMap map[string]interface{}) (urlPathNew string, err error) {
	params := url.Values{}
	var parseURL *url.URL
	parseURL, err = url.Parse(urlPath)
	if err != nil {
		return "", err
	}
	if paramMap != nil {
		for k, v := range paramMap {
			params.Set(k, datautils.ToString(v))
		}
	}
	// 如果参数中有中文参数,这个方法会进行URLEncode
	parseURL.RawQuery = params.Encode()
	urlPathNew = parseURL.String()
	return urlPathNew, nil
}

// GetWithParams
/**
 * @Description: 携带Params参数的http请求
 * @receiver h
 * @param urlPath http接口地址
 * @param paramMap 参数map
 * @return data 返回的数据
 * @return err
 */
func (g *GetRequest) GetWithParams(urlPath string, paramMap map[string]interface{}) (data string, err error) {
	urlPath, err = g.getUrlPathWithParams(urlPath, paramMap)
	if err != nil {
		return "", err
	}
	data, err = g.Get(urlPath)
	if err != nil {
		return "", err
	}
	return data, nil
}

// GetPro
/**
 * @Description: Get请求（专业版）（支持自定义header、cookie）
 * @receiver h
 * @param urlPath http接口地址
 * @param paramMap 参数map
 * @param headerMap 请求头Map
 * @param cookies cookie列表
 * @return data 返回的数据
 * @return err 错误
 */
func (g *GetRequest) GetPro(urlPath string, paramMap map[string]interface{}, headerMap map[string]interface{}, cookies []*http.Cookie) (data string, err error) {
	urlPath, err = g.getUrlPathWithParams(urlPath, paramMap)
	if err != nil {
		return "", err
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
	req, err = http.NewRequest(http.MethodGet, urlPath, nil)
	// 添加请求头
	if headerMap != nil {
		for k, v := range headerMap {
			req.Header.Add(k, datautils.ToString(v))
		}
	}

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
