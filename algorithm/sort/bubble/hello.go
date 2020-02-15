package bubble

import "fmt"

//sort 冒泡排序，从底至上逐个对比，比较相邻两个数大小，如果比后面数大，则交换
//一轮下来，最大的就浮到最顶上。
//继续循环其余数列，则所有数从小到大排列。
//时间很杂度 len(src):  例 8+7+6+5+4+3+2+1=36
func sort(src []int) {
	len := len(src)
	compare := 0
	for i := 0; i < len-1; i++ {
		for j := 0; j < len-i-1; j++ {
			compare++
			if src[j] > src[j+1] {
				fmt.Println("swap ", j, src[j], src[j+1])
				src[j], src[j+1] = src[j+1], src[j]
			}
		}
	}
	fmt.Println("Compare: ", compare)
}

//sort2 冒泡排序，增加一个标记位，如果一轮比较不发生交换则说明数列己经排好。
//这种情况是指数列前几位己经排好序
func sort2(src []int) {
	len := len(src)
	compare := 0
	for i := 0; i < len-1; i++ {
		done := true
		for j := 0; j < len-i-1; j++ {
			compare++
			if src[j] > src[j+1] {
				fmt.Println("swap ", j, src[j], src[j+1])
				src[j], src[j+1] = src[j+1], src[j]
				done = false
			}
		}
		if done {
			fmt.Println("Sort complete: ", i)
			break
		}
	}
	fmt.Println("Compare: ", compare)
}

//sort3 冒泡排序，增加一个标记记。
//1。 如果发生交换，记录最后交换的位置
//这种情况是指数列前几位己经排好序
func sort3(src []int) {
	len := len(src)
	compare := 0
	flag := len - 1

	for flag > 0 {
		k := flag
		flag = 0
		for j := 0; j < k; j++ {
			compare++
			if src[j] > src[j+1] {
				fmt.Println("swap ", j, src[j], src[j+1])
				src[j], src[j+1] = src[j+1], src[j]
				flag = j
			}
		}
	}
	fmt.Println("Compare: ", compare)
}
