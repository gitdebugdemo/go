package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/h1", func(writer http.ResponseWriter, request *http.Request) {
		//第一步必须:解释
		request.ParseForm()
		//第二步:获取值
		rt := request.FormValue("a")
		writer.Write([]byte(rt))
	})
	//x 指定路径上监听消息 - 如果ListenAndServe指定了第2个参数,那么这个就报废了
	http.HandleFunc("/xx", func(resp http.ResponseWriter, req *http.Request) {
		//客户端返回的信息
		resp.Write([]byte("OK\n"))
	})
	//监听端口,错误的时候,输出错误信息
	s := ":8080"
	log.Print("开启服务地址:" + s)
	log.Fatal(http.ListenAndServe(s, mux))
}
