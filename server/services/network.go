package services

import "net"

type NetworkService struct {
}

func (receiver NetworkService) GetAddresses() ([]string, error) {
	ifaces, err := net.Interfaces()

	if err != nil {
		return nil, err
	}

	var addresses []string

	for _, i := range ifaces {
		addrs, err := i.Addrs()

		if err != nil {
			return nil, err
		}

		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			addresses = append(addresses, ip.String())
		}
	}

	return addresses, err
}
