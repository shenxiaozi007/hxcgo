package core

type Config struct {
	ServerAddr string                   `json:"server_addr"`
	Redis      []*RedisConfig           `json:"redis"`
	Wechat     map[string]*WechatConfig `json:"wechat"`
	DB         []*DBConfig              `json:"db"`
}

type WechatConfig struct {
	AppID  string `json:"app_id"`
	Secret string `json:"secret"`
}

type DBConfig struct {
	Driver      string `json:"driver"`
	Alias       string `json:"alias"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Host        string `json:"host"`
	Port        int    `json:"port"`
	Database    string `json:"database"`
	TablePrefix string `json:"table_prefix"`
}

type RedisConfig struct {
	Alias    string `json:"alias"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Password string `json:"password"`
	Database int    `json:"database"`
}
