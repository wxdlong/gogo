package beibao

import "fmt"

//n个物品和1个背包，每个物品i对应的重量为wi，价值为vi，背包的容量为W。
//每个物品只有一件，要么装入，要么不装入，不可拆分。
//如何选取物品装入背包，使背包所装入的物品的总价值最大？要求算法输出最优值和最优解。

type dongxi struct {
	name   string
	weight int
	value  int
}

func (dx *dongxi) String() string {
	return fmt.Sprintf("[%s, 重量: %d, 价值: %d]", dx.name, dx.weight, dx.value)
}

//参考
//https://www.jianshu.com/p/c289cd8ae0ed
func backpack(dx []dongxi, total int) int {
	table := make([][]int, len(dx))

	for i := 0; i < len(dx); i++ {
		table[i] = make([]int, total+1)
	}

	for i := dx[0].weight; i < total; i++ {
		table[0][i] = dx[0].value
	}

	for i := 1; i < len(dx); i++ {
		for weight := dx[i].weight; weight <= total; weight++ {
			pre := table[i-1][weight]                             //在当前重量下，前几个物品最大价值
			cMax := dx[i].value + table[i-1][weight-dx[i].weight] //当前物品价值+ （当前重量-当前物品重量时最大价值）
			table[i][weight] = max(pre, cMax)
		}
	}

	index := len(dx) - 1
	for weight := total; index >= 0; index-- {
		pre := index - 1
		if index == 0 && weight > dx[index].weight || table[index][weight] > table[pre][weight] {
			fmt.Println(&dx[index])
			weight = weight - dx[index].weight //减去当前重量时，余下重量在表格内的最大价值列数
		}
	}
	return table[len(dx)-1][total]
}

//max 取两个数的最大值
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
