package validator

// RequestCodeParams 获取保存参数
type RequestCodeParams struct {
	GID string `form:"gid" json:"gid" binding:"required,min=5"`
}

func CodeGet() *RequestCodeParams {
	return &RequestCodeParams{}
}
