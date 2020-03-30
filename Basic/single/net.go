package main

import (
	"fmt"
	"net/url"
)

func urlTest() {
	s := "postgres://user:pass@host.com:5432/path?urlkey=thisinfo#f"
	//开始解析
	u, err := url.Parse(s)
	if err != nil {
		println(err)
	}
	//协议名
	fmt.Println(u.Scheme)
	m, _ := url.ParseQuery(u.RawQuery)
	//获取url的key-value
	fmt.Println(m["urlkey"][0])
}
func test5() {
	urlTest()
}
