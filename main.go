package main

import (
	"goblog/app/http/middlewares"
	"goblog/bootstrap"
	"net/http"
)

func main() {

	bootstrap.SetupDB()
	// 初始化路由
	router := bootstrap.SetupRoute()

	http.ListenAndServe(":3000", middlewares.RemoveTrailingSlash(router))

}
