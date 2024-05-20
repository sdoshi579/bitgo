package workers

import (
	"context"
	"fmt"
	"github.com/sodhi579/bitgo/app/notification/service"
	"github.com/sodhi579/bitgo/app/notification/value_objects"
	"github.com/sodhi579/bitgo/clients/email"
	"time"
)

const TickerDuration = 1 * time.Second

type emailNotification struct {
	notificationService service.Service
	emailClient         email.Email
}

func NewEmailWorker(notificationService service.Service, emailClient email.Email) Worker {
	return &emailNotification{
		notificationService: notificationService,
		emailClient:         emailClient,
	}
}

func (e *emailNotification) Run(ctx context.Context) error {
	ticker := time.NewTicker(TickerDuration)

	for {
		select {
		case <-ctx.Done():
			fmt.Println("main server stopped")
		case <-ticker.C:
			statusEnum := value_objects.StatusOutstanding
			notifications, err := e.notificationService.GetNotifications(ctx, &statusEnum)
			if err != nil {
				fmt.Println("error in fetching notification: ", err)
				continue
			}
			for _, n := range notifications {
				if err := e.emailClient.SendEmail(ctx, n.UserID.String(), "notification triggered"); err != nil {
					fmt.Println("error in sending mail: ", err)
				}
			}
		}
	}
}
