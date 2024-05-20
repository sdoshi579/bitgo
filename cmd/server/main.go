package main

import (
	"github.com/sodhi579/bitgo/api/notification"
	repository2 "github.com/sodhi579/bitgo/app/notification/repository"
	service2 "github.com/sodhi579/bitgo/app/notification/service"
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

	http.HandleFunc("/", notificationHandler.CreateNotification)
	http.HandleFunc("/get", notificationHandler.GetNotification)

	log.Fatal(http.ListenAndServe(":8081", nil))
}
