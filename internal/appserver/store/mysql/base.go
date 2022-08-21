package mysql

import (
	"database/sql/driver"
	"fmt"
	"time"
	"web_complier/core"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt formatDate     `json:"created_at"`
	UpdatedAt formatDate     `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (model BaseModel) DB() *gorm.DB {
	return DB()
}

func DB() *gorm.DB {
	return core.MysqlDB
}

type formatDate struct {
	time.Time
}

func (t formatDate) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%s\"", t.Format("2006-01-02 15:04:05"))
	return []byte(formatted), nil
}

func (t formatDate) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

func (t *formatDate) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = formatDate{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
