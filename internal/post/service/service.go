package service

import (
	"app/internal/post/domain"
	"app/internal/post/repository"
	"context"
)

type PostService interface {
	GetPost(ctx context.Context, id int) (*domain.Post, error)
	GetAllPosts(ctx context.Context) ([]*domain.Post, error)
	CreatePost(ctx context.Context, post *domain.Post) error
	DeletePosts(ctx context.Context, ids []int) error
}

type postService struct {
	repo repository.PostRepository
}

func NewPostService(repo repository.PostRepository) PostService {
	return &postService{repo: repo}
}

func (s *postService) GetPost(ctx context.Context, id int) (*domain.Post, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *postService) GetAllPosts(ctx context.Context) ([]*domain.Post, error) {
	return s.repo.GetAll(ctx)
}

func (s *postService) CreatePost(ctx context.Context, post *domain.Post) error {
	return s.repo.Create(ctx, post)
}

func (s *postService) DeletePosts(ctx context.Context, ids []int) error {
	return s.repo.Delete(ctx, ids)
}
