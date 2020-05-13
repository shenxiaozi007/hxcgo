package core

type Config struct {
	ServerAddr  string         `json:"server_addr"`
	Node        int            `json:"node"`
	ResourceDir string         `json:"resource_dir"`
	AppName     string         `json:"app_name"`
	ImageDir    string         `json:"image_dir"`
	ImageHost   string         `json:"image_host"`
	Redis       []*RedisConfig `json:"redis"`
	Session     *SessionConfig `json:"session"`
	RPC         []*RPCConfig   `json:"rpc"`
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

type SessionConfig struct {
	Driver   string `json:"driver"` //redis
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Password string `json:"password"`
	KeyPairs string `json:"key_pairs"`
}

type RPCConfig struct {
	ServiceName string `json:"service_name"`
	Addr        string `json:"addr"`
}
