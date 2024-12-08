package auth

import (
	"context"
	"errors"
	"time"

	"fmt"
	"livecode/internal/database"
	"livecode/internal/models"
	"livecode/internal/services/jwt"

	"github.com/google/uuid"
)

type Auth struct {
	usrSaver    UserSaver
	usrProvider UserProvider
}

type UserSaver interface {
	SaveUser(
		ctx context.Context,
		email string,
		password string,
	) (userUUID uuid.UUID, err error)
}

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
)

type UserProvider interface {
	User(ctx context.Context, email string) (models.User, error)
}

func New(
	userSaver UserSaver,
	userProvider UserProvider,
) *Auth {
	return &Auth{
		usrSaver:    userSaver,
		usrProvider: userProvider,
	}
}

func (a *Auth) Login(
	ctx context.Context,
	email string,
	password string,
) (string, error) {
	const op = "Auth.Login"

	user, err := a.usrProvider.User(ctx, email)
	if err != nil {
		if errors.Is(err, database.ErrUserNotFound) {
			return "", fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		}

		return "", fmt.Errorf("%s: %w", op, err)
	}

	if ok, err := ComparePassword(password, user.PasswordHash); !ok || err != nil {

		return "", fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
	}

	token, err := jwt.NewToken(user, time.Duration(60400))
	if err != nil {

		return "", fmt.Errorf("%s: %w", op, err)
	}

	return token, nil
}

func (a *Auth) RegisterNewUser(ctx context.Context, email string, pass string) (uuid.UUID, error) {
	const op = "Auth.RegisterNewUser"

	passHash, err := GeneratePassword(pass)
	if err != nil {

		return uuid.Nil, fmt.Errorf("%s: %w", op, err)
	}

	userUUID, err := a.usrSaver.SaveUser(ctx, email, passHash)
	if err != nil {

		return uuid.Nil, fmt.Errorf("%s: %w", op, err)
	}

	return userUUID, nil
}
