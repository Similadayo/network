package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/similadayo/subnet"
	vlsm "github.com/similadayo/vslm"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter command (subnet/vlsm/exit): ")
		command, _ := reader.ReadString('\n')
		command = strings.TrimSpace(command)
		if command == "exit" {
			fmt.Println("Goodbye!")
			break
		}

		switch command {
		case "subnet":
			fmt.Print("Enter CIDR (e.g., 192.168.0.0/24): ")
			cidr, _ := reader.ReadString('\n')
			cidr = strings.TrimSpace(cidr)
			result, err := subnet.Calcualte(cidr)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}
			fmt.Println("\nSubnet Calculation:")
			fmt.Println(result)
			printExplanations()
		case "vlsm":
			fmt.Print("Enter base CIDR (e.g., 192.168.0.0/24): ")
			baseCIDR, _ := reader.ReadString('\n')
			baseCIDR = strings.TrimSpace(baseCIDR)
			fmt.Print("Enter host counts (space-separated, e.g., 100 50 20): ")
			hostsInput, _ := reader.ReadString('\n')
			hostCounts := strings.Fields(strings.TrimSpace(hostsInput))
			if len(hostCounts) == 0 {
				fmt.Println("Error: At least one host count required")
				continue
			}
			allocations, err := vlsm.Calculate(baseCIDR, hostCounts)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}
			fmt.Println("\nVLSM Allocation:")
			for _, alloc := range allocations {
				fmt.Println(alloc)
			}
			printExplanations()
		default:
			fmt.Println("Invalid command. Use 'subnet', 'vlsm', or 'exit'.")
		}
	}
}

func printExplanations() {
	fmt.Println("\nExplanations:")
	fmt.Println("- Network: The first IP address in the subnet, used to identify the network.")
	fmt.Println("- Broadcast: The last IP address in the subnet, used for broadcasting messages.")
	fmt.Println("- Usable Range: IP addresses assignable to hosts (from first usable to last usable).")
	fmt.Println("- Usable Hosts: Number of assignable IP addresses (total minus network and broadcast).")
	fmt.Println("- Host Bits: Number of bits in the host portion of the subnet (32 - prefix length).")
	fmt.Println("- Mask: Subnet mask derived from the CIDR prefix (e.g., /24 -> 255.255.255.0).")
	fmt.Println("- In VLSM, subnets are allocated starting with the largest host requirement to optimize address space.")
}
