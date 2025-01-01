package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

const configPath = "/etc/ftp-auth-handler.yaml"

var (
	cmdGetUserAccounts bool
)

func main() {
	flag.BoolVar(&cmdGetUserAccounts, "get-user-accounts", false, "fetch user accounts from remote database")
	flag.Parse()

	// read configuration file
	cfg, err := readConfig(configPath)
	if err != nil {
		log.Fatal(err)
	}

	// write logs to file
	logFD, err := os.OpenFile(cfg.Log.Path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer logFD.Close()

	log.SetOutput(logFD)

	if cmdGetUserAccounts {
		accounts, err := getAccountsConfigFromRemoteDB(cfg.Database.User, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.Name, cfg.Table.Name)
		if err != nil {
			log.Fatal(err)
		}

		if len(accounts) > 0 {
			err = writeLocalDB(accounts, cfg.Accounts.Path)
			if err != nil {
				log.Fatal(err)
			}
		}
	} else {
		payload, err := genAuthenticationPayload(cfg.Accounts.Path)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(strings.Join(payload, "\n"))
	}

	os.Exit(0)
}
