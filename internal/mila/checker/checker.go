package checker

import (
	"errors"

	"github.com/julienbreux/mila/internal/mila"
)

var (
	// ErrNoSchema if there is no schema to check
	ErrNoSchema = errors.New("no schema")
)

// Checker checks the schema according to the environment
func Checker(sc *mila.Schema) error {
	if sc == nil {
		return ErrNoSchema
	}

	// TODO: Do checker

	return nil
}
