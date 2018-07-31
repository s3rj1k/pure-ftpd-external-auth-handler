package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	yaml "gopkg.in/yaml.v2"
)

type appConfig struct {
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

func readConfig(file string) (appConfig, error) {

	var cfg appConfig

	f, err := ioutil.ReadFile(filepath.Clean(file))
	if err != nil {
		return appConfig{}, fmt.Errorf("failed to read application config: %s", err.Error())
	}

	err = yaml.Unmarshal(f, &cfg)
	if err != nil {
		return appConfig{}, fmt.Errorf("failed to parse application config: %s", err.Error())
	}

	// Below default values for config

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
