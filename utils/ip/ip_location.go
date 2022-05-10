/**
 * @Time    :2022/4/28 18:11
 * @Author  :ZhangXiaoyu
 */

package ip

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"github.com/melf-xyzh/gin-start/global"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	// TIME_OUT 第三方接口请求超时时间
	TIME_OUT = 3 * time.Second
)

// amapSig
/**
 *  @Description: 按高德要求对算法进行签名
 *  @param params
 *  @param sig
 *  @return sn
 */
func amapSig(params url.Values, sig string) (sn string) {
	// 对key进行排序
	keys := make([]string, len(params), len(params))
	i := 0
	for k, _ := range params {
		keys[i] = k
		i += 1
	}
	sort.Strings(keys)

	// 按顺序对参数进行拼接
	var build strings.Builder
	for j, key := range keys {
		if j == 0 {
			build.WriteString(key + "=" + params.Get(key))
		} else {
			build.WriteString("&" + key + "=" + params.Get(key))
		}
	}
	// 拼接数字秘钥
	build.WriteString(sig)
	paramUrl := build.String()

	// md5签名
	s := md5.New()
	s.Write([]byte(paramUrl))
	return hex.EncodeToString(s.Sum(nil))
}

// 参考文档：https://lbs.amap.com/api/webservice/guide/api/ipconfig

// GetGPSByIpAmap
/**
 *  @Description: 获取IP对应得GPS坐标(内网地址则匹配当前地址所对应得公网地址)
 *  @param ip ip地址
 *  @return data
 *  @return err 错误信息
 */
func GetGPSByIpAmap(ip string) (data map[string]interface{}, err error) {
	// 高德Key
	amapKey := strings.TrimSpace(global.V.GetString("amap.amap-key"))
	if amapKey == "" {
		return data, errors.New("高德Key未配置")
	}
	// 高德数字签名
	sig := strings.TrimSpace(global.V.GetString("amap.sig"))

	params := url.Values{}
	var Url *url.URL
	var resp *http.Response
	Url, err = url.Parse(strings.TrimSpace(global.V.GetString("amap.url")))
	if err != nil {
		return
	}
	params.Set("key", amapKey)
	params.Set("type", "4")
	params.Set("ip", ip)
	// 数字签名
	if sig != "" {
		sig = amapSig(params, sig)
		params.Set("sig", sig)
	}
	// 如果参数中有中文参数,这个方法会进行URLEncode
	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	// 设置超时时间
	client := http.Client{
		Timeout: TIME_OUT,
	}
	// 发送Get请求
	resp, err = client.Get(urlPath)
	if err != nil {
		return data, err
	}
	if resp.StatusCode != 200 {
		return data, errors.New("StatusCode：" + strconv.Itoa(resp.StatusCode))
	}
	// 关闭流
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("关闭流失败:", err.Error())
			return
		}
	}(resp.Body)
	body, _ := ioutil.ReadAll(resp.Body)
	// 解析json结果
	err = json.Unmarshal(body, &data)
	if err != nil {
		return data, errors.New("解析数据失败")
	}
	return
}
