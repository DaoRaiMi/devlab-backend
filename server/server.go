package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/daoraimi/devlab-backend/application"
	"github.com/gin-gonic/gin"
)

type Server struct {
	engin *gin.Engine
	user  application.UserAppIface
}

func New(user application.UserAppIface) *Server {
	return &Server{
		engin: gin.Default(),
		user:  user,
	}
}

func (s *Server) Run(addr string) error {
	s.registerRoute()
	srv := http.Server{
		Addr:    addr,
		Handler: s.engin,
	}
	go func() {
		log.Print("listening and serving HTTP on ", addr)
		if err := srv.ListenAndServe(); err != nil {
			log.Print(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	sig := <-quit
	log.Printf("received %s signal, shutting down server...", sig)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("server forced to shutdown: ", err)
	}
	log.Print("server exiting...")
	return nil
}
