package router

import (
	"web_complier/internal/appserver/controller"

	"github.com/gin-gonic/gin"
)

func setApiRoute(r *gin.Engine) {

	hellov1 := r.Group("/api/v1/test")
	{
		hello := new(controller.HelloController)
		hellov1.GET("/hello-world", hello.HelloWorld)
	}
}
