package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"golang-streaming/video-server/scheduler/handler"
	"golang-streaming/video-server/scheduler/taskrunner"
	"net/http"
)

func RegisterHandler() *httprouter.Router  {
	router := httprouter.New()

	router.GET("/video-delete-record/:vid-id", handler.VidDelRecHandler)

	return router
}


func main() {
	go taskrunner.Start()
	r := RegisterHandler()
	fmt.Println("Start Listen server at :9001")
	http.ListenAndServe(":9001", r)
}
