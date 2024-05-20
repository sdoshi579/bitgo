package repository

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/sodhi579/bitgo/app/notification/entity"
	"github.com/sodhi579/bitgo/app/notification/value_objects"
	"log"
)

type Repository interface {
	CreateNotification(ctx context.Context, notification entity.Notification) (*entity.Notification, error)
	GetNotifications(ctx context.Context, status *value_objects.Status) ([]entity.Notification, error)
	DeleteNotification(ctx context.Context, id uuid.UUID) (bool, error)
}

type repository struct {
	db     map[uuid.UUID]entity.Notification
	logger log.Logger
}

func NewRepository() Repository {
	return &repository{
		db: make(map[uuid.UUID]entity.Notification),
		//logger: log.New(),
	}
}

func (r *repository) CreateNotification(ctx context.Context, notification entity.Notification) (*entity.Notification, error) {
	if notification.ID == uuid.Nil {
		notification.ID = uuid.New()
	}

	r.db[notification.ID] = notification
	return &notification, nil
}

func (r *repository) GetNotifications(ctx context.Context, status *value_objects.Status) ([]entity.Notification, error) {
	var response []entity.Notification

	for _, n := range r.db {
		if status == nil || n.Status == *status {
			response = append(response, n)
		}
	}
	return response, nil
}

func (r *repository) DeleteNotification(ctx context.Context, id uuid.UUID) (bool, error) {
	if _, ok := r.db[id]; ok {
		delete(r.db, id)
		return true, nil
	}

	return false, fmt.Errorf("error in deleting notification: %s", id.String())
}
