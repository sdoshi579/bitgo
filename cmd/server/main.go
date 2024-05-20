package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sodhi579/bitgo/api/notification"
	repository2 "github.com/sodhi579/bitgo/app/notification/repository"
	service2 "github.com/sodhi579/bitgo/app/notification/service"
	"github.com/sodhi579/bitgo/clients/email"
	"github.com/sodhi579/bitgo/workers"
	"log"
	"net/http"
)

func main() {

	//listener, err := net.Listen("localhost", "8081")
	//
	//if err != nil {
	//	fmt.Println("error in listener: ", err)
	//	os.Exit(1)
	//}

	repository := repository2.NewRepository()
	service := service2.NewService(repository)
	notificationHandler := notification.Handler{
		NotificationService: service,
	}

	//http.ListenAndServe(listener, )

	r := mux.NewRouter()
	r.HandleFunc("/", notificationHandler.CreateNotification).Methods(http.MethodPost)
	r.HandleFunc("/", notificationHandler.GetNotification).Methods(http.MethodGet)
	r.HandleFunc("/{id}", notificationHandler.DeleteNotification).Methods(http.MethodDelete)

	http.Handle("/", r)

	emailClient := email.NewEmailClient()
	w := workers.NewEmailWorker(service, emailClient)

	ctx, cancel := context.WithCancel(context.Background())
	go w.Run(ctx)
	fmt.Println("worker started")
	log.Fatal(http.ListenAndServe(":8081", nil))
	cancel()
}
