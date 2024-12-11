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
	ErrSessionNotFound  = errors.New("session not found")
	ErrMessagesNotFound = errors.New("messages not found")
)

type SessionService struct {
	ssnUpdater  SessionUpdater
	ssnProvider SessionProvider
	usrBlocker  UserBlocker
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
	SaveMessage(ctx context.Context, idSession uuid.UUID, idSessionParticipant uuid.UUID, message string) (uuid.UUID, error)
	DeleteAllMessagesBySession(ctx context.Context, idSession uuid.UUID) error
}

type SessionProvider interface {
	GetSessionById(ctx context.Context,
		sessionUUID uuid.UUID) (models.Session, error)
	GetAllMessagesBySession(ctx context.Context, idSession uuid.UUID) ([]models.Message, error)
}

type UserBlocker interface {
	AddBlockedUser(ctx context.Context, blockedIp string,
		sessionUUID uuid.UUID) (uuid.UUID, error)
	DeleteAllBySession(ctx context.Context,
		sessionUUID uuid.UUID) error
}

func New(
	SessionSaver SessionUpdater,
	SessionProv SessionProvider,
	UserBlock UserBlocker,
) *SessionService {
	return &SessionService{
		ssnUpdater:  SessionSaver,
		ssnProvider: SessionProv,
		usrBlocker:  UserBlock,
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
	const op = "Session.DeleteSession"

	err := ssn.usrBlocker.DeleteAllBySession(ctx, sessionUUID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	err = ssn.ssnUpdater.DeleteAllMessagesBySession(ctx, sessionUUID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	err = ssn.ssnUpdater.DeleteSessionById(ctx, sessionUUID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (ssn *SessionService) BlockUser(ctx context.Context, blockedIp string, sessionUUID uuid.UUID) (uuid.UUID, error) {
	const op = "Session.BlockUser"

	blockedUserUUID, err := ssn.usrBlocker.AddBlockedUser(ctx, blockedIp, sessionUUID)
	if err != nil {
		return uuid.Nil, fmt.Errorf("%s: %w", op, err)
	}

	return blockedUserUUID, nil
}

func (ssn *SessionService) WriteMessage(ctx context.Context, idSession uuid.UUID, idSessionParticipant uuid.UUID, message string) (uuid.UUID, error) {
	const op = "Session.WriteMessage"

	messageUUID, err := ssn.ssnUpdater.SaveMessage(ctx, idSession, idSessionParticipant, message)
	if err != nil {
		return uuid.Nil, fmt.Errorf("%s: %w", op, err)
	}

	return messageUUID, nil
}

func (ssn *SessionService) GetMessagesBySession(ctx context.Context, idSession uuid.UUID) ([]models.Message, error) {
	const op = "Session.WriteMessage"

	messages, err := ssn.ssnProvider.GetAllMessagesBySession(ctx, idSession)
	if err != nil {
		if errors.Is(err, database.ErrUserNotFound) {
			return nil, ErrMessagesNotFound
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return messages, nil
}
