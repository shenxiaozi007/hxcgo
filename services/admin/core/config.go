package core

type Config struct {
	ServerAddr string
	DB         []*DBConfig
	Redis      []*RedisConfig
}

type DBConfig struct {
	Alias       string
	Driver      string
	Username    string
	Password    string
	Host        string
	Port        int
	Database    string
	TablePrefix string
}

type RedisConfig struct {
	Alias    string
	Host     string
	Port     int
	Password string
	Database int
}
