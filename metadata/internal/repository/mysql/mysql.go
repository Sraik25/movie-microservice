package mysql

import (
	"context"
	"database/sql"
	"errors"
	"movieexample.com/metadata/internal/repository"
	"movieexample.com/metadata/pkg/model"

	_ "github.com/go-sql-driver/mysql"
)

type Repository struct {
	db *sql.DB
}

func NewRepository() (*Repository, error) {
	db, err := sql.Open("mysql", "root:password@/movieexample")
	if err != nil {
		return nil, err
	}
	return &Repository{db}, nil
}

// Get retrieves movie metadata for by movie id.
func (r *Repository) Get(ctx context.Context, id string) (*model.Metadata, error) {
	var title, description, director string
	query := `SELECT title, description, director FROM movies WHERE id = ?`
	row := r.db.QueryRowContext(ctx, query, id)

	if err := row.Scan(&title, &description, &director); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}

	return &model.Metadata{
		ID:          id,
		Title:       title,
		Description: director,
		Director:    director,
	}, nil
}

// Put adds movie metadata for a given movie id.
func (r *Repository) Put(ctx context.Context, id string, metadata *model.Metadata) error {
	query := `INSERT INTO movies (id, title, description, director) VALUES (?, ?, ?, ?)`
	_, err := r.db.ExecContext(ctx, query, id, metadata.Title, metadata.Description, metadata.Director)
	return err
}
