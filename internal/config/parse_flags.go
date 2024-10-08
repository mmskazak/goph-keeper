package config

import "flag"

func parseFlags() *Config {
	config := NewConfig()
	flag.StringVar(&config.Address, "a", config.Address, "IP-адрес сервера")
	flag.StringVar(&config.DataBaseDSN, "d", "", "Database connection string")
	flag.StringVar(&config.SecretKey, "secret", config.SecretKey, "Secret key for authorization JWT token")
	flag.StringVar((*string)(&config.LogLevel), "l", string(config.LogLevel), "log level")

	// Разбор командной строки
	flag.Parse()

	return config
}
