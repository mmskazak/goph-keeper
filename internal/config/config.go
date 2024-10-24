package config

import (
	"errors"
	"fmt"

	"dario.cat/mergo"
)

// Config содержит поля конфигурации.
type Config struct {
	LogLevel            LogLevel `json:"log_level"`          // Уровень логирования
	EncryptionKey       [32]byte `json:"-"`                  // 32-байтный ключ в байтах (не сериализуется)
	Address             string   `json:"address"`            // Адрес сервера
	DataBaseDSN         string   `json:"database_dsn"`       // Строка подключения к базе данных
	SecretKey           string   `json:"secret_key"`         // Секретный ключ JWT токена
	EncryptionKeyString string   `json:"encryption_key_hex"` // 32-байтный ключ для шифрования в hex
	DirSavedFiles       string   `json:"dir_saved_files"`    // Папка для сохранных файлов
}

func NewConfig() *Config {
	return &Config{
		Address:             ":8080",
		LogLevel:            "info",
		SecretKey:           "secret",
		DataBaseDSN:         "postgresql://gkuser:gkpass@localhost:5432/goph_keeper?sslmode=disable",
		EncryptionKeyString: "MySecretEncryptionKey1234567890a",
		DirSavedFiles:       "E:/test/",
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

	// Проверка и преобразование ключа из строки в байты
	if config.EncryptionKeyString == "" {
		return nil, errors.New("encryption key not provided")
	}

	encryptionKey := []byte(config.EncryptionKeyString)

	if len(encryptionKey) != 32 { //nolint:gomnd //количество байтов ключа шифрования
		return nil, fmt.Errorf("encryption key must be 32 bytes long, got %d bytes", len(encryptionKey))
	}

	var keyArray [32]byte
	copy(keyArray[:], encryptionKey)

	config.EncryptionKey = keyArray
	return config, nil
}

func mergeFlags(config *Config) (*Config, error) {
	configFromFlags := parseFlags()
	err := mergo.Merge(config, configFromFlags, mergo.WithOverride)
	if err != nil {
		return nil, fmt.Errorf("mergo from flags error: %w", err)
	}
	return config, nil
}

func mergeEnv(config *Config) (*Config, error) {
	configFromEnv := parseEnv()
	err := mergo.Merge(config, configFromEnv, mergo.WithOverride)
	if err != nil {
		return nil, fmt.Errorf("mergo from env error: %w", err)
	}

	return config, nil
}
