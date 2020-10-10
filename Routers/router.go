package Routers

import (
	"github.com/gin-gonic/gin"

	. "icxl/Services"
)

func RegisterRoutes(engine *gin.Engine)  {
	v1 := engine.Group("/v1")
	{
		v1.GET("/Students", GetStudents)
		v1.POST("/Students", PostStudents)
		v1.PUT("/Students", PutStudents)
		v1.DELETE("/Students", DeleteStudents)
	}
}
