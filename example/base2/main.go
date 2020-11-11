package main

import (
	"fmt"
	"log"
	"net/http"
)


type Engine struct {}

/**
* 实现 http 中的接口 Handler中的 ServeHTTP方法
* 把所有的请求自定义函数处理
*/
func (engine * Engine) ServeHTTP(w http.ResponseWriter, r * http.Request) {
	switch r.URL.Path {
	case "/":
		fmt.Fprintf(w, "URL.PATH = %q \n", r.URL.Path)
	case "/hello":
		for k, v := range r.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	default:
		fmt.Fprintf(w, "404 NOT FOUND:%s\n", r.URL.Path)
	}

}

func main() {
	engine := new(Engine)
	log.Fatal(http.ListenAndServe(":9998", engine))

}