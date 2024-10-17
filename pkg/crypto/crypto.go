package crypto

import (
	"crypto/rand"                  // Для генерации случайных данных
	"fmt"                          // Пакет для форматирования строк
	"golang.org/x/crypto/chacha20" // Пакет для работы с ChaCha20
)

// Encrypt Функция для шифрования текста
func Encrypt(key [32]byte, plaintext []byte) (string, error) {
	// Создаём новый объект шифрования ChaCha20 с заданным ключом и инициализационным вектором (nonce)
	var nonce [12]byte
	if _, err := rand.Read(nonce[:]); err != nil {
		return "", err // Ошибка при генерации nonce
	}

	// Создаём срез для зашифрованного текста
	ciphertext := make([]byte, len(plaintext))

	// Создаём объект шифрования
	stream, _ := chacha20.NewUnauthenticatedCipher(key[:], nonce[:])
	stream.XORKeyStream(ciphertext, plaintext) // Шифруем текст

	// Возвращаем зашифрованный текст в hex-формате вместе с nonce для дальнейшего использования
	return fmt.Sprintf("%x:%x", nonce, ciphertext), nil
}

// Decrypt Функция для расшифровки текста
func Decrypt(key [32]byte, encrypted string) ([]byte, error) {
	// Разделяем зашифрованный текст и nonce
	var nonce [12]byte
	var ciphertext []byte
	if _, err := fmt.Sscanf(encrypted, "%x:%x", &nonce, &ciphertext); err != nil {
		return nil, err // Ошибка при парсинге зашифрованного текста
	}

	// Создаём срез для расшифрованного текста
	plaintext := make([]byte, len(ciphertext))

	// Создаём объект шифрования
	stream, _ := chacha20.NewUnauthenticatedCipher(key[:], nonce[:])
	stream.XORKeyStream(plaintext, ciphertext) // Расшифровываем текст

	return plaintext, nil
}
