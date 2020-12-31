package main

import (
	"fmt"
	"net"
	"strings"
)

func getMacAddr() ([]string, error) {
	ifas, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	var as []string
	for _, ifa := range ifas {
		a := ifa.HardwareAddr.String()
		if a != "" {
			as = append(as, a)
		}
	}
	return as, nil
}

func main() {
	list, err := net.Interfaces()
	if err != nil {
		panic(err)
	}

	for i, iface := range list {
		if string(iface.Name) != "lo" && string(iface.Name) != "Loopback" {
			fmt.Printf("%d name=%s %v\n", i, iface.Name, iface.HardwareAddr)
			addrs, err := iface.Addrs()
			if err != nil {
				panic(err)
			}
			for j, addr := range addrs {
				fmt.Printf(" %d %v\n", j, addr)
			}
		}
	}

	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().String()
	idx := strings.LastIndex(localAddr, ":")
	fmt.Printf("IP WAN: %v\n", localAddr[0:idx])

}
