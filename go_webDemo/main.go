package main

import (
	"net/http"
	"fmt"
	"time"
)



func main()  {
	test1()
}

func test2()  {
	server := http.Server{
		Addr:":8080",
		ReadTimeout:time.Second * 15,
	}
	
	server.ListenAndServe()
}


func test1()  {
	http.HandleFunc("/hello" , sayHello)
	err := http.ListenAndServe(":8080",nil)
	if err != nil {
		fmt.Println(err)
	}
}

func sayHello(w http.ResponseWriter , r *http.Request)  {
	r.ParseForm() // 解析参数
	fmt.Println(r.Form)
	fmt.Println("r.URL.Path - " , r.URL.Path)
	fmt.Println("r.URL.Scheme - " , r.URL.Scheme)
	fmt.Println("r.URL.Host - " , r.URL.Host)
	
	for k , v := range r.Form {
		fmt.Println("key - " , k)
		fmt.Println("value - " , v)
	}
	fmt.Println("token " ,r.Header.Get("token"))
	
	by := make([]byte , 4096)
	r.Body.Read(by)
	fmt.Println("body " , string(by))
	r.Body.Close()
	
	fmt.Println()
	
	for k ,v := range r.Header {
		fmt.Println("header key - " ,k)
		fmt.Println("header v - " ,v)
		
	}
	
	w.Write([]byte("这是一个hello world! 服务器"))
	
}