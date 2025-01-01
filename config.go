package main

import (
	"fmt"
	"os"
	"path/filepath"

	yaml "gopkg.in/yaml.v3"
)

type AppConfig struct {
	Database struct {
		Name     string `yaml:"Name"`
		User     string `yaml:"User"`
		Host     string `yaml:"Host"`
		Password string `yaml:"Password"`
		Port     int    `yaml:"Port"`
	} `yaml:"Database"`
	Table struct {
		Name string `yaml:"Name"`
	} `yaml:"Table"`
	Accounts struct {
		Path string `yaml:"Path"`
	} `yaml:"Accounts"`
	Log struct {
		Path string `yaml:"Path"`
	} `yaml:"Log"`
}

func readConfig(file string) (AppConfig, error) {
	var cfg AppConfig

	f, err := os.ReadFile(filepath.Clean(file))
	if err != nil {
		return AppConfig{}, fmt.Errorf("failed to read application config: %w", err)
	}

	err = yaml.Unmarshal(f, &cfg)
	if err != nil {
		return AppConfig{}, fmt.Errorf("failed to parse application config: %w", err)
	}

	if len(cfg.Database.Name) == 0 {
		cfg.Database.Name = "ftp"
	}

	if cfg.Database.Port == 0 {
		cfg.Database.Port = 3306
	}

	if len(cfg.Table.Name) == 0 {
		cfg.Table.Name = "users"
	}

	if len(cfg.Accounts.Path) == 0 {
		cfg.Accounts.Path = "/dev/shm/accounts.db"
	}

	if len(cfg.Log.Path) == 0 {
		cfg.Log.Path = "/tmp/ftp-auth-handler.log"
	}

	return cfg, nil
}
