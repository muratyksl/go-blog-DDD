package repository

import (
	"app/internal/post/domain"
	"context"
)

type PostRepository interface {
	GetByID(ctx context.Context, id int) (*domain.Post, error)
	GetAll(ctx context.Context) ([]*domain.Post, error)
	Create(ctx context.Context, post *domain.Post) error
	Delete(ctx context.Context, ids []int) error
}
