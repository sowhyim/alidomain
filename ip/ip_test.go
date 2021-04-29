package ip

import "testing"

func TestGetIP(t *testing.T) {
	ip := GetIP()
	t.Log(ip)
}