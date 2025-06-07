package config

type Config struct {
	APIKey    string
	APISecret string
	Debug     bool // 是否启用调试模式
	Sandbox   bool // 是否为沙箱环境
	Timeout   int  // HTTP 超时设定（单位：秒）
}

func LoadConfig() *Config {
	return &Config{
		APIKey:    "ITC0893791",
		APISecret: "axzc2utvPbfc9UbJDOh+7w==",
		Debug:     true,
		Sandbox:   true,
		Timeout:   360,
	}
}
