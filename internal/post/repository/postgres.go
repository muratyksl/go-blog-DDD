package repository

import (
	"app/internal/post/domain"
	"context"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v4/pgxpool"
)

type postgresRepository struct {
	db *pgxpool.Pool
}

func NewPostgresRepository(db *pgxpool.Pool) PostRepository {
	return &postgresRepository{db: db}
}

func (r *postgresRepository) GetByID(ctx context.Context, id int) (*domain.Post, error) {
	var post domain.Post
	err := r.db.QueryRow(ctx, "SELECT id, title, body FROM posts WHERE id = $1", id).Scan(&post.ID, &post.Title, &post.Body)
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *postgresRepository) GetAll(ctx context.Context) ([]*domain.Post, error) {
	rows, err := r.db.Query(ctx, "SELECT id, title, body FROM posts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []*domain.Post
	for rows.Next() {
		var post domain.Post
		err := rows.Scan(&post.ID, &post.Title, &post.Body)
		if err != nil {
			return nil, err
		}
		posts = append(posts, &post)
	}
	return posts, nil
}

func (r *postgresRepository) Create(ctx context.Context, post *domain.Post) error {
	_, err := r.db.Exec(ctx, "INSERT INTO posts (title, body) VALUES ($1, $2) RETURNING id", post.Title, post.Body)
	return err
}

func (r *postgresRepository) Delete(ctx context.Context, ids []int) error {
	// Create a string with placeholders for each ID
	placeholders := make([]string, len(ids))
	for i := range ids {
		placeholders[i] = fmt.Sprintf("$%d", i+1)
	}

	query := fmt.Sprintf("DELETE FROM posts WHERE id IN (%s)", strings.Join(placeholders, ","))

	// Convert []int to []interface{} for the Exec function
	args := make([]interface{}, len(ids))
	for i, id := range ids {
		args[i] = id
	}

	_, err := r.db.Exec(ctx, query, args...)
	return err
}
