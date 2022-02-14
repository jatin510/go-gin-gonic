package config

type Config struct {
	serverPort string
}

func GetConfig() *Config {
	config := Config{serverPort: "8000"}
	return &config
}
