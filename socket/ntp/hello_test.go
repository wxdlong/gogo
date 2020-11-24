package ntp

import (
	"golang.org/x/sys/unix"
	"log"
	"reflect"
	"testing"
)

type sock_opt struct {
	names, desc string
	level, name int
	typ         reflect.Kind
}

var opts []sock_opt

func init() {
	opts = []sock_opt{
		{names: "SO_RCVBUF", desc: "允许发送广播数据报", level: unix.SOL_SOCKET, name: unix.SO_RCVBUF, typ: reflect.Int},
		{names: "SO_BROADCAST", level: unix.SOL_SOCKET, name: unix.SO_BROADCAST, typ: reflect.Int},
		{names: "SO_DEBUG", level: unix.SOL_SOCKET, name: unix.SO_DEBUG, typ: reflect.Int},
		{names: "SO_DONTROUTE", level: unix.SOL_SOCKET, name: unix.SO_DONTROUTE, typ: reflect.Int},
		{names: "SO_KEEPALIVE", level: unix.SOL_SOCKET, name: unix.SO_KEEPALIVE, typ: reflect.Int},
		{names: "SO_OOBINLINE", level: unix.SOL_SOCKET, name: unix.SO_OOBINLINE, typ: reflect.Int},
		{names: "SO_SNDBUF", level: unix.SOL_SOCKET, name: unix.SO_SNDBUF, typ: reflect.Int},
		{names: "SO_RCVLOWAT", level: unix.SOL_SOCKET, name: unix.SO_RCVLOWAT, typ: reflect.Int},
		{names: "SO_SNDLOWAT", level: unix.SOL_SOCKET, name: unix.SO_SNDLOWAT, typ: reflect.Int},
		{names: "SO_REUSEADDR", desc: "允许重用本地地址", level: unix.SOL_SOCKET, name: unix.SO_REUSEADDR, typ: reflect.Int},
		{names: "SO_REUSEPORT", desc: "允许重用本地端口", level: unix.SOL_SOCKET, name: unix.SO_REUSEPORT, typ: reflect.Int},
		{names: "SO_TYPE", desc: "套接字类型", level: unix.SOL_SOCKET, name: unix.SO_TYPE, typ: reflect.Int},

		{names: "IP_HDRINCL", desc: "数据包含的IP首部", level: unix.SOL_IP, name: unix.IP_HDRINCL, typ: reflect.Int},
		{names: "IP_TTL", desc: "存活时间", level: unix.SOL_IP, name: unix.IP_TTL, typ: reflect.Int},
		{names: "IP_TOS", desc: "服务类型优先权", level: unix.SOL_IP, name: unix.IP_TOS, typ: reflect.Int},
		{names: "IP_MULTICAST_LOOP", desc: "是否换回", level: unix.SOL_IP, name: unix.IP_MULTICAST_LOOP, typ: reflect.Int},
		//{names: "IPV6_V6ONLY", desc:"禁止IPV4",level: unix.SOL_IPV6, name: unix.IPV6_V6ONLY, typ: reflect.Int},
	}
}

func TestHello(t *testing.T) {

	fd, err := unix.Socket(unix.AF_INET, unix.SOCK_DGRAM, 0)
	if err != nil {
		log.Fatalln("get socket: ", err)
	}

	for _, v := range opts {
		switch v.typ {
		case reflect.Int:
			s, err := unix.GetsockoptInt(fd, v.level, v.name)
			if err != nil {
				log.Fatalln("getSocket ", v.names, err)
			}
			log.Println(v.names, "==", v.desc, ":", s)
		}
	}

	s, err := unix.GetsockoptByte(fd, unix.SOL_IP, unix.IP_MULTICAST_LOOP)
	if err != nil {
		log.Fatalln("IP_MULTICAST_LOOP ", err)
	}
	log.Println("IP_MULTICAST_LOOP", s)
}
