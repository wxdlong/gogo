package beibao

import (
	"fmt"
	"testing"
)

//
//       1  2  3  4  5  6  7  8  9  10 11 12 13 14 15
//5,12   0  0  0  0  12 12 12 12 12 12 12 12 12 12 12
//4,3    0  0  0  3  12 12 12 12 15 15 15 15 15 15 15
//7,10   0  0  0  3  12 12 12 12 12 15 15 22 22 22 22
//2,3    0  3  3  3  12 12 15 15 15 15 15 18 22 25 25
//6,6    0  3  3  3  12 12 15 15 15 15 18 18 22 25 25

//[刀, 重量: 2, 价值: 3]
//[西瓜, 重量: 7, 价值: 10]
//[书, 重量: 5, 价值: 12]
//max: 25
func TestHello(t *testing.T) {
	dx := []dongxi{{name: "书", weight: 5, value: 12}, {name: "笔", weight: 4, value: 3},
		{name: "西瓜", weight: 7, value: 10}, {name: "刀", weight: 2, value: 3},
		{name: "水", weight: 6, value: 6},
	}

	value := backpack(dx, 15)
	fmt.Println("max:", value)
}
