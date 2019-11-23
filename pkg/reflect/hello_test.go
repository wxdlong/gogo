package reflect

import (
	"fmt"
	"reflect"
	"testing"
)

func TestValueToString(t *testing.T) {
	var i3 int
	int_value := reflect.ValueOf(i3)
	fmt.Println("canAddr", int_value.CanAddr())
	fmt.Println("canSet ", int_value.CanSet())
	fmt.Println("value ", int_value.Int())
	int_type := reflect.TypeOf(3)
	fmt.Printf("%v\n", int_value)
	fmt.Printf("%v\n", int_type)
}
