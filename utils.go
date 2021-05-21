package main

import (
	"fmt"
	"net"
)

func printDNS(host string, ips []net.IP) {
	for _, ip := range ips {
		fmt.Println("DNS ResolveIP:", host, ":", ip.String())
	}

	fmt.Println()
}
