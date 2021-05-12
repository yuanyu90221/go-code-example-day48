package main

import (
	"github.com/yuanyu90221/go-code-problem-day48/sol"
)

func main() {
	a := []int{1, 2, 3, 4}
	l := sol.ConvertList(&a)
	sol.Transerval(l)
	l = sol.RightShiftLinkedListByKSteps(l, 5)
	sol.Transerval(l)
}
