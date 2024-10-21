package config

import "flag"

func parseFlags() *Config {
	config := NewConfig()

	flag.StringVar(&config.Address, "a", config.Address, "IP-адрес сервера")
	flag.StringVar(&config.DataBaseDSN, "d", "", "Database connection string")
	flag.StringVar(&config.SecretKey, "secret", config.SecretKey, "Secret key for authorization JWT token")
	flag.StringVar((*string)(&config.LogLevel), "l", string(config.LogLevel), "Log level")
	flag.StringVar(&config.EncryptionKeyString, "encryption_key", "", "32-byte encryption key in hex format")

	// Разбор командной строки
	flag.Parse()

	return config
}
