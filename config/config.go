package config

type Config struct {
	APIKey       string
	APISecret    string
	CustomerCode string // 客户编码 月结帐号
	Version      string // 接口版本号
	Debug        bool   // 是否启用调试模式
	Sandbox      bool   // 是否为沙箱环境
	Timeout      int    // HTTP 超时设定（单位：秒）
}

func LoadConfig() *Config {
	return &Config{
		APIKey:       "0f8ea10d-9804-427a-adf4-63c8ee5e6189",
		APISecret:    "8b4eb91a-1651-4c06-8f5f-5694d92edf37",
		CustomerCode: "7310969",
		Version:      "1.1.0",
		Debug:        true,
		Sandbox:      false,
		Timeout:      360,
	}
}

// 官方文档 https://open.4px.com/v2
