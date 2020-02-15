package quick

import "fmt"

//qucikSort 递归
//1. 在数据集之中，选择一个元素作为"基准"（pivot）。
//2. 所有小于"基准"的元素，都移到"基准"的左边；所有大于"基准"的元素，都移到"基准"的右边。
//3. 对"基准"左边和右边的两个子集，不断重复第一步和第二步，直到所有子集只剩下一个元素为止。
func qucikSort(src []int, start, end int) {
	if start >= end {
		fmt.Println("QuickSort end.", start, end)
		return
	}

	pivotIndex := partition(src, start, end)

	qucikSort(src, start, pivotIndex-1)
	qucikSort(src, pivotIndex+1, end)
}

//填坑法
func partition(src []int, start, end int) int {
	left := start
	right := end
	keng := start + (end-start)/2 //取中间位置为第一个坑位
	pivot := src[keng]            //第一个坑位的为基准值
	fmt.Println("partition", start, end, " Keng[", keng, ":", pivot)

	for right > left {
		//重右至左扫描，如果值大于基准值，则将值填到对应坑的位置。
		for ; right >= keng; right-- {
			if src[right] < pivot {
				fmt.Printf("Right, put [%d,%d] to [%d,%d]\n", right, src[right], keng, src[keng])
				src[keng] = src[right]
				keng = right //调换位置的地方变成坑位。
				right--
				break
			}
		}

		for ; left <= keng; left++ {
			if src[left] > pivot {
				fmt.Printf("Left, put [%d,%d] to [%d,%d]\n", left, src[left], keng, src[keng])

				src[keng] = src[left] //调换位置的地方变成坑位。
				keng = left
				left++
				break
			}
		}
	}

	src[keng] = pivot //将基准值放入最后的坑位中。
	return keng
}
