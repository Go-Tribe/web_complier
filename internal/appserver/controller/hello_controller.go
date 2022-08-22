package controller

import (
	"fmt"
	"web_complier/internal/pkg/docker"
	"web_complier/internal/pkg/response"
	"web_complier/internal/pkg/response/code"

	"github.com/gin-gonic/gin"
)

type ComplierController struct {
}

// RequestExecParams 编译器参数校验
type RequestExecParams struct {
	Code  string `form:"code" json:"code" binding:"required"`
	Input string `form:"input" json:"input"`
	Lang  string `form:"lang" json:"lang" binding:"required"`
}

func (h *ComplierController) HelloWorld(c *gin.Context) {
	str, ok := c.GetQuery("name")
	if !ok {
		str = "web_complier"
	}

	response.Success(c, fmt.Sprintf("hello %s", str))
}

func (h *ComplierController) Run(c *gin.Context) {
	params := RequestExecParams{}

	if err := c.ShouldBindJSON(&params); err != nil {
		response.Fail(c, code.ParamBindError, err.Error())
		return
	}

	if len(params.Code) > 1024*400 {
		response.Fail(c, code.CodeLenError, "提交的代码太长，最多允许400KB")
		return
	}

	langexists, _ := docker.LangExists(params.Lang)
	if !langexists {
		response.Fail(c, code.CodeLenError, "暂不支持该语言")
		return
	}

	tpl := docker.Run(params.Lang)
	output := docker.DockerRun(tpl.Image, params.Code, tpl.File, tpl.Cmd, tpl.Timeout, tpl.Memory)
	// 返回数据
	data := make(map[string]string)
	data["stdout"] = output
	data["stderr"] = ""

	response.Success(c, data)
}
