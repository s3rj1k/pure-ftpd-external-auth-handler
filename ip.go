package main

import (
	"net"
	"strings"
)

func isIPInList(remoteIP net.IP, list string) bool {

	for _, value := range strings.Split(list, ",") {

		if net.ParseIP(value).Equal(remoteIP) {
			return true
		}

		ip, network, err := net.ParseCIDR(value)
		if err == nil {
			if ip.Equal(remoteIP) {
				return true
			}

			if network.Contains(remoteIP) {
				return true
			}
		}

		addrs, err := net.LookupIP(value)
		if err == nil {
			for _, ip := range addrs {
				if ip.Equal(remoteIP) {
					return true
				}
			}
		}

	}

	return false
}
