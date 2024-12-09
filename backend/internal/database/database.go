package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
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
	saveNewSession = "INSERT INTO Sessions VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)"
	getSessionById = "SELECT * FROM Sessions WHERE id = $1"
	deleteSession  = "DELETE FROM Sessions WHERE id = $1"
)

func New(storagePath string) (*Storage, error) {
	const op = "database.New"

	db, err := sql.Open("postgres", storagePath)
	if err != nil {
		return nil, fmt.Errorf("%s %w", op, err)
	}

	if err := db.Ping(); err != nil {
		fmt.Println(storagePath)
		return nil, fmt.Errorf("%s: unable to connect to database %w", op, err)
	}

	return &Storage{db: db}, nil
}

func (s *Storage) Stop() error {
	return s.db.Close()
}

func (s *Storage) SaveUser(ctx context.Context, email string, passHash string) (uuid.UUID, error) {
	const op = "database.SaveUser"

	stmt, err := s.db.Prepare(saveNewUser)
	if err != nil {
		return uuid.Nil, fmt.Errorf("%s: %w", op, err)
	}

	defer stmt.Close()

	timeNow := time.Now()
	newUUID := uuid.New()
	_, err = stmt.ExecContext(ctx, newUUID, email, "", passHash, timeNow, timeNow)

	if err != nil {
		pgErr, ok := err.(*pq.Error)
		if ok && pgErr.Code == "23505" {
			return uuid.Nil, fmt.Errorf("%s: %w", op, ErrUserExists)
		}

		return uuid.Nil, fmt.Errorf("%s: %w", op, err)
	}

	return newUUID, nil
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

func (s *Storage) SaveSession(ctx context.Context,
	ID uuid.UUID,
	ownerId uuid.UUID,
	title string,
	lang string,
	access string,
	maxUsers int64,
	isEditable bool) (uuid.UUID, error) {
	const op = "database.SaveSession"

	stmt, err := s.db.Prepare(saveNewSession)
	if err != nil {
		return uuid.Nil, fmt.Errorf("%s: %w", op, err)
	}

	defer stmt.Close()

	timeNow := time.Now()

	_, err = stmt.ExecContext(ctx, ID, ownerId,
		title, lang, access, timeNow.Add(time.Hour*24),
		maxUsers, isEditable, timeNow, timeNow, '1',
	)

	if err != nil {
		return uuid.Nil, fmt.Errorf("%s: %w", op, err)
	}

	return ID, nil
}

func (s *Storage) GetSessionById(ctx context.Context, sessionUUID uuid.UUID) (models.Session, error) {
	const op = "database.GetSessionById"

	stmt, err := s.db.Prepare(getSessionById)
	if err != nil {
		return models.Session{}, fmt.Errorf("%s: %w", op, err)
	}

	row := stmt.QueryRowContext(ctx, sessionUUID)

	var ssn models.Session
	err = row.Scan(&ssn)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Session{}, fmt.Errorf("%s: %w", op, ErrUserNotFound)
		}
		return models.Session{}, fmt.Errorf("%s: %w", op, err)
	}

	return ssn, nil
}

func (s *Storage) DeleteSessionById(ctx context.Context, sessionUUID uuid.UUID) error {
	const op = "database.GetSessionById"

	stmt, err := s.db.Prepare(deleteSession)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.ExecContext(ctx, sessionUUID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
