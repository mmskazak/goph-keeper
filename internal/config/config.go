package config

import (
	"dario.cat/mergo"
	"fmt"
)

// Config содержит поля конфигурации.
type Config struct {
	Address     string   `json:"address"`      // Адрес сервера
	DataBaseDSN string   `json:"database_dsn"` // Строка подключения к базе данных
	SecretKey   string   `json:"secret_key"`   // Секретный ключ JWT токена
	LogLevel    LogLevel `json:"log_level"`    // Уровень логирования
}

func NewConfig() *Config {
	return &Config{
		Address:     ":8080",
		LogLevel:    "info",
		SecretKey:   "secret",
		DataBaseDSN: "postgresql://gkuser:gkpass@localhost:5432/goph_keeper?sslmode=disable",
	}
}

// InitConfig инициализирует конфигурацию из флагов командной строки и переменных окружения.
// Возвращает указатель на структуру Config и ошибку в случае её возникновения.
func InitConfig() (*Config, error) {
	config := NewConfig()
	config, err := mergeFlags(config)
	if err != nil {
		return nil, fmt.Errorf("error merge: %w", err)
	}

	config, err = mergeEnv(config)
	if err != nil {
		return nil, fmt.Errorf("error merge: %w", err)
	}
	return config, nil
}

func mergeFlags(config *Config) (*Config, error) {
	configFromFlags := parseFlags()
	err := mergo.Merge(config, configFromFlags, mergo.WithOverride)
	if err != nil {
		return nil, fmt.Errorf("mergo from flags error: %v", err)
	}
	return config, nil
}

func mergeEnv(config *Config) (*Config, error) {
	configFromEnv := parseEnv()
	err := mergo.Merge(config, configFromEnv, mergo.WithOverride)
	if err != nil {
		return nil, fmt.Errorf("mergo from env error: %v", err)
	}

	return config, nil
}
