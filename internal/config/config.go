package config

import (
	"dario.cat/mergo"
	"encoding/hex"
	"errors"
	"fmt"
)

// Config содержит поля конфигурации.
type Config struct {
	Address          string   `json:"address"`            // Адрес сервера
	DataBaseDSN      string   `json:"database_dsn"`       // Строка подключения к базе данных
	SecretKey        string   `json:"secret_key"`         // Секретный ключ JWT токена
	LogLevel         LogLevel `json:"log_level"`          // Уровень логирования
	EncryptionKeyHex string   `json:"encryption_key_hex"` // 32-байтный ключ для шифрования в hex
	EncryptionKey    []byte   `json:"-"`                  // 32-байтный ключ в байтах (не сериализуется)
	DirSavedFiles    string   `json:"dir_saved_files"`    // Папка для сохранных файлов
}

func NewConfig() *Config {
	return &Config{
		Address:          ":8080",
		LogLevel:         "info",
		SecretKey:        "secret",
		DataBaseDSN:      "postgresql://gkuser:gkpass@localhost:5432/goph_keeper?sslmode=disable",
		EncryptionKeyHex: "MySecretEncryptionKey1234567890",
		DirSavedFiles:    "/saved_files",
	}
}

// InitConfig инициализирует конфигурацию из флагов командной строки и переменных окружения.
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

	// Проверка и преобразование ключа из hex-строки в байты
	if config.EncryptionKeyHex == "" {
		return nil, errors.New("encryption key not provided")
	}

	encryptionKey, err := hex.DecodeString(config.EncryptionKeyHex)
	if err != nil {
		return nil, fmt.Errorf("invalid encryption key: %w", err)
	}

	if len(encryptionKey) != 32 {
		return nil, fmt.Errorf("encryption key must be 32 bytes long, got %d bytes", len(encryptionKey))
	}

	config.EncryptionKey = encryptionKey
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
