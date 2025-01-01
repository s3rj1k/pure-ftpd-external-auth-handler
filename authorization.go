package main

import (
	"fmt"
	"os"
	"strings"
)

func isAuthorized(auth AuthorizationData, accounts map[string]AccountConfig) (bool, error) {
	account, ok := accounts[auth.UserName]
	if !ok {
		return false, fmt.Errorf("account %q does not exist", auth.UserName)
	}

	ok = checkPasswordHash(auth.Password, account.Password)
	if !ok {
		return false, fmt.Errorf("password for %q is wrong", auth.UserName)
	}

	// if list is empty, skip check
	if len(strings.TrimSpace(account.AuthorizedClientIPs)) != 0 {
		// is authorized remote IP
		if !isIPInList(auth.RemoteIP, account.AuthorizedClientIPs) {
			return false, fmt.Errorf("remote IP %q for %q not in authorized list", auth.RemoteIP, auth.UserName)
		}
	}

	// if list is empty, skip check
	if len(strings.TrimSpace(account.RefuzedClientIPs)) != 0 {
		// is denied remote IP
		if isIPInList(auth.RemoteIP, account.RefuzedClientIPs) {
			return false, fmt.Errorf("remote IP %q for %q in denied list", auth.RemoteIP, auth.UserName)
		}
	}

	// user UID/GID must be greater that 1000
	if account.UID < 1000 || account.GID < 1000 {
		return false, fmt.Errorf("UID/GID for %q is lesser than 1000", auth.UserName)
	}

	fi, err := os.Lstat(account.HomeDirectory)
	if err != nil {
		return false, err
	}

	switch mode := fi.Mode(); {
	// deny if regular file
	case mode.IsRegular():
		return false, fmt.Errorf("home dir for %q is file", auth.UserName)
	// deny if symlink
	case mode&os.ModeSymlink != 0:
		return false, fmt.Errorf("home dir for %q is symlink", auth.UserName)
	// deny if named pipe
	case mode&os.ModeNamedPipe != 0:
		return false, fmt.Errorf("home dir for %q is named pipe", auth.UserName)
	// deny if device
	case mode&os.ModeDevice != 0:
		return false, fmt.Errorf("home dir for %q is device", auth.UserName)
	// deny if socket
	case mode&os.ModeSocket != 0:
		return false, fmt.Errorf("home dir for %q is socket", auth.UserName)
	}

	// home_dir must be in allowed path prefixes list
	allowedPathPrefixes := []string{"/home/", "/tmp/", "/media/", "/mnt/", "/storage/"}
	for _, prefix := range allowedPathPrefixes {
		if strings.HasPrefix(account.HomeDirectory, prefix) {
			return true, nil
		}
	}

	return false, fmt.Errorf("%q not authorized", auth.UserName)
}

func genAuthenticationPayload(path string) ([]string, error) {
	out := make([]string, 0)

	auth, err := getPureFTPdAuthData()
	if err != nil {
		return []string{"auth_ok:-1", "end"}, err
	}

	accounts, err := readLocalDB(path)
	if err != nil {
		return []string{"auth_ok:-1", "end"}, err
	}

	ok, err := isAuthorized(auth, accounts)
	if err != nil {
		return []string{"auth_ok:-1", "end"}, err
	}

	if !ok {
		return []string{"auth_ok:-1", "end"}, nil
	}

	out = append(out, "auth_ok:1",
		fmt.Sprintf("uid:%d", accounts[auth.UserName].UID),
		fmt.Sprintf("gid:%d", accounts[auth.UserName].GID),
		fmt.Sprintf("dir:%s", accounts[auth.UserName].HomeDirectory),
		"slow_tilde_expansion:1",
	)

	if accounts[auth.UserName].UploadBandwidth != 0 {
		out = append(out, fmt.Sprintf("throttling_bandwidth_ul:%d", accounts[auth.UserName].UploadBandwidth))
	}

	if accounts[auth.UserName].DownloadBandwidth != 0 {
		out = append(out, fmt.Sprintf("throttling_bandwidth_dl:%d", accounts[auth.UserName].DownloadBandwidth))
	}

	if accounts[auth.UserName].MaxNumberOfConnections != 0 {
		out = append(out, fmt.Sprintf("per_user_max:%d", accounts[auth.UserName].MaxNumberOfConnections))
	}

	if accounts[auth.UserName].SizeQuota != 0 {
		out = append(out, fmt.Sprintf("user_quota_size:%d", accounts[auth.UserName].SizeQuota))
	}

	if accounts[auth.UserName].FilesQuota != 0 {
		out = append(out, fmt.Sprintf("user_quota_files:%d", accounts[auth.UserName].FilesQuota))
	}

	out = append(out, "end")

	return out, nil
}
