package notification

import (
	"encoding/json"
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
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	if err := json.Unmarshal(payload, &request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	resp, err := h.NotificationService.CreateNotification(req.Context(), entity.Notification{
		UserID: request.UserID,
		Status: request.Status,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	bytesResponse, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(bytesResponse)
}

func (h *Handler) GetNotification(w http.ResponseWriter, req *http.Request) {
	status := req.URL.Query().Get("status")

	var resp []entity.Notification
	var err error
	if status == "" {
		resp, err = h.NotificationService.GetNotifications(req.Context(), nil)
	} else {
		statusEnum := value_objects.GetStatusEnum(status)
		resp, err = h.NotificationService.GetNotifications(req.Context(), &statusEnum)
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	bytesResponse, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(bytesResponse)
}
