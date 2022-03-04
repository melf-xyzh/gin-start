/**
 * @Time    :2022/2/28 16:22
 * @Author  :ZhangXiaoyu
 */

package httpapi

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func Get(url string) (err error) {
	client := http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil
	}
	// 添加请求头
	req.Header.Add("Content-type", "application/json;charset=utf-8")
	// 添加cookie
	cookie := &http.Cookie{
		Name:  "aaa",
		Value: "aaa-value",
	}
	req.AddCookie(cookie)
	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		log.Println("err")
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("err")
	}
	fmt.Println(string(b))
	return
}
