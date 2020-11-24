package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"os"
	"proto/entity"
)

func write() {
	p1 := &entity.Person{
		Id:   0,
		Name: "bingo",
		Phones: []*entity.Phone{&entity.Phone{
			Type:   0,
			Number: "17070858080",
		}},
	}
	data, _ := proto.Marshal(p1)
	ioutil.WriteFile("test.txt", data, os.ModePerm)
}

func read() {
	data, _ := ioutil.ReadFile("test.txt")
	p1 := &entity.Person{}

	proto.Unmarshal(data, p1)
	fmt.Printf("%+v", p1)
}

func main() {
	//write()
	read()
}
