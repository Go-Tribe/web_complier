package controller

import (
	"fmt"
	"web_complier/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

type HelloController struct {
}

func (h *HelloController) HelloWorld(c *gin.Context) {
	str, ok := c.GetQuery("name")
	if !ok {
		str = "web_complier"
	}

	response.Success(c, fmt.Sprintf("hello %s", str))
}
