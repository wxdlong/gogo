package arp

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"golang.org/x/sys/unix"
	"log"
)

//Frame 574: 62 bytes on wire (496 bits), 62 bytes captured (496 bits) on interface 0
//Interface id: 0 (any)
//Interface name: any
//Encapsulation type: Linux cooked-mode capture (25)
//Arrival Time: Nov 24, 2019 11:40:37.145140024 CST
//[Time shift for this packet: 0.000000000 seconds]
//Epoch Time: 1574566837.145140024 seconds
//[Time delta from previous captured frame: 0.001620000 seconds]
//[Time delta from previous displayed frame: 8.610262478 seconds]
//[Time since reference or first frame: 48.371350076 seconds]
//Frame Number: 574
//Frame Length: 62 bytes (496 bits)
//Capture Length: 62 bytes (496 bits)
//[Frame is marked: False]
//[Frame is ignored: False]
//[Protocols in frame: sll:ethertype:arp:vssmonitoring]
//[Coloring Rule Name: ARP]
//[Coloring Rule String: arp]
//Linux cooked capture
//Packet type: Broadcast (1)
//Link-layer address type: 1
//Link-layer address length: 6
//Source: XiaomiEl_dd:63:09 (28:6c:07:dd:63:09)
//Unused: 0000
//Protocol: ARP (0x0806)
//Padding: 00000000000000000000000000000000
//Address Resolution Protocol (request/gratuitous ARP)
//Hardware type: Ethernet (1)
//Protocol type: IPv4 (0x0800)
//Hardware size: 6
//Protocol size: 4
//Opcode: request (1)
//[Is gratuitous: True]
//Sender MAC address: XiaomiEl_dd:63:09 (28:6c:07:dd:63:09)
//Sender IP address: 192.168.31.1
//Target MAC address: 00:00:00_00:00:00 (00:00:00:00:00:00)
//Target IP address: 192.168.31.1
//VSS-Monitoring ethernet trailer, Source Port: 0
//Src Port: 0

//0001080006040001286c07dd6309c0a81f01000000000000c0a81f01

type Arp struct {
	HwType    uint16
	ProtoType uint16
	HwSize    uint8
	ProtoSize uint8
	OpCode    uint16
	SrcMac    []byte
	SrcIP     []byte
	DestMac   []byte
	DestIp    []byte
	Tail      uint16
}

func (arp Arp) String() string {
	return fmt.Sprintf("HwType:%d, ProtoType:%d, HwSize:%d, ProtoSize:%d \n OpCode:%d SrcMac:%v, SrcIP:%s, DestMac:%v, DestIP:%s ",
		arp.HwType, arp.ProtoType, arp.HwSize, arp.ProtoSize,
		arp.OpCode, arp.SrcMac, IPString(arp.SrcIP), arp.DestMac, IPString(arp.DestIp))
}

func IPString(ip []byte) string {
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}

func Marshal(data []byte) Arp {
	var arp Arp
	arp.HwType = binary.BigEndian.Uint16(data[:2])
	arp.ProtoType = binary.BigEndian.Uint16(data[2:4])
	arp.HwSize = data[4]
	arp.ProtoSize = data[5]
	arp.OpCode = binary.BigEndian.Uint16(data[6:8])
	arp.SrcMac = data[8:14]
	arp.SrcIP = data[14:18]
	arp.DestMac = data[18:24]
	arp.DestIp = data[24:28]

	return arp
}

func (arp *Arp) UnMarshal() []byte {
	packet := make([]byte, 62)

	return packet
}

func Hello() {
	arpData, _ := hex.DecodeString("000100010006286c07dd6309000008060001080006040001286c07dd6309c0a81f01000000000000c0a81f02000000000000000000000000000000000000")

	fd, err := unix.Socket(unix.AF_PACKET, unix.SOCK_RAW, unix.ETH_P_ALL)
	if err != nil {
		log.Fatalln("create raw socket: ", err)
	}
	var addr unix.SockaddrLinklayer
	addr.Protocol = unix.ETH_P_ARP
	addr.Ifindex = 1
	addr.Hatype = unix.ARPHRD_ETHER
	log.Println(fd)

	err = unix.Sendto(fd, arpData, 0, &addr)
	if err != nil {
		log.Fatalln("send raw socket: ", err)
	}
	//unix.Sendto(fd,)
}
