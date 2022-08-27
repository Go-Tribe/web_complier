package validator

// RequestShareParams 分享&保存参数
type RequestShareParams struct {
	Code string `form:"code" json:"code" binding:"required,min=10"`
	Lang string `form:"lang" json:"lang" binding:"required"`
}

func SharePost() *RequestShareParams {
	return &RequestShareParams{}
}
