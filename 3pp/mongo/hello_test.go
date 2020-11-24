package mongo

import (
	"fmt"
	"net"
	"testing"
)

func Test_hello(t *testing.T) {
	//hello()

    ip, ipNet, _ := net.ParseCIDR("3ffe::/112")
    fmt.Println(ip,ipNet)

	var ones, bits = ipNet.Mask.Size()
	fmt.Println(ones,bits)
}

