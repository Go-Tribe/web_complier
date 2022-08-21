package application

type DivisionTime struct {
	MaxAge       int ` yaml:"max_age"`
	RotationTime int ` yaml:"rotation_time"`
}

type DivisionSize struct {
	MaxSize    int  ` yaml:"max_size"`
	MaxBackups int  ` yaml:"max_backups"`
	MaxAge     int  ` yaml:"max_age"`
	Compress   bool ` yaml:"compress"`
}

type LoggerConfig struct {
	DefaultDivision string       `yaml:"default_division"`
	Filename        string       `yaml:"file_name"`
	DivisionTime    DivisionTime `yaml:"division_time"`
	DivisionSize    DivisionSize `yaml:"division_size"`
}

var Logger = LoggerConfig{
	// time 按时间切割，默认一天, size 按文件大小切割
	DefaultDivision: "time",
	Filename:        "sys.log",
	DivisionTime: DivisionTime{
		MaxAge:       15,
		RotationTime: 24,
	},
	DivisionSize: DivisionSize{
		MaxSize:    2,
		MaxBackups: 2,
		MaxAge:     15,
		Compress:   false,
	},
}
