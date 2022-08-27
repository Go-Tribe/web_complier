package controller

import (
	"fmt"
	"web_complier/internal/appserver/service"
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
		response.Fail(c, code.CodeTypeError, "暂不支持该语言")
		return
	}

	tpl := docker.Run(execPost.Lang)
	output := docker.DockerRun(tpl.Image, execPost.Code, tpl.File, tpl.Cmd, tpl.Timeout, tpl.Memory)
	// 返回数据
	response.Success(c, &response.RunResponse{Stdout: output})
}

func (h *ComplierController) Share(c *gin.Context) {
	sharePost := validator.SharePost()
	if err := validator.CheckPostParams(c, &sharePost); err != nil {
		return
	}

	if len(sharePost.Code) > 1024*400 {
		response.Fail(c, code.CodeLenError, "提交的代码太长，最多允许400KB")
		return
	}

	gid, err := service.ComplierService.Create(sharePost.Code, sharePost.Lang)
	if err != nil {
		response.Fail(c, code.CodeShareError, "分享失败")
		return
	}
	response.Success(c, &response.ShareResponse{URL: gid})
}

func (h *ComplierController) GetCode(c *gin.Context) {
	codeGet := validator.CodeGet()
	if err := validator.CheckQueryParams(c, &codeGet); err != nil {
		return
	}
	sharRes, err := service.ComplierService.FindOne(codeGet.GID)
	if err != nil {
		response.Fail(c, code.CodeGetInfoError, "获取信息失败")
		return
	}
	response.Success(c, sharRes)
}
