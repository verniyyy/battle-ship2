package lib

import (
	"net"
)

func IPAddr() (string, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}

	var ipAddrs []string
	const LOCALHOST = "127.0.0.1"
	for _, inter := range interfaces {
		addrs, err := inter.Addrs()
		if err != nil {
			return "", err
		}
		for _, addr := range addrs {
			if ipnet, ok := addr.(*net.IPNet); ok {
				if ipnet.IP.To4() != nil && ipnet.IP.String() != LOCALHOST {
					ipAddrs = append(ipAddrs, ipnet.IP.String())
				}
			}
		}
	}

	if len(ipAddrs) > 0 {
		return ipAddrs[0], nil
	} else {
		return "", err
	}
}
