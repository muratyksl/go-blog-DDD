package service

import (
	"app/internal/common/errors"
	"app/internal/post/domain"
	"app/internal/post/repository"
	"context"

	"go.uber.org/zap"
)

type PostService interface {
	GetPost(ctx context.Context, id int) (*domain.Post, error)
	GetAllPosts(ctx context.Context) ([]*domain.Post, error)
	CreatePost(ctx context.Context, post *domain.Post) error
	DeletePosts(ctx context.Context, ids []int) error
}

type postService struct {
	repo   repository.PostRepository
	logger *zap.Logger
}

func NewPostService(repo repository.PostRepository, logger *zap.Logger) PostService {
	return &postService{repo: repo, logger: logger}
}

func (s *postService) GetPost(ctx context.Context, id int) (*domain.Post, error) {
	post, err := s.repo.GetByID(ctx, id)
	if err != nil {
		s.logger.Error("Failed to get post", zap.Int("id", id), zap.Error(err))
		return nil, errors.NewAppError("NOT_FOUND", "Post not found", err)
	}
	return post, nil
}

func (s *postService) GetAllPosts(ctx context.Context) ([]*domain.Post, error) {
	posts, err := s.repo.GetAll(ctx)
	if err != nil {
		s.logger.Error("Failed to get all posts", zap.Error(err))
		return nil, errors.NewAppError("INTERNAL_ERROR", "Failed to retrieve posts", err)
	}
	return posts, nil
}

func (s *postService) CreatePost(ctx context.Context, post *domain.Post) error {
	err := s.repo.Create(ctx, post)
	if err != nil {
		s.logger.Error("Failed to create post", zap.Error(err))
		return errors.NewAppError("INTERNAL_ERROR", "Failed to create post", err)
	}
	return nil
}

func (s *postService) DeletePosts(ctx context.Context, ids []int) error {
	err := s.repo.Delete(ctx, ids)
	if err != nil {
		s.logger.Error("Failed to delete posts", zap.Ints("ids", ids), zap.Error(err))
		return errors.NewAppError("INTERNAL_ERROR", "Failed to delete posts", err)
	}
	return nil
}
