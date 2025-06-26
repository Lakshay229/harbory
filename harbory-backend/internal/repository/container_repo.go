package repository

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

type containerRepository struct {
	client *client.Client
}

func NewContainerRepository(client *client.Client) *containerRepository {
	return &containerRepository{client: client}
}

func (r *containerRepository) List(ctx context.Context, options container.ListOptions) ([]types.Container, error) {
	return r.client.ContainerList(ctx, options)
}

func (r *containerRepository) Start(ctx context.Context, containerID string, options container.StartOptions) error {
	return r.client.ContainerStart(ctx, containerID, options)
}

func (r *containerRepository) Stop(ctx context.Context, containerID string, timeout container.StopOptions) error {
	return r.client.ContainerStop(ctx, containerID, timeout)
}

func (r *containerRepository) Delete(ctx context.Context, containerID string, options container.RemoveOptions) error {
	return r.client.ContainerRemove(ctx, containerID, options)
}

func (r *containerRepository) GetLogs(ctx context.Context, containerID string, options container.LogsOptions) (string, error) {
	logs, err := r.client.ContainerLogs(ctx, containerID, options)
	if err != nil {
		return "", err
	}
	defer logs.Close()

	buf := new([]byte)
	_, err = logs.Read(*buf)
	if err != nil {
		return "", err
	}

	return string(*buf), nil
}
