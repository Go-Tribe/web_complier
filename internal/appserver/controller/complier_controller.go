package controller

import (
	"fmt"
	"web_complier/internal/pkg/docker"
	"web_complier/internal/pkg/response"
	"web_complier/internal/pkg/response/code"
	"web_complier/internal/pkg/validator"

	"github.com/gin-gonic/gin"
)

type ComplierController struct {
}

func (h *ComplierController) HelloWorld(c *gin.Context) {
	str, ok := c.GetQuery("name")
	if !ok {
		str = "web_complier"
	}

	response.Success(c, fmt.Sprintf("hello %s", str))
}

func (h *ComplierController) Run(c *gin.Context) {
	execPost := validator.ExecPost()
	if err := validator.CheckPostParams(c, &execPost); err != nil {
		return
	}

	if len(execPost.Code) > 1024*400 {
		response.Fail(c, code.CodeLenError, "提交的代码太长，最多允许400KB")
		return
	}

	langexists, _ := docker.LangExists(execPost.Lang)
	if !langexists {
		response.Fail(c, code.CodeLenError, "暂不支持该语言")
		return
	}

	tpl := docker.Run(execPost.Lang)
	output := docker.DockerRun(tpl.Image, execPost.Code, tpl.File, tpl.Cmd, tpl.Timeout, tpl.Memory)
	// 返回数据
	data := make(map[string]string)
	data["stdout"] = output
	data["stderr"] = ""

	response.Success(c, data)
}
