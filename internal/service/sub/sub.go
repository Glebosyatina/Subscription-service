package sub

import (
	"log/slog"
	"strconv"

	"glebosyatina/test_project/internal/domain"
)

type SubRepository interface {
	CreateSub(userId uint64, nameService string, price uint64, startDate string, endDate string) (*domain.Sub, error)
	GetSubByID(idSub uint64) (*domain.Sub, error)
	DeleteSubByID(idSub uint64) error
}

type SubService struct {
	subRepo SubRepository
	lg      *slog.Logger
}

func NewSubService(sr SubRepository, logger *slog.Logger) *SubService {
	return &SubService{
		subRepo: sr,
		lg:      logger,
	}
}

func (s *SubService) AddSub(userId uint64, nameService string, price uint64, start string, end string) (*domain.Sub, error) {
	sub, err := s.subRepo.CreateSub(userId, nameService, price, start, end)
	if err != nil {
		return nil, err
	}

	s.lg.Info("Subscription created:", slog.String("NameService", nameService), slog.String("Price", strconv.Itoa(int(price))))
	return sub, nil
}

func (s *SubService) GetSubscription(subId uint64) (*domain.Sub, error) {
	sub, err := s.subRepo.GetSubByID(subId)
	if err != nil {
		return nil, err
	}
	return sub, nil
}

func (s *SubService) DeleteSubByID(subId uint64) error {
	err := s.subRepo.DeleteSubByID(subId)
	if err != nil {
		return err
	}
	return nil
}
