package main

import (
	"fmt"
	"net"
	"os"
	"sync"
	"time"
)

func main() {
	// Check if the required command-line arguments are provided
	if len(os.Args) < 3 {
		// If not enough arguments provided, print the usage information and exit
		fmt.Fprintf(os.Stderr, "Usage: %s <Starting_IP_address> <Ending_IP_address>\n", os.Args[0])
		os.Exit(1)
	}

	// Extract starting and ending IP addresses from command-line arguments
	startIP := net.ParseIP(os.Args[1]).To4()
	endIP := net.ParseIP(os.Args[2]).To4()

	// Validate the format of the provided IP addresses
	if startIP == nil || endIP == nil {
		fmt.Println("Invalid IP address format.")
		os.Exit(1)
	}

	// Create a WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup
	// Map to store open ports discovered during scanning
	openPorts := make(map[string]bool)

	// Iterate through the IP range and launch a scan for each IP address
	for ip := startIP; !isEqualIP(ip, endIP); incrementIP(ip) {
		wg.Add(1)
		go func(ip net.IP) {
			defer wg.Done()
			scanIP(ip, openPorts)
		}(ip)
	}

	// Wait for all scans to finish
	wg.Wait()

	// Print the list of open ports discovered
	fmt.Println("Open Ports:")
	for port := range openPorts {
		fmt.Println(port)
	}
}

// Function to check if two IP addresses are equal
func isEqualIP(ip1, ip2 net.IP) bool {
	return ip1.Equal(ip2)
}

// Function to increment an IP address
func incrementIP(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

// Function to initiate port scanning for a given IP address
func scanIP(ip net.IP, openPorts map[string]bool) {
	fmt.Printf("Scanning started: %s\n", ip.String())

	// Create a WaitGroup for port scanning
	var wg sync.WaitGroup

	// Iterate through all possible port numbers and launch a goroutine for each port
	for port := 1; port <= 65535; port++ {
		wg.Add(1)
		go func(port int) {
			defer wg.Done()
			scanPort(ip, port, openPorts)
		}(port)
	}

	// Wait for all port scans to finish
	wg.Wait()
}

// Function to scan a specific port of a given IP address
func scanPort(ip net.IP, port int, openPorts map[string]bool) {
	address := fmt.Sprintf("%s:%d", ip.String(), port)
	conn, err := net.DialTimeout("tcp", address, 2*time.Second)
	if err != nil {
		return
	}
	defer conn.Close()

	// If port is open, record it in the openPorts map
	openPorts[address] = true

	// Print the open port information
	fmt.Printf("Port %d open - %s\n", port, address)

	// Perform vulnerability scanning for the open port
	checkVulnerabilities(ip, port)
}

// Function to perform vulnerability scanning for a specific port
func checkVulnerabilities(ip net.IP, port int) {
	switch port {
	case 21:
		// FTP vulnerability check
		checkFTP(ip, port)
	case 22:
		// SSH vulnerability check
		checkSSH(ip, port)
	case 80:
		// HTTP vulnerability check
		checkHTTP(ip, port)
	case 443:
		// HTTPS vulnerability check
		checkHTTPS(ip, port)
	default:
		// General message for other ports
		fmt.Printf("General vulnerability check - %s:%d\n", ip.String(), port)
	}
}

// Function to check for FTP vulnerabilities
func checkFTP(ip net.IP, port int) {
	// FTP vulnerability check code can be added here
	// For example, checking for common FTP vulnerabilities
	// For instance, checking for anonymous access
	fmt.Printf("FTP vulnerability check - %s:%d\n", ip.String(), port)
}

// Function to check for SSH vulnerabilities
func checkSSH(ip net.IP, port int) {
	// SSH vulnerability check code can be added here
	// For example, checking for weak password usage
	fmt.Printf("SSH vulnerability check - %s:%d\n", ip.String(), port)
}

// Function to check for HTTP vulnerabilities
func checkHTTP(ip net.IP, port int) {
	// HTTP vulnerability check code can be added here
	// For example, examining HTTP header information
	fmt.Printf("HTTP vulnerability check - %s:%d\n", ip.String(), port)
}

// Function to check for HTTPS vulnerabilities
func checkHTTPS(ip net.IP, port int) {
	// HTTPS vulnerability check code can be added here
	// For example, checking for weak encryption usage
	fmt.Printf("HTTPS vulnerability check - %s:%d\n", ip.String(), port)
}
