package main

import (
	"fmt"
)

// Sqrt 下面的 z² − x 是 z² 到它所要到达的值（即 x）的距离，除以的 2z 为 z² 的导数，我们通过 z² 的变化速度来改变 z 的调整量。这种通用方法叫做牛顿法。它对很多函数，特别是平方根而言非常有效。
func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
	}
	return z
}

func mainFun() {
	fmt.Println(Sqrt(2))
}
