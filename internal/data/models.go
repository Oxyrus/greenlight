package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
)

// Models wraps the MovieModel, eventually it will hold other models such as
// UserModel, PermissionModel, etc.
type Models struct {
	Movies MovieModel
	Users  UserModel
}

// NewModels returns a Models struct containing the initialized MovieModel for easy of use.
func NewModels(db *sql.DB) Models {
	return Models{
		Movies: MovieModel{DB: db},
		Users:  UserModel{DB: db},
	}
}
