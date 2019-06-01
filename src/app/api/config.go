package api


type Config struct {
	DB *DBConfig
}
type DBConfig struct {
	Dialect  string
	Username string
	Password string
	Name     string
	Charset  string
	PortNo   int
	Host     string
}

func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Dialect:  "postgres",
			Username: "postgres",
			Password: "root",
			Name:     "testDb",
			Host:     "localhost",
			Charset:  "utf8",
			PortNo:   5432,
		},
	}
}
