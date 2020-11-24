package kmp

import (
	"testing"
)

func TestFind(t *testing.T){
	 var src String
	 src="abcdef"
	 println(src.find("bcd"))
	println(src.find("dbc"))
}
