package main

import (
	"github.com/gin-gonic/gin"

	. "icxl/Routers"
)
/*
1.注册路由
2.对一个表增删改查
3.docker部署
*/
func main() {
	r := gin.Default()

	RegisterRoutes(r)

	// Listen and serve on 0.0.0.0:8080
	r.Run()
}
