package main

import (
	"encoding/gob"
	"fmt"
	"os"
	"path/filepath"
)

func writeLocalDB(data map[string]AccountConfig, path string) error {
	f, err := os.Create(filepath.Clean(path))
	if err != nil {
		return fmt.Errorf("failed to create local DB: %w", err)
	}
	defer f.Close()

	enc := gob.NewEncoder(f)

	err = enc.Encode(data)
	if err != nil {
		return fmt.Errorf("failed to write data to local DB: %w", err)
	}

	return nil
}

func readLocalDB(path string) (map[string]AccountConfig, error) {
	data := &map[string]AccountConfig{}

	f, err := os.Open(filepath.Clean(path))
	if err != nil {
		return nil, fmt.Errorf("failed to open local DB: %w", err)
	}
	defer f.Close()

	dec := gob.NewDecoder(f)

	err = dec.Decode(data)
	if err != nil {
		return nil, fmt.Errorf("failed to read data from local DB: %w", err)
	}

	return *data, nil
}
