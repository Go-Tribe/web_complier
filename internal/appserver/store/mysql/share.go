package mysql

// share 表
type Share struct {
	BaseModel
	GID  string `gorm:"column:gid;size:30;unique;" json:"gid"` // 全局唯一 ID
	Code string `gorm:"size:2000" json:"code"`                 // 代码
	Type string `gorm:"size:20" json:"type"`                   // 资源类型
}

func (Share) TableName() string {
	//实现TableName接口，以达到结构体和表对应，如果不实现该接口，并未设置全局表名禁用复数，gorm会自动扩展表名为articles（结构体+s）
	return "share"
}
