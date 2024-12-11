package jwt

import (
	"fmt"
	"livecode/internal/models"
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

	var authUser models.User

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		authUser.ID = uuid.MustParse(claims.UID)
		authUser.Email = claims.Email
		return authUser, nil
	} else {
		return models.User{}, err
	}
}
