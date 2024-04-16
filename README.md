# Port Scanner

This is a simple port scanner tool written in Go. It allows you to scan a range of IP addresses for open ports and performs basic vulnerability checks on the open ports.

## Usage

To use this tool, you need to have Go installed on your system.

### Running the Executable

If you have already built the executable, you can run it with the following command:



For example:


This command will scan the IP addresses from `192.168.1.1` to `192.168.1.10` for open ports.

### Using `go run`

If you want to run the tool without building the executable, you can use the `go run` command:

go run main.go <Starting_IP_address> <Ending_IP_address> 

### Using `.exe `
./main.exe <Starting_IP_address> <Ending_IP_address> 
