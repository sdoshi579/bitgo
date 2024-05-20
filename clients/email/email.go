package email

import (
	"context"
	"fmt"
)

type Email interface {
	SendEmail(ctx context.Context, toEmail string, message string) error
}

type email struct {
	// ses client
}

func NewEmailClient() Email {
	return &email{}
}

func (e *email) SendEmail(ctx context.Context, toEmail string, message string) error {
	fmt.Printf("sending mail to: %s, message: %s \n", toEmail, message)
	return nil
}
