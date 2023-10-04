package website

import (
	"context"
	"database/sql"
)

type PostgresSQLClassicRepository struct {
	db *sql.DB
}

func NewPostgresSQLClassicRepository(db *sql.DB) *PostgresSQLClassicRepository {
	return &PostgresSQLClassicRepository{
		db: db,
	}
}

func (r *PostgresSQLClassicRepository) Migrate(ctx context.Context) error {
	query := `
	CREATE TABLE IF NOT EXISTS websites(
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL UNIQUE,
		url TEXT NOT NULL,
		rank INT NOT NULL
	);
	`
	_, err := r.db.ExecContext(ctx, query)
	return err
}