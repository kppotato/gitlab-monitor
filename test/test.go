package main

import (
	"sort"
	"fmt"
)
//1b430d78acb4775d621c66eab8bdb93a4b73abc1695c60347c8369150c7a316a
//08a54ec5da45ab5c9efba259d3fa42b31a46c64a7e9581317e5220ceaa2bea8e

func main()  {
	var f=[]int{1,2,3,4,8,5}
	sort.Slice(f, func(i, j int) bool {
		if f[i]<f[j]{
			return false
		}
		return true
	})
	fmt.Println(f)

}
