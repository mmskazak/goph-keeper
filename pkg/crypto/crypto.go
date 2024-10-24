package crypto

import (
	"crypto/rand" // Для генерации случайных данных
	"encoding/hex"
	"fmt" // Пакет для форматирования строк
	"strings"

	"golang.org/x/crypto/chacha20" // Пакет для работы с ChaCha20
)

// Encrypt Функция для шифрования текста.
func Encrypt(key [32]byte, plaintext []byte) (string, error) {
	// Создаём новый объект шифрования ChaCha20 с заданным ключом и инициализационным вектором (nonce)
	var nonce [12]byte
	if _, err := rand.Read(nonce[:]); err != nil {
		return "", fmt.Errorf("error generating nonce: %w", err) // Ошибка при генерации nonce
	}

	// Создаём срез для зашифрованного текста
	ciphertext := make([]byte, len(plaintext))

	// Создаём объект шифрования
	stream, err := chacha20.NewUnauthenticatedCipher(key[:], nonce[:])
	if err != nil {
		return "", fmt.Errorf("error creating cipher: %w", err) // Обработка ошибки
	}

	stream.XORKeyStream(ciphertext, plaintext) // Шифруем текст

	// Возвращаем зашифрованный текст в hex-формате вместе с nonce для дальнейшего использования
	return fmt.Sprintf("%x:%x", nonce, ciphertext), nil
}

// Decrypt Функция для расшифровки текста.
func Decrypt(key [32]byte, encrypted string) ([]byte, error) {
	// Разделяем зашифрованный текст и nonce
	parts := strings.Split(encrypted, ":")
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid format: expected nonce:ciphertext")
	}

	// Парсим nonce
	var nonce [12]byte
	nonceBytes, err := hex.DecodeString(parts[0])
	if err != nil || len(nonceBytes) != 12 {
		return nil, fmt.Errorf("error parsing nonce: %w", err)
	}
	copy(nonce[:], nonceBytes)

	// Парсим ciphertext
	ciphertext, err := hex.DecodeString(parts[1])
	if err != nil {
		return nil, fmt.Errorf("error parsing ciphertext: %w", err)
	}

	// Создаём срез для расшифрованного текста
	plaintext := make([]byte, len(ciphertext))

	// Создаём объект шифрования
	stream, err := chacha20.NewUnauthenticatedCipher(key[:], nonce[:])
	if err != nil {
		return nil, fmt.Errorf("error creating cipher: %w", err)
	}

	stream.XORKeyStream(plaintext, ciphertext) // Расшифровываем текст

	return plaintext, nil
}
