package main

func GenerateFibonacci() chan int {
	out := make(chan int, 10)
	ch := make(chan int, 10)
	ch <- 1
	ch <- 1
	go func() {
		for {
			a := <-ch
			b := <-ch
			ch <- b
			ch <- a + b
			out <- a
		}
	}()
	return out
}

func main() {
	ch := GenerateFibonacci()

	var a int
	b := <-ch
	for {
		a = b
		b = <-ch
		if a > b {
			panic("GenerateFibonacci lied")
		}
	}
}
