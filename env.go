package main

import (
	"fmt"
	"net"
	"os"
)

type AuthorizationData struct {
	UserName string // AUTHD_ACCOUNT
	Password string // AUTHD_PASSWORD
	RemoteIP net.IP // AUTHD_REMOTE_IP
}

func getPureFTPdAuthData() (AuthorizationData, error) {
	var (
		auth AuthorizationData
		ok   bool
	)

	auth.UserName, ok = os.LookupEnv("AUTHD_ACCOUNT")
	if !ok {
		return AuthorizationData{}, fmt.Errorf("failed to lookup AUTHD_ACCOUNT")
	}

	auth.Password, ok = os.LookupEnv("AUTHD_PASSWORD")
	if !ok {
		return AuthorizationData{}, fmt.Errorf("failed to lookup AUTHD_PASSWORD")
	}

	ip, ok := os.LookupEnv("AUTHD_REMOTE_IP")
	if !ok {
		return AuthorizationData{}, fmt.Errorf("failed to lookup AUTHD_REMOTE_IP")
	}

	auth.RemoteIP = net.ParseIP(ip)
	if auth.RemoteIP == nil {
		return AuthorizationData{}, fmt.Errorf("failed to parse RemoteIP")
	}

	return auth, nil
}
