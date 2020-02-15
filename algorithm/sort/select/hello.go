package _select

import "fmt"

//选择排序（Selection sort）是一种简单直观的排序算法。
//它的工作原理是：第一次从待排序的数据元素中选出最小（或最大）的一个元素，存放在序列的起始位置，
//然后再从剩余的未排序元素中寻找到最小（大）元素，然后放到已排序的序列的末尾。
//以此类推，直到全部待排序的数据元素的个数为零。选择排序是不稳定的排序方法。
func sort(src []int) {
	for i := 0; i < len(src); i++ {

		for j := i + 1; j < len(src); j++ {
			if src[i] > src[j] {
				fmt.Printf("swap [%d-%d] <--> [%d-%d]\n", i, src[i], j, src[j])
				src[i], src[j] = src[j], src[i]
			}
		}
	}
}
