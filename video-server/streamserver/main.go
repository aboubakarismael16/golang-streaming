package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"golang-streaming/video-server/streamserver/handler"
	"net/http"
)

type middleWareHandler struct {
	r *httprouter.Router
	l  *handler.ConnLimiter
}

func NewMiddleWareHandler(r *httprouter.Router, cc int) http.Handler {
	m := middleWareHandler{}
	m.r = r
	m.l = handler.NewConnLimiter(cc)

	return m

}

func RegisterHandlers() *httprouter.Router  {
	router := httprouter.New()

	router.GET("/videos/:vid-id", handler.StreamHandler)

	router.POST("/upload/:vid-id", handler.UploadHandler)

	router.GET("/testpage", handler.TestPageHandler)

	return router
}

func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	if !m.l.GetConn() {
		handler.SendErrorResponse(w, http.StatusTooManyRequests, "Too many requests")
		return
	}

	m.r.ServeHTTP(w, r)
	defer m.l.ReleaseConn()
}

func main() {
	r := RegisterHandlers()
	mh := NewMiddleWareHandler(r, 2)
	fmt.Println("StreamServer start at http://localhost:9000")
	http.ListenAndServe(":9000", mh)
}
