package repository

import (
	"context"
	"io"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
)

type imageRepository struct {
	client *client.Client
}

func NewImageRepository(client *client.Client) *imageRepository {
	return &imageRepository{client: client}
}

func (r *imageRepository) List(ctx context.Context, options image.ListOptions) ([]image.Summary, error) {
	return r.client.ImageList(ctx, options)
}

func (r *imageRepository) Inspect(ctx context.Context, imageID string) (types.ImageInspect, error) {
	inspect, _, err := r.client.ImageInspectWithRaw(ctx, imageID)
	return inspect, err
}

func (r *imageRepository) Delete(ctx context.Context, imageID string, options image.RemoveOptions) error {
	_, err := r.client.ImageRemove(ctx, imageID, options)
	return err
}

func (r *imageRepository) Pull(ctx context.Context, imageName string, options image.PullOptions) (io.ReadCloser, error) {
	return r.client.ImagePull(ctx, imageName, options)
}
