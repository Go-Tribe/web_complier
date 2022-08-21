package router

import (
	"io/ioutil"
	"net/http"
	c "web_complier/configs"
	"web_complier/internal/appserver/middleware"
	"web_complier/internal/pkg/response"
	"web_complier/internal/pkg/response/code"

	"github.com/gin-gonic/gin"
)

func SetRouters() *gin.Engine {
	var r *gin.Engine

	if c.Config.Debug == false {
		// 生产模式
		r = ReleaseRouter()
		r.Use(
			middleware.RequestHandler(),
			middleware.CustomLogger(),
			middleware.CustomRecovery(),
			middleware.CorsHandler(),
		)
	} else {
		// 开发调试模式
		r = gin.New()
		r.Use(
			middleware.RequestHandler(),
			gin.Logger(),
			middleware.CustomRecovery(),
			middleware.CorsHandler(),
		)
	}

	// ping
	r.GET("/ping", func(c *gin.Context) {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"message": "pong!",
		})
	})

	// 设置 API 路由
	setApiRoute(r)

	r.NoRoute(func(c *gin.Context) {
		response.Resp().SetHttpCode(http.StatusNotFound).FailCode(c, code.ServerError, "资源不存在")
	})

	return r
}

// ReleaseRouter 生产模式使用官方建议设置为 release 模式
func ReleaseRouter() *gin.Engine {
	// 切换到生产模式
	gin.SetMode(gin.ReleaseMode)
	// 禁用 gin 输出接口访问日志
	gin.DefaultWriter = ioutil.Discard

	engine := gin.New()

	return engine
}
