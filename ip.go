package util

import (
	"net"
	"strings"
)

func IPAddress() (ipAddr string, err error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return
	}

	for _, i := range ifaces {
		var addrs []net.Addr
		addrs, err = i.Addrs()
		if err != nil {
			return
		}

		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			// process IP address
			ipAddr = ip.String()
			if strings.Index(ipAddr, ".") > 0 && !strings.HasPrefix(ipAddr, "127") {
				return ipAddr, nil
			}
		}
	}
	return
}
