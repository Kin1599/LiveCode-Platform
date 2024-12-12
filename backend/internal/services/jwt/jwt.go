package jwt

import (
	"fmt"
	"livecode/internal/models"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

const hmacSecret = "SecretValueReplaceThis"

type Claims struct {
	UID   string `json:"uid"`
	Email string `json:"Email"`
	jwt.RegisteredClaims
}

func NewToken(user models.User, duration time.Duration) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["uid"] = user.ID
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(duration).Unix()

	tokenString, err := token.SignedString([]byte(hmacSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(tokenString string) (models.User, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(hmacSecret), nil
	})

	// Проверка ошибки парсинга
	if err != nil {
		log.Printf("Error parsing token: %v", err)
		return models.User{}, fmt.Errorf("failed to parse token: %w", err)
	}

	// Проверка валидности токена
	if !token.Valid {
		return models.User{}, fmt.Errorf("invalid token")
	}

	// Извлечение claims
	claims, ok := token.Claims.(*Claims)
	if !ok {
		return models.User{}, fmt.Errorf("invalid claims")
	}

	var authUser models.User
	authUser.ID = uuid.MustParse(claims.UID)
	authUser.Email = claims.Email
	return authUser, nil
}

func NewRefreshToken(user models.User) (string, error) {
	// Создание claims для refreshToken
	claims := jwt.MapClaims{
		"uid": user.ID.String(),
		"exp": time.Now().Add(time.Hour * 24 * 7).Unix(), // refreshToken действителен 7 дней
	}

	// Создание токена
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Подпись токена
	tokenString, err := token.SignedString([]byte(hmacSecret))
	if err != nil {
		return "", fmt.Errorf("failed to sign refresh token: %w", err)
	}

	return tokenString, nil
}
