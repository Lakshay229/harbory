package domain

import (
	"context"
	"io"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/image"
)

type ImageService interface {
	List(ctx context.Context, all bool) ([]image.Summary, error)
	Inspect(ctx context.Context, imageID string) (types.ImageInspect, error)
	Delete(ctx context.Context, imageID string) error
	Pull(ctx context.Context, imageName string, writer io.Writer) error
}

type ImageRepository interface {
	List(ctx context.Context, options image.ListOptions) ([]image.Summary, error)
	Inspect(ctx context.Context, imageID string) (types.ImageInspect, error)
	Delete(ctx context.Context, imageID string, options image.RemoveOptions) error
	Pull(ctx context.Context, imageName string, options image.PullOptions) (io.ReadCloser, error)
}
