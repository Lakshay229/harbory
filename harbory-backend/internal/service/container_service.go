package service

import (
	"context"

	"harbory-backend/internal/domain"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
)

type containerService struct {
	repo domain.ContainerRepository
}

func NewContainerService(repo domain.ContainerRepository) domain.ContainerService {
	return &containerService{repo: repo}
}

func (s *containerService) List(ctx context.Context, all bool) ([]types.Container, error) {
	return s.repo.List(ctx, container.ListOptions{All: all})
}

func (s *containerService) Start(ctx context.Context, containerID string) error {
	return s.repo.Start(ctx, containerID, container.StartOptions{})
}

func (s *containerService) Stop(ctx context.Context, containerID string) error {
	return s.repo.Stop(ctx, containerID, container.StopOptions{})
}

func (s *containerService) Delete(ctx context.Context, containerID string) error {
	return s.repo.Delete(ctx, containerID, container.RemoveOptions{Force: true})
}

func (s *containerService) GetLogs(ctx context.Context, containerID string, options container.LogsOptions) (string, error) {
	return s.repo.GetLogs(ctx, containerID, options)
}
