package models

import (
	"context"
	"database/sql"
	"time"
)

// DBModel is a wrapper for a database models
type DBModel struct {
	DB *sql.DB
}

// Get returns a single movie
func (m *DBModel) Get(id int) (*Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `
		select
			id, title, description, year, release_date, runtime, rating, mpaa_rating,
				created_at, updated_at
		from
			movies
		where
			id = $1
	`

	row := m.DB.QueryRowContext(ctx, query, id)
	var movie Movie
	err := row.Scan(
		&movie.ID,
		&movie.Title,
		&movie.Description,
		&movie.Year,
		&movie.ReleaseDate,
		&movie.Runtime,
		&movie.Rating,
		&movie.MPAARating,
		&movie.CreatedAt,
		&movie.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &movie, nil
}

// All returns all movies
func (m *DBModel) All(id int) ([]*Movie, error) {
	return nil, nil
}
