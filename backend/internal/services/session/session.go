package session

import (
	"context"
	"errors"
	"fmt"
	"livecode/internal/database"
	"livecode/internal/models"

	"github.com/google/uuid"
)

var (
	ErrSessionNotFound = errors.New("session not found")
)

type SessionService struct {
	ssnUpdater  SessionUpdater
	ssnProvider SessionProvider
}

type SessionUpdater interface {
	SaveSession(ctx context.Context,
		ID uuid.UUID,
		ownerId uuid.UUID,
		title string,
		lang string,
		acces string,
		maxUsers int64,
		isEditable bool) (uuid.UUID, error)
	DeleteSessionById(ctx context.Context, sessionUUID uuid.UUID) error
}

type SessionProvider interface {
	GetSessionById(ctx context.Context,
		sessionUUID uuid.UUID) (models.Session, error)
}

func New(
	SessionSaver SessionUpdater,
	SessionProv SessionProvider,
) *SessionService {
	return &SessionService{
		ssnUpdater:  SessionSaver,
		ssnProvider: SessionProv,
	}
}

func (ssn *SessionService) CreateNewSession(ctx context.Context,
	ID uuid.UUID,
	ownerId uuid.UUID,
	title string,
	lang string,
	access string,
	maxUsers int64,
	isEditable bool,
) (uuid.UUID, error) {
	const op = "Session.CreateNewSession"

	sessionUUID, err := ssn.ssnUpdater.SaveSession(ctx, ID, ownerId, title, lang, access, maxUsers, isEditable)
	if err != nil {
		return uuid.Nil, fmt.Errorf("%s: %w", op, err)
	}

	return sessionUUID, nil
}

func (ssn *SessionService) GetSession(ctx context.Context, sessionUUID uuid.UUID) (models.Session, error) {
	const op = "Session.GetSession"

	session, err := ssn.ssnProvider.GetSessionById(ctx, sessionUUID)
	if err != nil {
		if errors.Is(err, database.ErrUserNotFound) {
			return session, ErrSessionNotFound
		}
		return session, fmt.Errorf("%s: %w", op, err)
	}

	return session, nil
}

func (ssn *SessionService) DeleteSession(ctx context.Context, sessionUUID uuid.UUID) error {
	const op = "Session.GetSession"

	err := ssn.ssnUpdater.DeleteSessionById(ctx, sessionUUID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
