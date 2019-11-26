package arp

import (
	"encoding/hex"
	"fmt"
	"log"
	"testing"
)

func TestIPString(t *testing.T) {
	ip := []byte{0xc0, 0xa8, 0x1f, 0x01}
	log.Println(IPString(ip))
}

func TestMarshal(t *testing.T) {
	arpData, _ := hex.DecodeString("0001080006040001286c07dd6309c0a81f01000000000000c0a81f01")
	arp := Marshal(arpData)
	fmt.Println("Arp: ", arp)
}

func TestHello(t *testing.T) {
	hello()
}
