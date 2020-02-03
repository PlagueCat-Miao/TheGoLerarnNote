package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func hellofunction(w http.ResponseWriter, r *http.Request) {
	fmt.Println("通过GET访问 http://.../hello 确实地调用了其对应函数hellofunction Nya~")
//	fmt.Println("请求消息体：" ,r.Body)
//	fmt.Println("请求的方法：" ,r.Method)
//	fmt.Println("请求的头部：" ,r.Header)
	fmt.Fprint(w, "╭(￣▽￣)╯Resp! hello meow~")
}
func postAndjson(writer http.ResponseWriter, request *http.Request) {
	// 检查是否POST请求
	if request.Method != "POST" {
		writer.WriteHeader(405)
		return
	}
	data, err := ioutil.ReadAll(request.Body)
	if err != nil {
		fmt.Println("post get body fail")
	}
	fmt.Println("POST数据content-type类型：" ,request.Header["Content-Type"])
	fmt.Println("POST消息体：" ,string(data))
	fmt.Println("POST的Body与content-type无关")

	fmt.Fprint(writer, "(`･ω･´)ゞ Resp! i get post meow~")
}

//httpsever
func SeverHttpStart() {
	http.HandleFunc("/hello", hellofunction)
	http.HandleFunc("/post", postAndjson)

	err := http.ListenAndServe("127.0.0.1:8888", nil)
	if err != nil {
		fmt.Println("http listen failed")
	}

}
//httpclient
func main() {
	go SeverHttpStart();
	fmt.Println("sever set~")
	fmt.Println("get test")
	res, err := http.Get("http://127.0.0.1:8888/hello")

	if err != nil {
		fmt.Println("Get url failed:", err)
		return
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("get data error:", err)
		return
	}
	fmt.Println(string(data))
	//post
	fmt.Println("post test")
	res2, err := http.Post("http://127.0.0.1:8888/post"," (¦3[▓▓]",strings.NewReader("miao~"))
	data2, err := ioutil.ReadAll(res2.Body)
	fmt.Println(string(data2))
}

