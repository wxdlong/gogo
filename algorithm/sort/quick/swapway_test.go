package quick

import (
	"fmt"
	"testing"
)

func Test_swap(t *testing.T) {
	arr := []int{4, 7, 6, 5, 3, 2, 8, 1}
	qucikSort2(arr, 0, len(arr)-1)

	for _, v := range arr {
		fmt.Print(v, " ")
	}
}
