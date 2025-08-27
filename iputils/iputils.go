package iputils

import "net"

// Broadcast calculates the broadcast address for a given IP network.
func Broadcast(ipnet *net.IPNet) net.IP {
	ip := ipnet.IP.To4()
	if ip == nil {
		return nil // Not an IPv4 address
	}

	mask := ipnet.Mask
	broadcast := make(net.IP, len(ip))
	for i := 0; i < len(ip); i++ {
		broadcast[i] = ip[i] | ^mask[i]
	}
	return broadcast
}
