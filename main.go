package main

import (
	"fmt"
	"net/http"
	"strings"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>Hello,欢迎来到 goblog</h1>")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "<h1>请求页面未找到：(</h1>"+"<p>如有疑惑，请联系我们！")
	}
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "此博客是用以记录变成笔记，如您有反馈或建议，请联系 "+"<a href=\"mailto:fengniancong@163.com\"> fengniancong@163.com</a> ")
}

func main() {
	router := http.NewServeMux()

	router.HandleFunc("/", handlerFunc)
	router.HandleFunc("/about", aboutHandler)

	// 文章详情
	router.HandleFunc("/articles/", func(w http.ResponseWriter, r *http.Request) {
		id := strings.SplitN(r.URL.Path, "/", 3)[2]
		fmt.Fprint(w, "文章详情，文章 ID 为："+id)
	})

	http.ListenAndServe(":3000", router)
}
