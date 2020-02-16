package insert

import (
	"fmt"
	"testing"
)

func Test_sort(t *testing.T) {
	arr := []int{4, 7, 6, 5, 3, 2, 8, 1}
	fmt.Println("Origin", arr)
	sort(arr)
	fmt.Println("Sorted", arr)

}
