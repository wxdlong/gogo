package kmp

import (
	"fmt"
	"testing"
)

func TestNext(t *testing.T){
	fmt.Printf("%v",next("ababaaaba"))
	println(index("12abababaaabasef", "ababaaaba"))
}
