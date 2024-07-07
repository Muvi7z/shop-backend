package main

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"os"
	"shop-backend/internal/auth/handler"
	"shop-backend/internal/auth/service"
	"shop-backend/internal/config"
	"time"
)

func main() {
	router := gin.Default()

	cfg := config.GetConfig()

	logger := *slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	authService := service.New(&logger, nil, 10*time.Second)

	handlerAuth := handler.NewHandler(&logger, authService)

	handlerAuth.Register(router)

	s := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	s.ListenAndServe()
}
