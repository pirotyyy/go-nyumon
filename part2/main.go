package main

const text = "Hello, World"

const (
	COMPLETE = iota
	ERROR
	WARNING
)

func main() {
	println(text)

	for i, v := range []int{1, 2, 3} {
		println(i, v)
	}
}
