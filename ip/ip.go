package ip

import (
	"net"
	"strings"
)

// GetPublicIP .
func GetPublicIP() {

}

// GetIP .
func GetIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().String()
	idx := strings.LastIndex(localAddr, ":")
	return localAddr[0:idx]

}

// IsPublicIP .
func IsPublicIP(ip net.IP) bool {

	return false
}
