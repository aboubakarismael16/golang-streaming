package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"golang-streaming/video-server/web/handler"
	"net/http"
)

func RegisterHandler() *httprouter.Router  {
	router := httprouter.New()

	router.GET("/", handler.HomeHandler)

	router.POST("/", handler.HomeHandler)

	router.GET("/userhome", handler.UserHomeHandler)

	router.POST("/userhome", handler.UserHomeHandler)

	router.POST("/api", handler.ApiHandler)

	router.POST("/upload/:vid-id", handler.ProxyHandler)

	router.ServeFiles("/statics/*filepath", http.Dir("./templates"))

	return router
}

func main() {

	r := RegisterHandler()
	fmt.Println("Starting server at :8080")
	http.ListenAndServe(":8080", r)
}
