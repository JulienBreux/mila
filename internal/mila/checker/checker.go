package checker

import (
	"errors"
	"fmt"
	"os"

	"github.com/julienbreux/mila/internal/mila"
)

var (
	// ErrNoSchema if there is no schema to check
	ErrNoSchema = errors.New("no schema")

	// ErrNoVars if there are no variables
	ErrNoVars = errors.New("no variable")
)

// Checker checks the schema according to the environment
func Checker(sc *mila.Schema) error {
	if sc == nil {
		return ErrNoSchema
	}

	if err := required(&sc.Vars); err != nil {
		return err
	}

	if err := validated(&sc.Vars); err != nil {
		return err
	}

	return nil
}

func required(vars *mila.Vars) error {
	if vars == nil {
		return ErrNoVars
	}

	for name, variable := range *vars {
		_, ok := os.LookupEnv(name)
		if variable.Required && !ok {
			return fmt.Errorf("variable %s not found", name)
		}
	}

	return nil
}

func validated(vars *mila.Vars) error {
	if vars == nil {
		return ErrNoVars
	}

	// TODO: Write validated core logic

	return nil
}
