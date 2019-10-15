package kmp

//next计算next数组值.
// ababaaba

func next(src string) []int {
	next := make([]int, len(src))

	pre, suf := 0, 1
	next[suf] = pre
	for suf < len(src)-1 {

		//if pre == -1, next[0]=-1
		//
		if pre == 0 || src[pre] == src[suf] {
			suf++
			pre++
			next[suf] = pre

		} else {
			pre = next[pre] //pre 回溯
		}
	}

	return next
}

func index(src, dest string) int {
	next := next(dest)
	i, j := 0, 0

	for i < len(src) {
		if src[i] == dest[j] {
			j++
			i++
			if len(dest) == j {
				return i - j
			}
			continue
		}else if j==0{
			  i++
		}
		j = next[j]
	}

	return -1
}
