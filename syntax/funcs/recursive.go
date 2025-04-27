package main

func Recursive(n int) {
	if n > 10 {
		print("over")
		return
	}
	n++
	Recursive(n)
}
