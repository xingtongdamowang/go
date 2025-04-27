package main

func defer1() {
	for i := 0; i < 10; i++ {
		defer func() {
			println(i)
		}()
	}
}

func defer2() {
	for i := 0; i < 10; i++ {
		defer func(val int) {
			print(val)
		}(i)
	}
}

func defer3() {
	for i := 0; i < 10; i++ {
		j := i
		defer func() {
			print(j)
		}()
	}
}
