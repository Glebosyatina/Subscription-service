package user

import (
	"log/slog"

	"glebosyatina/test_project/internal/domain"
)

// интерфейсы репозиторного слоя
type UserRepository interface {
	CreateUser(name string, surname string) (*domain.User, error)
	GetUserById(id uint64) (*domain.User, error)
	DeleteUserById(id uint64) error
}

// юзер сервис
type UserService struct {
	userRepo UserRepository
	lg       *slog.Logger
}

func NewUserService(ur UserRepository, logger *slog.Logger) *UserService {
	return &UserService{
		userRepo: ur,
		lg:       logger,
	}
}

func (s *UserService) AddUser(name string, surname string) (*domain.User, error) {
	user, err := s.userRepo.CreateUser(name, surname)
	if err != nil {
		s.lg.Error("Не удалось создать пользователся в бд", err)
		return nil, err
	}

	s.lg.Info("User created:", slog.String("name", name), slog.String("surname", surname))
	return user, nil
}
