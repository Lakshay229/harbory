package service

import (
	"context"

	"harbory-backend/internal/domain"

	"github.com/docker/docker/api/types/system"
)

type systemService struct {
	repo domain.SystemRepository
}

func NewSystemService(repo domain.SystemRepository) domain.SystemService {
	return &systemService{repo: repo}
}
func (s *systemService) GetInfo(ctx context.Context) (system.Info, error) {
	return s.repo.GetInfo(ctx)
}
