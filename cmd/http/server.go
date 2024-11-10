package main

import (
	"context"
	httpHandler "github.com/davidsugianto/go-chi-http/internal/handler/http"
	"github.com/davidsugianto/go-chi-http/internal/pkg/provider"
	userRepos "github.com/davidsugianto/go-chi-http/internal/repository/user"
	userUseCase "github.com/davidsugianto/go-chi-http/internal/usecase/user"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func start(addressHTTP string) error {
	ctx := context.Background()

	// Initialize Client
	db := provider.InitGorm(ctx)

	// Initialize Repository
	userRepo, err := userRepos.New(db)
	if err != nil {
		return err
	}

	// Initialize UseCase
	userUC := userUseCase.New(userRepo)

	// Initialize Handler
	handler := httpHandler.New(userUC)

	// Initialize Router
	router := httpHandler.NewRouters(handler)

	return serve(addressHTTP, router)
}

func serve(address string, handler http.Handler) error {
	srv := http.Server{
		Handler: handler,
		Addr:    address,
	}
	log.Infof("service listen to %s", address)

	return srv.ListenAndServe()
}
