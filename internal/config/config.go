package config

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
		DataBaseDSN: "",
	}
}

// InitConfig инициализирует конфигурацию из флагов командной строки и переменных окружения.
// Возвращает указатель на структуру Config и ошибку в случае её возникновения.
func InitConfig() (*Config, error) {
	config := NewConfig()
	config = parseFlags()
	config = parseEnv()
	//TODO: надо будет использовать mergo
	return config, nil
}
