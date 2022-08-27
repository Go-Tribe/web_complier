package validator

// RequestExecParams 编译器参数校验
type RequestExecParams struct {
	Code  string `form:"code" json:"code" binding:"required,min=10"`
	Input string `form:"input" json:"input"`
	Lang  string `form:"lang" json:"lang" binding:"required"`
}

func ExecPost() *RequestExecParams {
	return &RequestExecParams{}
}
