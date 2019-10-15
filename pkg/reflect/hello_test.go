package reflect

import (
	"fmt"
	"reflect"
	"testing"
)

func TestValueToString(t *testing.T) {
	int_value := reflect.ValueOf(3)
	fmt.Printf("%v",int_value)
}