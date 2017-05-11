package main

import "fmt"

func fibonacci() func() int {
	f0 := 0
	f1 := 1
	fn := f0 + f1
	count := 0
	return func() int {
		var ret int
		switch {
		case count == 0:
			count++
			ret = 0
		case count == 1:
			count++
			ret = 1
		default:
			ret = fn
			f0 = f1
			f1 = fn
			fn = f0 + f1
		}

		return ret
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
