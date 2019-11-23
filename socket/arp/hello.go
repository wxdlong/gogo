package arp

import (
	"golang.org/x/sys/unix"
	"log"
)

func hello() {

	fd, err := unix.Socket(unix.AF_PACKET, unix.SOCK_RAW, unix.ETH_P_ALL)
	if err != nil {
		log.Fatalln("create raw socket: ", err)
	}
	var addr unix.SockaddrLinklayer
	addr.Protocol = unix.ETH_P_ARP
	addr.Ifindex = 1
	addr.Hatype = unix.ARPHRD_ETHER

	//unix.Sendto(fd,)
}
