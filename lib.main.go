package main

import "fmt"

func addEx(strs []string, funcs []func(x int, y int)) (resStrs []string, resFuncs []func(x int, y int)) {
	resStrs, resFuncs = clickBase(strs, ".png", funcs)
	return
}

func addExFunc(strs []string, funcs []func(x int, y int)) (resStrs []string, resFuncs []func(x int, y int)) {
	strs = append(strs, getSystemImg(".png"))
	funcs = append(funcs, func(x, y int) {
		a := ""
		fmt.Println(a)
	})
	resStrs = strs
	resFuncs = funcs
	return
}
