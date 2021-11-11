package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/lapitskyss/go_backend_2/internal/server/httpserver"
	"github.com/lapitskyss/go_backend_2/internal/services/scopeservice"
	"github.com/lapitskyss/go_backend_2/internal/services/userservice"
	"github.com/lapitskyss/go_backend_2/internal/storage/pgstore/scopestore"
	"github.com/lapitskyss/go_backend_2/internal/storage/pgstore/userstore"
)

func main() {
	userStore := userstore.NewUserStore()
	userService := userservice.NewUserService(userStore)

	scopeStore := scopestore.NewScopeStore()
	scopeService := scopeservice.NewScopeService(scopeStore)

	server := httpserver.NewHTTPServer(userService, scopeService)

	server.Start()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case x := <-interrupt:
		log.Println("Received a signal.", "signal", x.String())
	case err := <-server.Notify():
		log.Println("Received an error from http server.", "err", err)
	}
}
