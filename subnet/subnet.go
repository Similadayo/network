package subnet

import (
	"fmt"
	"math"
	"net"

	"github.com/similadayo/iputils"
)

// Calcualte compute subnet details for a given CIDR notation
func Calcualte(cidr string) (string, error) {
	_, netw, err := net.ParseCIDR(cidr)
	if err != nil {
		return "", fmt.Errorf("invalid CIDR: %w", err)
	}

	if netw.IP.To4() == nil {
		return "", fmt.Errorf("only IPv4 is supported")
	}

	masklen, _ := netw.Mask.Size()
	hostbits := 32 - masklen
	usableHosts := int(math.Pow(2, float64(hostbits))) - 2
	if usableHosts < 0 {
		usableHosts = 0
	}

	broadcast := iputils.Broadcast(netw)
	firstUseable := iputils.IncrementIP(netw.IP, 1)
	lastUsable := iputils.IncrementIP(broadcast, ^uint64(0))

	usableRange := fmt.Sprintf("%s - %s", firstUseable, lastUsable)
	result := fmt.Sprintf("Network: %s\nNetmask: %s\nBroadcast: %s\nUsable Hosts: %d\nUsable Range: %s",
		netw.IP, netw.Mask, broadcast, usableHosts, usableRange)

	return result, nil
}
