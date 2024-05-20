package notification

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/sodhi579/bitgo/api/notification/models"
	"github.com/sodhi579/bitgo/app/notification/entity"
	"github.com/sodhi579/bitgo/app/notification/service"
	"github.com/sodhi579/bitgo/app/notification/value_objects"
	"io"
	"net/http"
)

type Handler struct {
	NotificationService service.Service
}

func (h *Handler) CreateNotification(w http.ResponseWriter, req *http.Request) {
	var request models.CreateNotificationRequest

	payload, err := io.ReadAll(req.Body)
	if err != nil {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	if err := json.Unmarshal(payload, &request); err != nil {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	resp, err := h.NotificationService.CreateNotification(req.Context(), entity.Notification{
		UserID:            request.UserID,
		CurrentPrice:      request.CurrentPrice,
		Volume:            request.Volume,
		MarketCap:         request.MarketCap,
		IntraDayHighPrice: request.IntraDayHighPrice,
		Status:            request.Status,
	})
	if err != nil {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	bytesResponse, err := json.Marshal(resp)
	if err != nil {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(bytesResponse)
}

func (h *Handler) GetNotification(w http.ResponseWriter, req *http.Request) {
	status := req.URL.Query().Get("status")

	var statusEnum *value_objects.Status
	if status != "" {
		innerStatus := value_objects.GetStatusEnum(status)
		statusEnum = &innerStatus
	}
	resp, err := h.NotificationService.GetNotifications(req.Context(), statusEnum)
	if err != nil {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	bytesResponse, err := json.Marshal(resp)
	if err != nil {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(bytesResponse)
}

func (h *Handler) DeleteNotification(w http.ResponseWriter, req *http.Request) {
	notificationID := mux.Vars(req)["id"]

	notificationUUID, err := uuid.Parse(notificationID)
	if err != nil {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	_, err = h.NotificationService.DeleteNotification(req.Context(), notificationUUID)
	if err != nil {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("content-type", "application/json")
}
