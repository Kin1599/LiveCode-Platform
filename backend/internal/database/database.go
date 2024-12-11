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

	ErrTemplateNotFound = errors.New("template not found")
)

const (
	saveNewUser          = "INSERT INTO \"Users\"(id, email, avatar, password_hash, created_at, updated_at) VALUES($1, $2, $3, $4, $5, $6);"
	getUserByEmail       = "SELECT id, email, password_hash FROM \"Users\" WHERE email = $1"
	saveNewSession       = "INSERT INTO \"Sessions\" VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)"
	getSessionById       = "SELECT * FROM \"Sessions\" WHERE id = $1"
	deleteSession        = "DELETE FROM \"Sessions\" WHERE id = $1"
	insertIp             = "INSERT INTO \"SessionBlock\" VALUES($1, $2, $3)"
	deleteBlockBySession = "DELETE FROM \"SessionBlock\" WHERE session_id = $1"
	getUserPublicInfo    = "SELECT id, nickname, avatar, email FROM \"Users\" WHERE email = $1"
	getAllTemplates      = "SELECT id, name, language, template_code, created_by FROM \"Templates\""
	getTempleByID        = "SELECT id, name, language, template_code, created_by FROM \"Templates\" WHERE id = $1"
	saveNewTemplate      = "INSERT INTO \"Templates\" VALUES($1, $2, $3, $4, $5, $6, $7)"
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

func (s *Storage) UserPublicInfo(ctx context.Context, email string) (models.User, error) {
	const op = "database.UserPublicInfo"

	stmt, err := s.db.Prepare(getUserPublicInfo)
	if err != nil {
		return models.User{}, fmt.Errorf("%s: %w", op, err)
	}

	row := stmt.QueryRowContext(ctx, email)

	var user models.User
	err = row.Scan(&user.ID, &user.Nickname, &user.Avatar, &user.Email)

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

	timeNow := time.Now()

	_, err = stmt.ExecContext(ctx, ID, ownerId,
		title, lang, "Public", timeNow.Add(time.Hour*24),
		maxUsers, isEditable, timeNow, timeNow, 1,
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
	err = row.Scan(&ssn.ID, &ssn.IdOwner, &ssn.Title, &ssn.Language, &ssn.AccessType,
		&ssn.ExpirationTime, &ssn.MaxUsers, &ssn.IsEditable, &ssn.CreatedAt, &ssn.UpdatedAt, &ssn.IsActive)
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

func (s *Storage) AddBlockedUser(ctx context.Context, blockedIp string, sessionUUID uuid.UUID) (uuid.UUID, error) {
	const op = "database.AddBlockedUser"

	stmt, err := s.db.Prepare(insertIp)
	if err != nil {
		return uuid.Nil, fmt.Errorf("%s: %w", op, err)
	}

	newUUID := uuid.New()
	_, err = stmt.ExecContext(ctx, newUUID, blockedIp, sessionUUID)

	if err != nil {
		pgErr, ok := err.(*pq.Error)
		if ok && pgErr.Code == "23505" {
			return uuid.Nil, fmt.Errorf("%s: %w", op, ErrUserExists)
		}

		return uuid.Nil, fmt.Errorf("%s: %w", op, err)
	}

	return newUUID, nil
}

func (s *Storage) DeleteAllBySession(ctx context.Context, sessionUUID uuid.UUID) error {
	const op = "database.GetSessionById"

	stmt, err := s.db.Prepare(deleteBlockBySession)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.ExecContext(ctx, sessionUUID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *Storage) GetAllTemplates(ctx context.Context) ([]models.Template, error) {
	const op = "database.GetAllTemplates"

	stmt, err := s.db.Prepare(getAllTemplates)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	rows, err := stmt.Query()
	if err != nil {
		defer rows.Close()
		if errors.Is(err, sql.ErrNoRows) {
			return []models.Template{}, nil
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	templates := []models.Template{}

	for rows.Next() {
		var template models.Template

		if err := rows.Scan(&template.ID, &template.Name, &template.Language, &template.TemplateCode, &template.CreatedBy); err != nil {
			return nil, err
		}
		templates = append(templates, template)
	}

	return templates, nil
}

func (s *Storage) GetTemplateByID(ctx context.Context, templateUUID uuid.UUID) (models.Template, error) {
	const op = "database.GetTemplateByID"

	stmt, err := s.db.Prepare(getTempleByID)
	if err != nil {
		return models.Template{}, fmt.Errorf("%s: %w", op, err)
	}

	row := stmt.QueryRowContext(ctx, templateUUID)

	var template models.Template
	err = row.Scan(&template.ID, &template.Name, &template.Language, &template.TemplateCode, &template.CreatedBy)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Template{}, fmt.Errorf("%s: %w", op, ErrUserNotFound)
		}

		return models.Template{}, fmt.Errorf("%s: %w", op, err)
	}

	return template, nil
}

func (s *Storage) SaveTemplate(ctx context.Context, templateName string,
	lang string, template_code string, creatorID uuid.UUID) (uuid.UUID, error) {
	const op = "database.SaveTemplate"

	stmt, err := s.db.Prepare(saveNewTemplate)
	if err != nil {
		return uuid.Nil, fmt.Errorf("%s: %w", op, err)
	}

	newUUID := uuid.New()
	timeNow := time.Now()
	_, err = stmt.ExecContext(ctx, newUUID, templateName, lang, template_code, creatorID, timeNow, timeNow)

	if err != nil {
		return uuid.Nil, fmt.Errorf("%s: %w", op, err)
	}

	return newUUID, nil
}
