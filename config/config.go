package config

type Config struct {
	Server struct {
		Address      string
		Port         int
		SecretKey    string `yaml:"secret_key"`
		TokenExpTime int    `yaml:"token_exp_time"`
	}

	DB struct {
		Host     string
		Port     int
		UserName string `yaml:"user_name"`
		Password string
		DBName   string `yaml:"dbname"`
	} `yaml:"db"`
}
