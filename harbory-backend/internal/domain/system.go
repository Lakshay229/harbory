package domain

import (
	"context"

	"github.com/docker/docker/api/types/system"
)

type SystemService interface {
	GetInfo(ctx context.Context) (system.Info, error)
}

type SystemRepository interface {
	GetInfo(ctx context.Context) (system.Info, error)
}
