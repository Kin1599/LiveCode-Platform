package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"livecode/internal/config"
	"livecode/internal/models"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Storage struct {
	db *sql.DB
}

var (
	ErrUserExists   = errors.New("user already exists")
	ErrUserNotFound = errors.New("user not found")
	ErrAppNotFound  = errors.New("app not found")
)

const (
	saveNewUser    = "INSERT INTO \"Users\"(id, email, avatar, password_hash, created_at, updated_at) VALUES($1, $2, $3, $4, $5, $6);"
	getUserByEmail = "SELECT id, email, password_hash FROM \"Users\" WHERE email = $1"
)

func conStringFromCfg(storageCfg config.StorageConfig) string {
	return fmt.Sprintf("postgres://%s:%s@localhost:%d/%s?sslmode=disable",
		storageCfg.User,
		storageCfg.Pass,
		storageCfg.Port,
		storageCfg.Name,
	)
}

func New(storageCfg config.StorageConfig) (*Storage, error) {
	const op = "database.New"

	db, err := sql.Open("postgres", conStringFromCfg(storageCfg))
	if err != nil {
		return nil, fmt.Errorf("%s %w", op, err)
	}

	return &Storage{db: db}, nil
}

func (s *Storage) Stop() error {
	return s.db.Close()
}

func (s *Storage) SaveUser(ctx context.Context, email string, passHash string) (int64, error) {
	const op = "databse.SaveUser"

	stmt, err := s.db.Prepare(saveNewUser)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	timeNow := time.Now()
	_, err = stmt.ExecContext(ctx, uuid.New(), email, "", passHash, timeNow, timeNow)

	if err != nil {
		pgErr, ok := err.(*pq.Error)
		if ok && pgErr.Code == "23505" {
			return 0, fmt.Errorf("%s: %w", op, ErrUserExists)
		}

		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return 0, nil
}

func (s *Storage) User(ctx context.Context, email string) (models.User, error) {
	const op = "database.User"

	stmt, err := s.db.Prepare(getUserByEmail)
	if err != nil {
		return models.User{}, fmt.Errorf("%s: %w", op, err)
	}

	row := stmt.QueryRowContext(ctx, email)

	var user models.User
	err = row.Scan(&user.ID, &user.Email, &user.PasswordHash)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.User{}, fmt.Errorf("%s: %w", op, ErrUserNotFound)
		}

		return models.User{}, fmt.Errorf("%s: %w", op, err)
	}

	return user, nil
}
