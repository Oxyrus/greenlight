package data

import (
	"database/sql"
	"errors"
)

// We'll return this from our Get() method when
// looking up a record that doesn't exist in our database.
var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
)

// Create a Models struct which wraps the MovieModel.
// Eventually we'll add more models to this.
type Models struct {
	Movies interface {
		Insert(movie *Movie) error
		Get(id int64) (*Movie, error)
		Update(movie *Movie) error
		Delete(id int64) error
	}
}

// For ease of use, we add a NewModels() method to return a struct
// containing all the initialized models.
func NewModel(db *sql.DB) Models {
	return Models{
		Movies: MovieModel{DB: db},
	}
}
