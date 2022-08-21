package application

type ServerConfig struct {
	Host string `yaml:"host"`
	Port uint16 `yaml:"port"`
}

var Server = ServerConfig{
	Host: "127.0.0.1",
	Port: 9999,
}
