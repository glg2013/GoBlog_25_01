// Package bootstrap 负责应用初始化相关工作，比如初始化路由。
package bootstrap

import (
	"github.com/gorilla/mux"
	"goblog/routes"
)

func SetupRoute() *mux.Router {
	router := mux.NewRouter()
	routes.RegisterWebRoutes(router)
	return router
}
