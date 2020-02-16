package insert

import "fmt"

//插入排序的基本思想是：每步将一个待排序的记录，按其关键码值的大小插入前面已经排序的文件中适当位置上，直到全部插入完为止。

func sort(scr []int) {
	for i := 1; i < len(scr); i++ {
		fmt.Println("Sort index: ", i)
		for j := i; j > 0; j-- {
			if scr[j] > scr[j-1] {
				fmt.Printf("Swap [%d-%d] <--> [%d-%d]\n", i, scr[i], j, scr[j])
				scr[j-1], scr[j] = scr[j], scr[j-1]
			}
		}
	}
}
