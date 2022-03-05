/**
 * @Time    :2022/3/4 16:06
 * @Author  :ZhangXiaoyu
 */

package ip

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
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

var (
	v *validator.Validate
)

// GetClientIP
/**
 *  @Description: 获取用户真实IP
 *  @param r
 *  @return ip
 */
func GetClientIP(r *http.Request) (ip string) {
	ip = exnet.ClientPublicIP(r)
	if ip == "" {
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

// IsIpAddress
/**
 * @Description: 判断是否为合法的ip地址
 * @param ip ip地址
 * @return ok 是否合法
 */
func IsIpAddress(ip string) (ok bool) {
	if v == nil {
		v = validator.New()
	}
	errs := v.Var(ip, "ip")
	if errs != nil {
		return false
	} else {
		return true
	}
}

// IsIPv4
/**
 * @Description: 判断是否为合法的ipv4地址
 * @param ip ip地址
 * @return ok 是否合法
 */
func IsIPv4(ip string) (ok bool) {
	if v == nil {
		v = validator.New()
	}
	errs := v.Var(ip, "ipv4")
	if errs != nil {
		return false
	} else {
		return true
	}
}

// IsIPv6
/**
 * @Description: 判断是否为合法的ipv6地址
 * @param ip ip地址
 * @return ok 是否合法
 */
func IsIPv6(ip string) (ok bool) {
	if v == nil {
		v = validator.New()
	}
	errs := v.Var(ip, "ipv6")
	if errs != nil {
		return false
	} else {
		return true
	}
}
