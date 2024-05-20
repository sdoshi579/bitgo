package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/sodhi579/bitgo/app/notification/entity"
	"github.com/sodhi579/bitgo/app/notification/repository"
	"github.com/sodhi579/bitgo/app/notification/value_objects"
)

type Service interface {
	CreateNotification(ctx context.Context, notification entity.Notification) (*entity.Notification, error)
	GetNotifications(ctx context.Context, status *value_objects.Status) ([]entity.Notification, error)
	DeleteNotification(ctx context.Context, id uuid.UUID) (bool, error)
}

type serviceImplementation struct {
	repository repository.Repository
}

func NewService(repository repository.Repository) Service {
	return &serviceImplementation{
		repository: repository,
	}
}

func (s *serviceImplementation) CreateNotification(ctx context.Context, notification entity.Notification) (*entity.Notification, error) {
	return s.repository.CreateNotification(ctx, notification)
}

func (s *serviceImplementation) GetNotifications(ctx context.Context, status *value_objects.Status) ([]entity.Notification, error) {
	return s.repository.GetNotifications(ctx, status)
}

func (s *serviceImplementation) DeleteNotification(ctx context.Context, id uuid.UUID) (bool, error) {
	return s.repository.DeleteNotification(ctx, id)
}
