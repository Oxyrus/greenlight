package data

import (
	"encoding/json"
	"fmt"
	"greenlight/internal/validator"
	"time"
)

type Movie struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	Title     string    `json:"title"`
	Year      int32     `json:"year,omitempty"`
	Runtime   Runtime   `json:"runtime,omitempty"`
	Genres    []string  `json:"genres,omitempty"`
	Version   int32     `json:"version"`
}

// MarshalJSON is used so it satisfies the json.Marshaler interface
func (m Movie) MarshalJSON() ([]byte, error) {
	// Declare a variable to hold the custom runtime string (empty by default)
	var runtime string

	// If the value of the runtime field is not zero, set the runtime variable to be
	// a string in the format "<runtime> mins"
	if m.Runtime != 0 {
		runtime = fmt.Sprintf("%d mins", m.Runtime)
	}

	// Create an anonymous struct to hold the data for JSON encoding. This has exactly
	// the same fields, types and tags as our Movie struct, except that the Runtime
	// field here is a string, instead of an int32. Also notice that we don't include
	// a CreatedAt field at all (there's no point including one, because we don't want
	// it to appear in the JSON output).
	aux := struct {
		ID      int64    `json:"id"`
		Title   string   `json:"title"`
		Year    int32    `json:"year,omitempty"`
		Runtime string   `json:"runtime,omitempty"`
		Genres  []string `json:"genres,omitempty"`
		Version int32    `json:"version"`
	}{
		ID:      m.ID,
		Title:   m.Title,
		Year:    m.Year,
		Runtime: runtime,
		Genres:  m.Genres,
		Version: m.Version,
	}

	return json.Marshal(aux)
}

func ValidateMovie(v *validator.Validator, movie *Movie) {
	// Use the Check() method to execute our validation checks.

	// Title validations
	v.Check(movie.Title != "", "title", "must be provided")
	v.Check(len(movie.Title) <= 500, "title", "must not be more than 500 bytes long")

	// Year validations
	v.Check(movie.Year != 0, "year", "must be provided")
	v.Check(movie.Year >= 1888, "year", "must be greater than 1888")
	v.Check(movie.Year <= int32(time.Now().Year()), "year", "must not be in the future")

	// Runtime validations
	v.Check(movie.Runtime != 0, "runtime", "must be provided")
	v.Check(movie.Runtime > 0, "runtime", "must be provided")

	// Genres validations
	v.Check(movie.Genres != nil, "genres", "must be provided")
	v.Check(len(movie.Genres) >= 1, "genres", "must contain at least 1 genre")
	v.Check(len(movie.Genres) <= 5, "genres", "must not contain more than 5 genres")
	v.Check(validator.Unique(movie.Genres), "genres", "must not contain duplicate values")
}
