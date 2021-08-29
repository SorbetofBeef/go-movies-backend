package models

import (
	"database/sql"
	"time"
)

// Models is a wrapper the database
type Models struct {
	DB DBModel
}

// NewModels returns a new database model
func NewModels(db *sql.DB) Models {
	return Models{
		DB: DBModel{DB: db},
	}
}

// Movie defines how the movies table looks in json for the PostgreSQL db
type Movie struct {
	ID          int          `json:"id"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Year        int          `json:"year"`
	ReleaseDate time.Time    `json:"release_date"`
	Runtime     int          `json:"runtime"`
	Rating      int          `json:"rating"`
	MPAARating  string       `json:"mpaa_rating"`
	CreatedAt   time.Time    `json:"-"`
	UpdatedAt   time.Time    `json:"-"`
	MovieGenre  map[int]string `json:"movie_genre"`
}

// Genre defines how the genre table looks in json for the PostgreSQL db
type Genre struct {
	ID        int       `json:"id"`
	GenreName string    `json:"genre_name"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

// MovieGenre defines how the movie_genre table looks in json for the PostgreSQL db
type MovieGenre struct {
	ID        int       `json:"-"`
	MovieId   int       `json:"-"`
	GenreId   int       `json:"-"`
	Genre     Genre     `json:"genre"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
