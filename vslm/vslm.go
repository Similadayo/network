// File: vlsm/vlsm.go
package vlsm

import (
	"fmt"
	"math"
	"net"
	"sort"
	"strconv"

	"github.com/similadayo/iputils"
)

// Calculate performs VLSM allocation for a base CIDR and host counts.
func Calculate(baseCIDR string, hostCounts []string) ([]string, error) {
	_, baseNet, err := net.ParseCIDR(baseCIDR)
	if err != nil {
		return nil, fmt.Errorf("invalid base CIDR: %w", err)
	}

	if baseNet.IP.To4() == nil {
		return nil, fmt.Errorf("only IPv4 supported")
	}

	baseMaskLen, _ := baseNet.Mask.Size()
	var reqs []struct {
		hosts int
		index int
	}
	for i, count := range hostCounts {
		h, err := strconv.Atoi(count)
		if err != nil || h < 1 {
			return nil, fmt.Errorf("invalid host count '%s' at position %d", count, i+1)
		}
		reqs = append(reqs, struct {
			hosts int
			index int
		}{h, i})
	}

	// Sort by hosts descending
	sort.Slice(reqs, func(i, j int) bool {
		return reqs[i].hosts > reqs[j].hosts
	})

	currentIP := baseNet.IP.To4()
	allocations := make([]string, len(reqs))

	for _, req := range reqs {
		hosts := req.hosts
		hostBits := int(math.Ceil(math.Log2(float64(hosts + 2))))
		prefix := 32 - hostBits
		if prefix < baseMaskLen {
			return nil, fmt.Errorf("requirement of %d hosts too large for base network", hosts)
		}

		subnetMask := net.CIDRMask(prefix, 32)
		subnet := &net.IPNet{IP: currentIP, Mask: subnetMask}

		// Check if subnet fits
		bc := iputils.Broadcast(subnet)
		if !baseNet.Contains(subnet.IP) || !baseNet.Contains(bc) {
			return nil, fmt.Errorf("not enough space in base network for %d hosts", hosts)
		}

		// Format output
		firstUsable := iputils.IncrementIP(subnet.IP, 1)
		lastUsable := iputils.IncrementIP(bc, ^uint64(0))
		usable := fmt.Sprintf("%s - %s", firstUsable, lastUsable)
		allocStr := fmt.Sprintf("Subnet for %d hosts: %s (Usable: %d, Range: %s, Network: %s, Broadcast: %s)",
			hosts, subnet.String(), (1<<hostBits)-2, usable, subnet.IP, bc)
		allocations[req.index] = allocStr

		// Advance to next subnet
		subnetSize := uint64(1) << (32 - prefix)
		currentIP = iputils.IncrementIP(currentIP, subnetSize)
	}

	return allocations, nil
}
