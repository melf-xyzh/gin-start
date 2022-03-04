/**
 * @Time    :2022/3/4 16:06
 * @Author  :ZhangXiaoyu
 */

package ip

import (
	"errors"
	"fmt"
	"github.com/thinkeridea/go-extend/exnet"
	"io"
	"io/ioutil"
	"net"
	"net/http"
)

const (
	// GET_PUBLIC_IP_URL 查询公网IP的URL
	GET_PUBLIC_IP_URL = "http://myexternalip.com/raw"
)

// GetClientIP
/**
 *  @Description: 获取用户真实IP
 *  @param r
 *  @return ip
 */
func GetClientIP(r *http.Request) (ip string) {
	ip = exnet.ClientPublicIP(r)
	if ip == ""{
		ip = exnet.ClientIP(r)
	}
	return
}

// GetLocalIp
/**
 *  @Description: 查询本机内网IP
 *  @return ips ip列表
 *  @return err 错误
 */
func GetLocalIp() (ips []string, err error) {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("get ip interfaces error:", err)
		return
	}

	for _, i := range netInterfaces {
		address, errRet := i.Addrs()
		if errRet != nil {
			continue
		}

		for _, addr := range address {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
				if ip.IsGlobalUnicast() {
					ips = append(ips, ip.String())
				}
			}
		}
	}
	return
}

// GetPublicIP
/**
 *  @Description: 获取本机公网ip
 *  @return ip
 *  @return err
 */
func GetPublicIP() (ip string, err error) {
	resp, errHttp := http.Get(GET_PUBLIC_IP_URL)
	if errHttp != nil {
		return ip, errHttp
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)
	// 读取标准输出
	body, errRead := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 || errRead != nil {
		err = errors.New("请求失败")
		return ip, err
	}
	ip = string(body)
	return ip, nil
}


func IsIpAddress(ip string) (ok bool) {

	return
}

func IsIPv4(ip string) (ok string) {
	return
}

func IsIPv6(ip string) (ok string) {
	return
}
