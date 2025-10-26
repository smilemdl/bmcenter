package main

import (
	"fmt"
	"net/http"
	"os"
)

// HelloHandler 处理 /hello 路由的函数
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

// TopHandler 处理 / 根页面路由的函数
func TopHandler(w http.ResponseWriter, r *http.Request) {
	// 设置响应头为 HTML 类型，确保浏览器正确渲染
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// 按钮页面的 HTML 内容（直接返回字符串）
	html, err := os.ReadFile("top.html")
	if err != nil {
		http.Error(w, "Could not load page", http.StatusInternalServerError)
		fmt.Println("Could not load page")
		return
	}

	// 写入Writer
	fmt.Fprintln(w, string(html))
}
