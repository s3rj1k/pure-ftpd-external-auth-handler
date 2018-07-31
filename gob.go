package main

import (
	"encoding/gob"
	"fmt"
	"os"
	"path/filepath"
)

func writeLocalDB(data map[string]accountConfig, path string) error {

	var err error

	f, err := os.Create(filepath.Clean(path))
	if err != nil {
		return fmt.Errorf("failed to create local DB: %s", err.Error())
	}
	defer f.Close()

	enc := gob.NewEncoder(f)

	err = enc.Encode(data)
	if err != nil {
		return fmt.Errorf("failed to write data to local DB: %s", err.Error())
	}

	return nil
}

func readLocalDB(path string) (map[string]accountConfig, error) {

	var err error
	data := &map[string]accountConfig{}

	f, err := os.Open(filepath.Clean(path))
	if err != nil {
		return map[string]accountConfig{}, fmt.Errorf("failed to open local DB: %s", err.Error())
	}
	defer f.Close()

	dec := gob.NewDecoder(f)

	err = dec.Decode(data)
	if err != nil {
		return map[string]accountConfig{}, fmt.Errorf("failed to read data from local DB: %s", err.Error())
	}

	return *data, nil
}
