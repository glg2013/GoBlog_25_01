// Package route 路由相关
package route

import (
	"github.com/gorilla/mux"
	"net/http"
)

var Router *mux.Router

func Initialize() {
	Router = mux.NewRouter()
}

// Name2URL 通过路由名称来获取 URL
func Name2URL(routeName string, pairs ...string) string {
	url, err := Router.Get(routeName).URL(pairs...)
	if err != nil {
		//checkError(err)
		return ""
	}

	return url.String()
}

// GetRouteVariable 获取 URI 路由参数
func GetRouteVariable(parameterName string, r *http.Request) string {
	vars := mux.Vars(r)
	return vars[parameterName]
}
