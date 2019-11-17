package hello

import (
	"fmt"
	"log"
	"net/rpc"
)

func client() {
	client, e := rpc.DialHTTP("tcp", ":1234")
	if e != nil {
		log.Fatal("connect to server ", e)
	}

	args := &Args{7, 8}
	var reply int
	err := client.Call("Arith.Multiply", args, &reply)

	if err != nil {
		log.Fatal("arith error ", err)
	}

	fmt.Printf("Arith: %d*%d=%d", args.A, args.B, reply)
}
