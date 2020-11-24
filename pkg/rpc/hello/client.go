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

	//调用命名函数,等待返回.
	err := client.Call("Arith.Multiply", args, &reply)

	if err != nil {
		log.Fatal("arith error ", err)
	}

	fmt.Printf("Arith: %d*%d=%d \n", args.A, args.B, reply)

	quto := new(Qutotient)

	//异步调用. 立马返回调用结构体.
	call := client.Go("Arith.Divide", args, &quto, nil)
	<-call.Done //等待channel返回结果

	log.Println(quto)

}
