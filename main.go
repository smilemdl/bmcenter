package main

import (
	"fmt"
	"net/http"
)

func main() {
	// 注册方法路由
	http.HandleFunc("/", LoginHandler) // todo
	http.HandleFunc("/hello", HelloHandler)
	http.HandleFunc("/login", LoginHandler)
	http.HandleFunc("/register", RegisterHandler)

	// 启动服务器
	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
