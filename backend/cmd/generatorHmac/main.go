package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

func generateHMACSecret(length int) (string, error) {
	// Создаем буфер для хранения случайных байтов
	key := make([]byte, length)

	// Заполняем буфер случайными байтами
	_, err := rand.Read(key)
	if err != nil {
		return "", err
	}

	// Преобразуем байты в строку в формате hex
	return hex.EncodeToString(key), nil
}

func main() {
	// Генерируем HMAC-секрет длиной 32 байта (256 бит)
	hmacSecret, err := generateHMACSecret(32)
	if err != nil {
		fmt.Println("Ошибка при генерации HMAC-секрета:", err)
		return
	}

	fmt.Println("HMAC Secret:", hmacSecret)
}
