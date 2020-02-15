package bubble

import (
	"fmt"
	"testing"
)

func Test_sort(t *testing.T) {
	src := []int{1, 5, 6, 20, 3, 29, 12, 17, 8}
	sort(src)
	for _, v := range src {
		fmt.Print(v, " ")
	}
}

func Test_sort2(t *testing.T) {
	src := []int{1, 3, 5, 6, 20, 7, 29, 12, 17, 8}
	sort2(src)
	for _, v := range src {
		fmt.Print(v, " ")
	}
}

func Test_sort3(t *testing.T) {
	src := []int{1, 3, 5, 6, 20, 7, 29, 12, 17, 8, 40, 60, 80}
	sort3(src)
	for _, v := range src {
		fmt.Print(v, " ")
	}
}
