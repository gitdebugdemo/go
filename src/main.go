package main

import (
	"log"
	"net/http"
)

func main() {
	//指定路径上监听消息
	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		//服务器上打印的日志
		log.Print("ok /")
		//客户端返回的信息
		resp.Write([]byte("OK\n"))
	})
	//监听端口,错误的时候,输出错误信息
	s := ":8080"
	log.Print("开启服务地址:" + s)
	log.Fatal(http.ListenAndServe(s, nil))
}
