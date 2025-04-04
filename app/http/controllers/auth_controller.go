// Package controllers 用户认证相关控制器
package controllers

import (
	"goblog/pkg/view"
	"net/http"
)

// AuthController 处理用户认证
type AuthController struct {
}

func (*AuthController) Register(w http.ResponseWriter, r *http.Request) {
	view.RenderSimple(w, view.D{}, "auth.register")
}

// DoRegister 处理注册逻辑
func (*AuthController) DoRegister(w http.ResponseWriter, r *http.Request) {
}
