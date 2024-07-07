package service

import (
	"context"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log/slog"
	"shop-backend/internal/domain/user"
	"shop-backend/internal/lib/jwt"
	"time"
)

type Service struct {
	logger   *slog.Logger
	storage  UserStorage
	tokenTTL time.Duration
}

type UserStorage interface {
	GetUser(ctx context.Context, id string) (user.User, error)
}

func New(logger *slog.Logger, storage UserStorage, tokenTTL time.Duration) *Service {
	return &Service{logger, storage, tokenTTL}
}

func (s *Service) Login(ctx context.Context, username, password string) (token string, err error) {
	s.logger.Info("attempting to login")

	user, err := s.storage.GetUser(ctx, username)
	if err != nil {
		s.logger.Warn("failed to get user", slog.String("error", err.Error()))
		return "", fmt.Errorf("%s", err)
	}

	//Проверка пароля
	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(password)); err != nil {
		s.logger.Warn("invalid credentials", slog.String("error", err.Error()))
		return "", err
	}

	token, err = jwt.NewToken(user, s.tokenTTL)
	if err != nil {
		s.logger.Warn("failed to generate token", slog.String("error", err.Error()))

		return "", fmt.Errorf("%s", err)
	}

	return token, nil
}
