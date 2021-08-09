package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"golang-streaming/video-server/api/handler"
	"golang-streaming/video-server/api/session"
	"net/http"
)


type middleWareHandler struct {
	r *httprouter.Router
}

func NewMiddleWareHandler(r *httprouter.Router) http.Handler {
	m := middleWareHandler{}
	m.r = r
	return r
}

func (m middleWareHandler) ServerHTTP(w http.ResponseWriter, r *http.Request)  {
	handler.ValidateUserSession(r)
	m.r.ServeHTTP(w, r)
}

func RegisterHandlers() *httprouter.Router  {
	router := httprouter.New()

	router.POST("/user",handler.CreateUser)

	router.POST("/user/:username", handler.Login)

	router.GET("/user/:username",handler.GetUserInfo)

	router.POST("/user/:username/videos", handler.AddNewVideo)

	router.GET("/user/:username/videos", handler.ListAllVideos)

	router.DELETE("/user/:username/videos/:vid-id", handler.DeleteVideo)

	router.POST("/videos/:vid-id/comments", handler.PostComment)

	router.GET("/videos/:vid-id/comments", handler.ShowComments)

	return router
}

func  Prepare()  {
	session.LoadSessionsFromDB()
}

func main() {
	Prepare()
	r := RegisterHandlers()
	mh := NewMiddleWareHandler(r)
	fmt.Println("Server Start at :8000")
	http.ListenAndServe(":8000", mh)
}
