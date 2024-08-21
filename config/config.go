package config

type Config struct {
	Version  int
	Env      string
	Network  Network
	Postgres Postgres
}

type Network struct {
	Address      string
	Port         string
	WriteTimeout int
	ReadTimeout  int
}

type Postgres struct {
	Host         string
	Port         string
	Username     string
	Password     string
	DatabaseName string
	Sslmode      string
}

func GetConfig() (*Config, error) {

	conf := &Config{
		Version: 1,
		Env:     "dev",
		Network: Network{
			Address:      "127.0.0.1:",
			Port:         "8080",
			WriteTimeout: 15,
			ReadTimeout:  15,
		},
		Postgres: Postgres{
			Host:         "localhost",
			Port:         "5432",
			Username:     "postgres",
			Password:     "0000",
			DatabaseName: "testDB",
			Sslmode:      "disable",
		},
	}

	return conf, nil
}
