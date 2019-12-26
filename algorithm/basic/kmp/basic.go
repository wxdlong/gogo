package kmp

import "fmt"

type String string


//find 从源字符窜s中找出目标字符串t起始位置.
//暴力循环查找.
//没有找到,则返回-1
//第一个匹配位置返回0
func (s String)find(t string) int{
	if len(s)<len(t) {
		return -1
	}

	for k, v := range s {
        println("source",k,v)
        find := true
        if(k+len(t)) > len(s){
        	return -1
		}
		for ks, vs := range t {
			if t[ks] != s[ks+k] {
				find = false
				break
			}
			fmt.Printf("match [%d,%d] %c\n",ks, ks+k, vs)
		}
		if find {
			return k
		}
	}
	return -1
}