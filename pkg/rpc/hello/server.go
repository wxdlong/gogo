package hello

import (
	"errors"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Args struct {
	A, B int
}

type Qutotient struct {
	Quo, Rem int
}

type Arith int

//函数必须是导出的(首字母大写)
//必须有两个导出类型的参数，
//第一个参数是接收的参数，第二个参数是返回给客户端的参数，第二个参数必须是指针类型的
//函数还要有一个返回值error
func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B

	log.Println("Multiply: ", args.A, args.B)
	return nil
}

func (t *Arith) Divide(args *Args, quo *Qutotient) error {
	if args.B == 0 {
		return errors.New("divide by Zero")
	}

	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B

	log.Println("Divide: ", args.A, args.B)

	return nil
}

func server() {
	arith := new(Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listener error ", e)
	}

	go http.Serve(l, nil)
	select {}
}
