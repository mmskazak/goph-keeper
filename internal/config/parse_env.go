package config

import "os"

func parseEnv() *Config {
	config := NewConfig()

	if envServAddr, ok := os.LookupEnv("SERVER_ADDRESS"); ok {
		config.Address = envServAddr
	}

	if envLogLevel, ok := os.LookupEnv("LOG_LEVEL"); ok {
		config.LogLevel = LogLevel(envLogLevel)
	}

	if dataBaseDSN, ok := os.LookupEnv("DATABASE_DSN"); ok {
		config.DataBaseDSN = dataBaseDSN
	}

	if secretKey, ok := os.LookupEnv("SECRET_KEY"); ok {
		config.SecretKey = secretKey
	}

	// Добавление поддержки 32-байтного ключа для шифрования
	if encryptionKeyHex, ok := os.LookupEnv("ENCRYPTION_KEY"); ok {
		config.EncryptionKeyString = encryptionKeyHex
	}

	return config
}
