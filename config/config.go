package config

type Config struct {
	Host string
	Port string
	// DatabaseUrl string
	DBname   string
	Username string
	TimeZone string
	Password string
	SSLmode  string
}

func New() Config {
	return Config{
		// DatabaseUrl: "go-test",
		// DatabaseUrl: "postgres://postgres:pass@127.0.0.1:5432/go-test",
		// DatabaseUrl: "postgres://postgres:pass@127.0.0.1:5432/testGo",
		DBname:   "testGo",
		Host:     "127.0.0.1",
		Port:     "9001",
		Username: "postgres",
		TimeZone: "Asia/Shanghai",
		Password: "",
		SSLmode:  "prefer",
	}
}
