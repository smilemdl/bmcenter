package login

import (
	"net/http"
)

// RegisterHandler 处理用户注册请求
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	// 设置响应头
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	
	// 简单的注册页面响应
	w.Write([]byte("<html><body><h1>Register Page</h1><form method='POST' action='/register'><input type='text' name='username' placeholder='Username'/><br/><input type='password' name='password' placeholder='Password'/><br/><input type='submit' value='Register'/></form></body></html>"))
	// 这里可以添加处理注册逻辑的代码
	// 判断请求方法
	if r.Method == http.MethodPost {
		// 处理注册逻辑，例如保存用户信息到数据库
		username := r.FormValue("username")
		password := r.FormValue("password")
}