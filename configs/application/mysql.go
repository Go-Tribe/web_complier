package application

import (
	"time"
)

type MysqlConfig struct {
	Enable       bool          `yaml:"enable"`
	Host         string        `yaml:"host"`
	Username     string        `yaml:"username"`
	Password     string        `yaml:"password"`
	Port         uint16        `yaml:"port"`
	Database     string        `yaml:"database"`
	Charset      string        `yaml:"charset"`
	TablePrefix  string        `yaml:"table_prefix"`
	MaxIdleConns int           `yaml:"max_idle_conns"`
	MaxOpenConns int           `yaml:"max_open_conns"`
	MaxLifetime  time.Duration `yaml:"max_lifetime"`
	LogLevel     int           `yaml:"log_level"`
	PrintSQL     bool          `yaml:"print_sql"`
}

var Mysql = MysqlConfig{
	Enable:       false,
	Host:         "127.0.0.1",
	Username:     "root",
	Password:     "root1234",
	Port:         3306,
	Database:     "test",
	Charset:      "utf8mb4",
	TablePrefix:  "",
	MaxIdleConns: 10,
	MaxOpenConns: 100,
	MaxLifetime:  time.Hour,
	LogLevel:     4,
	PrintSQL:     true,
}
