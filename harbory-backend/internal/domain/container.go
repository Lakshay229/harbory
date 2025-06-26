package domain

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
)

type ContainerService interface {
	List(ctx context.Context, all bool) ([]types.Container, error)
	Start(ctx context.Context, containerID string) error
	Stop(ctx context.Context, containerID string) error
	Delete(ctx context.Context, containerID string) error
	GetLogs(ctx context.Context, containerID string, options container.LogsOptions) (string, error)
}

type ContainerRepository interface {
	List(ctx context.Context, options container.ListOptions) ([]types.Container, error)
	Start(ctx context.Context, containerID string, options container.StartOptions) error
	Stop(ctx context.Context, containerID string, timeout container.StopOptions) error
	Delete(ctx context.Context, containerID string, options container.RemoveOptions) error
	GetLogs(ctx context.Context, containerID string, options container.LogsOptions) (string, error)
}
