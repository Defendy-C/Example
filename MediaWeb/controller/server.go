package controller

import (
	"net/http"
)

type mediaServer struct {
	server http.Server
	router *MyRouter
}

func NewServer(router *MyRouter, addr string)*mediaServer {
	router.AddHandler("user", DealUser)
	router.AddHandler("video", DealVideo)

	var s http.Server
	s.Addr = addr
	s.Handler = router
	return &mediaServer{server:s, router:router}
}

func(s *mediaServer)ListenAndServ()error {
	return s.server.ListenAndServe()
}




