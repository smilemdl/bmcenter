package main

import (
	"fmt"
	"net/http"
	"os"
)

// LoginHandler 处理 /login 路由的函数
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	Handler(w, r, "login/login.html")
}

// HelloHandler 处理 /hello 路由的函数
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

// TopHandler 处理 / 根页面路由的函数
func TopHandler(w http.ResponseWriter, r *http.Request) {
	Handler(w, r, "top.html")
}

// Handler 根据路由名调用对应的处理函数
func Handler(w http.ResponseWriter, r *http.Request, route string) {
	// 设置响应头
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// 读HTML内容（直接返回字符串）
	html, err := os.ReadFile(route)
	if err != nil {
		http.Error(w, "Could not load page", http.StatusInternalServerError)
		fmt.Println("Could not load page")
		return
	}

	// 写入Writer
	fmt.Fprintln(w, string(html))
}
