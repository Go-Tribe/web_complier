package application

type RedisConfig struct {
	Enable   bool   `yaml:"enable"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Password string `yaml:"password"`
	Database int    `yaml:"database"`
}

var Redis = RedisConfig{
	Enable:   false,
	Host:     "127.0.0.1",
	Password: "root1234",
	Port:     "6379",
	Database: 0,
}
