package IsGoogle

import (
	"net"
	"strings"
)

// IsGoogle should check if a certain address belongs to Google
func IsGoogle(addr string) (bool, error) {

	addrs, err := net.LookupAddr(addr)
	if err != nil {
		return false, err
	}

	if !(strings.HasSuffix(addrs[0], ".google.com.") || strings.HasSuffix(addrs[0], ".googlebot.com.")) {
		return false, nil
	}

	ips, err := net.LookupIP(addrs[0])
	if err != nil {
		return false, err
	}

	return ips[0].String() == addr, nil
}
