package code

import (
	c "web_complier/configs"
)

const (
	SUCCESS            = 0
	FAILURE            = 1
	NotFound           = 404
	ParamBindError     = 10000
	ServerError        = 10101
	TooManyRequests    = 10102
	AuthorizationError = 10103
	RBACError          = 10104
	CodeLenError       = 2001
	CodeTypeError      = 2002
	CodeShareError     = 2003
	CodeGetInfoError   = 2004
)

func Text(code int) (str string) {
	lang := c.Config.Language

	var ok bool
	switch lang {
	case "zh_CN":
		str, ok = zhCNText[code]
		break
	case "en":
		str, ok = enUSText[code]
		break
	}
	if !ok {
		return "unknown error"
	}
	return
}
