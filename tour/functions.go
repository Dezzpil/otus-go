package main

import "fmt"

func add(x int, y int) int {
	return 10
}

func main() {
	fmt.Println(add(5, 5))

	sum := func (x int, y int) int { return x + y }
	fmt.Println(sum(6, 6))

	a, b := swap("foo", "bar")
	fmt.Println(a, b)

	fmt.Println(nake(100))
}

func swap(a, b string) (string, string) {
	return b, a
}

func nake(sum int) (x, y int) {
	x = sum + 1
	y = sum * 10
	return
}