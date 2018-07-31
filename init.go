package main

import (
	"flag"
)

const configPath = "/etc/ftp-auth-handler.yaml"

var (
	cmdGetUserAccounts bool
)

func init() {
	flag.BoolVar(&cmdGetUserAccounts, "get-user-accounts", false, "fetch user accounts from remote database")
	flag.Parse()
}
