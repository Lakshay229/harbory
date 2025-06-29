package service

import (
	"context"
	"io"

	"harbory-backend/internal/domain"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/image"
)

type imageService struct {
	repo domain.ImageRepository
}

func NewImageService(repo domain.ImageRepository) domain.ImageService {
	return &imageService{repo: repo}
}

func (s *imageService) List(ctx context.Context, all bool) ([]image.Summary, error) {
	return s.repo.List(ctx, image.ListOptions{All: all})
}

func (s *imageService) Inspect(ctx context.Context, imageID string) (types.ImageInspect, error) {
	return s.repo.Inspect(ctx, imageID)
}

func (s *imageService) Delete(ctx context.Context, imageID string) error {
	return s.repo.Delete(ctx, imageID, image.RemoveOptions{Force: true})
}

func (s *imageService) Pull(ctx context.Context, imageName string, writer io.Writer) error {
	resp, err := s.repo.Pull(ctx, imageName, image.PullOptions{})
	if err != nil {
		return err
	}
	defer resp.Close()

	_, err = io.Copy(writer, resp)
	return err
}
