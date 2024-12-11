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
	ErrSessionNotFound   = errors.New("session not found")
	ErrTemplatenNotFound = errors.New("session not found")
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
	SaveTemplate(ctx context.Context, templateName string,
		lang string, template_code string, creatorID uuid.UUID) (uuid.UUID, error)
}

type SessionProvider interface {
	GetSessionById(ctx context.Context,
		sessionUUID uuid.UUID) (models.Session, error)
	GetTemplateByID(ctx context.Context, templateUUID uuid.UUID) (models.Template, error)
	GetAllTemplates(ctx context.Context) ([]models.Template, error)
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

func (ssn *SessionService) TemplateByID(ctx context.Context, templateUUID uuid.UUID) (models.Template, error) {
	const op = "Session.TemplateByID"

	template, err := ssn.ssnProvider.GetTemplateByID(ctx, templateUUID)
	if err != nil {
		if errors.Is(err, database.ErrTemplateNotFound) {
			return template, ErrTemplatenNotFound
		}
		return template, fmt.Errorf("%s: %w", op, err)
	}

	return template, nil
}

func (ssn *SessionService) AllTemplates(ctx context.Context) ([]models.Template, error) {
	const op = "Session.AllTemplates"

	templates, err := ssn.ssnProvider.GetAllTemplates(ctx)
	if err != nil {
		return templates, fmt.Errorf("%s: %w", op, err)
	}

	return templates, nil
}

func (ssn *SessionService) CreateTemplate(ctx context.Context, templateName string,
	lang string, template_code string, creatorID uuid.UUID) (uuid.UUID, error) {
	const op = "Session.CreateTemplate"

	templateUUID, err := ssn.ssnUpdater.SaveTemplate(ctx, templateName, lang, template_code, creatorID)
	if err != nil {
		return uuid.Nil, fmt.Errorf("%s: %w", op, err)
	}

	return templateUUID, nil
}