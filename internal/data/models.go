package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
)

// Models wraps the MovieModel, eventually it will hold other models such as
// UserModel, PermissionModel, etc.
type Models struct {
	Movies MovieModel
}

// NewModels returns a Models struct containing the initialized MovieModel for easy of use.
func NewModels(db *sql.DB) Models {
	return Models{
		Movies: MovieModel{DB: db},
	}
}
