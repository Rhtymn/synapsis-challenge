package domain

import "context"

type Category struct {
	ID       int64
	ImageURL *string
	Name     string
	Slug     string
	ParentID *int64
}

type CategoryRepository interface {
	GetBySlug(ctx context.Context, slug string) (Category, error)
}
