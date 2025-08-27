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

// IncrementIP increments an IP address by a given value.
func IncrementIP(ip net.IP, increment int) net.IP {
	ip = ip.To4()
	if ip == nil {
		return nil // Not an IPv4 address
	}

	val := uint64(ip[0])<<24 | uint64(ip[1])<<16 | uint64(ip[2])<<8 | uint64(ip[3])
	val += uint64(increment)
	newIP := make(net.IP, 4)
	newIP[0] = byte((val >> 24) & 0xFF)
	newIP[1] = byte((val >> 16) & 0xFF)
	newIP[2] = byte((val >> 8) & 0xFF)
	newIP[3] = byte(val & 0xFF)
	return newIP
}
