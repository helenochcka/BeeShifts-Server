package config

type Config struct {
	Server struct {
		Address   string
		Port      int
		SecretKey string `yaml:"secret_key"`
	}
}
