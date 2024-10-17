package crypto

import (
	"crypto/rand"                  // Для генерации случайных данных
	"fmt"                          // Пакет для форматирования строк
	"golang.org/x/crypto/chacha20" // Пакет для работы с ChaCha20
)

// Функция для шифрования текста
func encrypt(key [32]byte, plaintext []byte) (string, error) {
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

// Функция для расшифровки текста
func decrypt(key [32]byte, encrypted string) ([]byte, error) {
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

func main() {
	// Секретный ключ, который должен быть длиной 32 байта
	var key [32]byte
	if _, err := rand.Read(key[:]); err != nil {
		fmt.Println("Ошибка при генерации ключа:", err)
		return
	}

	// Открытый текст для шифрования
	plaintext := []byte("password123")

	// Шифруем текст
	encrypted, err := encrypt(key, plaintext)
	if err != nil {
		fmt.Println("Ошибка шифрования:", err)
		return
	}

	fmt.Println("Encrypted:", encrypted)

	// Расшифровываем текст
	decrypted, err := decrypt(key, encrypted)
	if err != nil {
		fmt.Println("Ошибка расшифровки:", err)
		return
	}

	fmt.Println("Decrypted:", string(decrypted))
}
