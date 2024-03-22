package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "Kullanım: %s <Hedef_IP_adresi_başlangıç> <Hedef_IP_adresi_bitiş>\n", os.Args[0])
		os.Exit(1)
	}

	hedefBaslangicIP := os.Args[1]
	hedefBitisIP := os.Args[2]

	startIP := net.ParseIP(hedefBaslangicIP).To4()
	endIP := net.ParseIP(hedefBitisIP).To4()

	if startIP == nil || endIP == nil {
		fmt.Println("Geçersiz IP adresi formatı.")
		os.Exit(1)
	}

	var wg sync.WaitGroup

	for ip := startIP; !isEqualIP(ip, endIP); incrementIP(ip) {
		wg.Add(1)
		go func(ip net.IP) {
			defer wg.Done()
			scanIP(ip)
		}(ip)
	}

	wg.Wait()
}

func isEqualIP(ip1, ip2 net.IP) bool {
	return ip1.Equal(ip2)
}

func incrementIP(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

func scanIP(ip net.IP) {
	fmt.Printf("Tarama başlatıldı: %s\n", ip.String())

	var wg sync.WaitGroup

	for port := 1; port <= 65535; port++ {
		wg.Add(1)
		go func(port int) {
			defer wg.Done()
			scanPort(ip, port)
		}(port)
	}

	wg.Wait()
}

func scanPort(ip net.IP, port int) {
	address := fmt.Sprintf("%s:%d", ip.String(), port)
	conn, err := net.DialTimeout("tcp", address, 2*time.Second)
	if err != nil {
		return
	}
	defer conn.Close()
	fmt.Printf("Port %d açık - %s\n", port, address)
}

func scanService(ip net.IP, port int) {
	address := fmt.Sprintf("%s:%d", ip.String(), port)
	conn, err := net.DialTimeout("tcp", address, 2*time.Second)
	if err != nil {
		return
	}
	defer conn.Close()

	// Servis tespiti
	// Örnek olarak, HTTP başlık bilgisi alınabilir.
}
