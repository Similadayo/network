# Network Utility Tool

A comprehensive Go-based command-line tool for network subnet calculations and VLSM (Variable Length Subnet Masking) allocations.

## Features

- **Subnet Calculator**: Calculate detailed subnet information for any given CIDR notation
- **VLSM Allocator**: Perform Variable Length Subnet Masking for efficient IP address allocation
- **Interactive CLI**: User-friendly command-line interface with step-by-step guidance
- **IPv4 Support**: Full support for IPv4 network calculations

## Installation

1. Ensure you have Go 1.23.2 or later installed
2. Clone or download this repository
3. Navigate to the project directory
4. Build the project:

   ```bash
   go build -o network-tool
   ```

## Usage

Run the compiled binary:

```bash
./network-tool
```

### Available Commands

- `subnet`: Calculate subnet details for a given CIDR notation
- `vlsm`: Perform VLSM allocation for multiple host requirements
- `exit`: Exit the application

### Subnet Calculation Example

```
Enter command (subnet/vlsm/exit): subnet
Enter CIDR (e.g., 192.168.0.0/24): 192.168.1.0/24

Subnet Calculation:
Subnet: 192.168.1.0 (Host Bits: 8, Usable Hosts: 254, Range: 192.168.1.1 - 192.168.1.254, Network: 192.168.1.0, Broadcast: 192.168.1.255, Mask: 255.255.255.0)
```

### VLSM Allocation Example

```
Enter command (subnet/vlsm/exit): vlsm
Enter base CIDR (e.g., 192.168.0.0/24): 192.168.0.0/22
Enter host counts (space-separated, e.g., 100 50 20): 100 50 20 10

VLSM Allocation:
Subnet for 100 hosts: 192.168.0.0/25 (Usable: 126, Range: 192.168.0.1 - 192.168.0.126, Network: 192.168.0.0, Broadcast: 192.168.0.127)
Subnet for 50 hosts: 192.168.0.128/26 (Usable: 62, Range: 192.168.0.129 - 192.168.0.190, Network: 192.168.0.128, Broadcast: 192.168.0.191)
Subnet for 20 hosts: 192.168.0.192/27 (Usable: 30, Range: 192.168.0.193 - 192.168.0.222, Network: 192.168.0.192, Broadcast: 192.168.0.223)
Subnet for 10 hosts: 192.168.0.224/28 (Usable: 14, Range: 192.168.0.225 - 192.168.0.238, Network: 192.168.0.224, Broadcast: 192.168.0.239)
```

## Project Structure

```
network/
├── main.go              # Main application entry point
├── go.mod              # Go module definition
├── iputils/
│   └── iputils.go      # IP utility functions
├── subnet/
│   └── subnet.go       # Subnet calculation logic
└── vslm/
    └── vslm.go         # VLSM allocation logic
```

## Key Components

### IP Utilities (`iputils/iputils.go`)

- `Broadcast()`: Calculates broadcast address for a given network
- `IncrementIP()`: Increments IP address by specified value
- `MaskToString()`: Converts CIDR prefix to subnet mask string

### Subnet Calculator (`subnet/subnet.go`)

- `Calculate()`: Computes detailed subnet information including:
  - Network address
  - Broadcast address
  - Usable IP range
  - Number of usable hosts
  - Subnet mask
  - Host bits calculation

### VLSM Allocator (`vslm/vslm.go`)

- `Calculate()`: Performs VLSM allocation with:
  - Automatic sorting by host requirements (largest first)
  - Efficient IP space utilization
  - Validation for sufficient address space
  - Detailed allocation output

## Terminology Explained

- **Network**: The first IP address in the subnet, used to identify the network
- **Broadcast**: The last IP address in the subnet, used for broadcasting messages
- **Usable Range**: IP addresses assignable to hosts (from first usable to last usable)
- **Usable Hosts**: Number of assignable IP addresses (total minus network and broadcast)
- **Host Bits**: Number of bits in the host portion of the subnet (32 - prefix length)
- **Mask**: Subnet mask derived from the CIDR prefix (e.g., /24 → 255.255.255.0)
- **VLSM**: Variable Length Subnet Masking - allocates subnets starting with largest host requirement to optimize address space

## Development

### Building from Source

```bash
go build -o network-tool
```

### Running Tests

```bash
go test ./...
```

### Module Dependencies

This project uses Go modules. Dependencies are managed through the `go.mod` file.

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## License

This project is open source and available under the MIT License.

## Support

For issues or questions, please open an issue in the project repository or contact the maintainers.
