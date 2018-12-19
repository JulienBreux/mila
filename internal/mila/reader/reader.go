package reader

import (
	"fmt"
	"os"
	"path/filepath"

	yaml "gopkg.in/yaml.v2"

	"github.com/julienbreux/mila/internal/mila"
)

const (
	defaultSchemaFilename = ".mila.yaml"
)

// Reader reads a Mila schema for a given directory
func Reader(dir string) (*mila.Schema, error) {
	path := filepath.Join(dir, defaultSchemaFilename)
	if _, err := os.Stat(path); err != nil {
		return nil, fmt.Errorf("no %s file found in %s", defaultSchemaFilename, dir)
	}

	t, err := readSchemaFile(path)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func readSchemaFile(file string) (*mila.Schema, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	var s mila.Schema

	return &s, yaml.NewDecoder(f).Decode(&s)
}
