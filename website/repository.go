package website

import "context"

type Repository interface {
	Migrate(ctx context.Context) error
	Create(ctx context.Context, website Website) (*Website, error)
	All(ctx context.Context) ([]Website, error)
	GetByName(ctx context.Context, name string) (*Website, error)
	Update(ctx context.Context, id int64, updated Website) (*Website, error)
	Delete(ctx context.Context, id int64) error
}