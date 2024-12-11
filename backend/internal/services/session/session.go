package session

import (
	"context"
	"errors"
	"fmt"
	"livecode/internal/database"
	"livecode/internal/models"
	"livecode/internal/services/filestorage"
	"time"

	"github.com/google/uuid"
)

var (
	ErrSessionNotFound = errors.New("session not found")
)

type SessionService struct {
	ssnUpdater  SessionUpdater
	ssnProvider SessionProvider
	usrBlocker  UserBlocker
	s3Serve     *filestorage.S3Service
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
	DeleteExpiredSession(ctx context.Context) ([]uuid.UUID, error)
}

type SessionProvider interface {
	GetSessionById(ctx context.Context,
		sessionUUID uuid.UUID) (models.Session, error)
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
	s3Serv *filestorage.S3Service,
) *SessionService {
	ssnService := &SessionService{
		ssnUpdater:  SessionSaver,
		ssnProvider: SessionProv,
		usrBlocker:  UserBlock,
		s3Serve:     s3Serv,
	}
	closeChan := make(chan struct{})
	go ssnService.startExpirationChecker(closeChan, time.Hour*24)

	return ssnService
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

func (ssn *SessionService) garbageCollector() {
	rows, err := ssn.ssnUpdater.DeleteExpiredSession(context.Background())
	for _, row := range rows {
		ssn.s3Serve.DeleteProject(row.String())
	}
	if err != nil {
		fmt.Println(err)
	}
}

func (snn *SessionService) startExpirationChecker(closeChan chan struct{}, tm time.Duration) {
	for {
		select {
		case <-closeChan:
			return
		case <-time.After(tm):
			snn.garbageCollector()
		}
	}
}
