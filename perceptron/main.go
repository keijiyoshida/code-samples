package main

import "fmt"

func main() {
	p := newPerceptron(0.1, 10)
	p.fit(x, y)
	fmt.Println(p.errors)
	fmt.Println(p.w)
}
