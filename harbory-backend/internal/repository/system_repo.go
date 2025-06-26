package repository

import (
	"context"

	"github.com/docker/docker/api/types/system"
	"github.com/docker/docker/client"
)

type systemRepository struct {
	client *client.Client
}

func NewSystemRepository(client *client.Client) *systemRepository {
	return &systemRepository{client: client}
}
func (r *systemRepository) GetInfo(ctx context.Context) (system.Info, error) {
	return r.client.Info(ctx)
}