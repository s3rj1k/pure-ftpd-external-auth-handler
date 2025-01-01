package main

import (
	"database/sql"
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"github.com/go-sql-driver/mysql"
)

type AccountConfig struct {
	UserName               string
	Password               string
	UID                    int64
	GID                    int64
	HomeDirectory          string
	UploadBandwidth        int64
	DownloadBandwidth      int64
	MaxNumberOfConnections int64
	FilesQuota             int64
	SizeQuota              int64
	AuthorizedClientIPs    string
	RefuzedClientIPs       string
}

const sqlQuery = `
SELECT 
	UserName, Password, UID, GID, HomeDirectory,
	UploadBandwidth, DownloadBandwidth, MaxNumberOfConnections,
	FilesQuota, SizeQuota,
	AuthorizedClientIPs, RefuzedClientIPs
FROM`

func getAccountsConfigFromRemoteDB(user, password, host string, _ int /*port*/, database, table string) (map[string]AccountConfig, error) {
	config := mysql.NewConfig()

	config.User = user
	config.Passwd = password
	config.Net = "tcp"
	config.Addr = host
	config.DBName = database

	config.Timeout = 30 * time.Second
	config.ReadTimeout = 60 * time.Second
	config.WriteTimeout = 60 * time.Second

	db, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		return map[string]AccountConfig{}, fmt.Errorf("failed to connect to remote DB: %w", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		return map[string]AccountConfig{}, fmt.Errorf("failed to ping remote DB server: %w", err)
	}

	results, err := db.Query(fmt.Sprintf("%s %s", sqlQuery, table))
	if err != nil {
		return map[string]AccountConfig{}, fmt.Errorf("failed to query remote DB server: %w", err)
	}

	records := make(map[string]AccountConfig)

	for results.Next() {
		var (
			record AccountConfig
			err    error
		)

		err = results.Scan(
			&record.UserName,
			&record.Password,
			&record.UID,
			&record.GID,
			&record.HomeDirectory,
			&record.UploadBandwidth,
			&record.DownloadBandwidth,
			&record.MaxNumberOfConnections,
			&record.FilesQuota,
			&record.SizeQuota,
			&record.AuthorizedClientIPs,
			&record.RefuzedClientIPs)
		if err != nil {
			return map[string]AccountConfig{}, fmt.Errorf("failed to parse returned table raw from remote DB server: %w", err)
		}

		// sanitizing
		record.UserName = strings.TrimSpace(strings.TrimSuffix(record.UserName, "\n"))
		record.Password = strings.TrimSpace(strings.TrimSuffix(record.Password, "\n"))
		record.AuthorizedClientIPs = strings.TrimSpace(strings.TrimSuffix(record.AuthorizedClientIPs, "\n"))
		record.RefuzedClientIPs = strings.TrimSpace(strings.TrimSuffix(record.RefuzedClientIPs, "\n"))
		record.HomeDirectory = filepath.Clean(record.HomeDirectory)

		// minimal config validation
		if record.GID == 0 || record.UID == 0 || len(record.Password) < 4 || len(record.UserName) < 2 || record.HomeDirectory == "/" {
			continue
		}

		record.Password, err = hashPassword(record.Password)
		if err != nil {
			return map[string]AccountConfig{}, fmt.Errorf("failed to hash account password: %w", err)
		}

		records[record.UserName] = record
	}

	return records, nil
}
