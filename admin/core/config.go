package core

type Config struct {
	ServerAddr  string
	Node        int
	UploadDir   string
	TemplateDir string
	ResourceDir string
	ImageHost   string
	DB          []*DBConfig
	Redis       []*RedisConfig
	Session     *SessionConfig
	RPC         []*RPCConfig
}

type DBConfig struct {
	Driver      string
	Alias       string
	Username    string
	Password    string
	Host        string
	Port        int
	Database    string
	TablePrefix string
}

type RedisConfig struct {
	Alias    string
	Password string
	Database int
	Host     string
	Port     int
}

type SessionConfig struct {
	Driver   string //redis
	Password string
	Host     string
	Port     int
	KeyPairs string
}

type RPCConfig struct {
	ServiceName string
	Addr        string
}
