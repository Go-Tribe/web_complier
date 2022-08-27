package router

import (
	"web_complier/internal/appserver/controller"

	"github.com/gin-gonic/gin"
)

// setApiRoute 路由
func setApiRoute(r *gin.Engine) {
	complierV1 := r.Group("/api/v1")
	{
		complier := new(controller.ComplierController)
		complierV1.GET("/hello-world", complier.HelloWorld)
		complierV1.GET("/code", complier.GetCode)
		complierV1.POST("/run", complier.Run)
		complierV1.POST("/share", complier.Share)
	}
}
