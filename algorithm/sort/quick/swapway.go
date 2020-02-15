package quick

import "fmt"

//qucikSort 递归
//1. 在数据集之中，选择一个元素作为"基准"（pivot）。
//2. 所有小于"基准"的元素，都移到"基准"的左边；所有大于"基准"的元素，都移到"基准"的右边。
//3. 对"基准"左边和右边的两个子集，不断重复第一步和第二步，直到所有子集只剩下一个元素为止。
func qucikSort2(src []int, start, end int) {
	if start >= end {
		fmt.Println("QuickSort end.", start, end)
		return
	}

	pivotIndex := swap(src, start, end)

	qucikSort2(src, start, pivotIndex-1)
	qucikSort2(src, pivotIndex+1, end)
}

//交换法
//返回基准值位置
func swap(src []int, start, end int) int {
	left, right := start, end
	pivot := src[start] //第一个值为基准值
	fmt.Println("pivot: ", pivot, start, end)
	for right != left {
		//从右扫描找出第一个比基准值小的位置
		for left < right && src[right] > pivot {
			fmt.Printf("Right: %d-%d \n", right, src[right])
			right--
		}

		//从左扫描找出第一个比基准值大的位置
		for left < right && src[left] <= pivot {
			fmt.Printf("Left: %d \n", src[left])
			left++
		}

		if left < right {
			//如果一轮扫描完毕， 交换左右两个值
			fmt.Printf("Scan Over: left>[%d:%d] - right>[%d:%d] \n", left, src[left], right, src[right])
			src[right], src[left] = src[left], src[right]
		}
	}

	fmt.Printf("Swap Pivot: pivot>[%d:%d] - right>[%d:%d] \n", 0, src[0], right, src[right])
	src[start], src[right] = src[right], src[start]

	return left
}
