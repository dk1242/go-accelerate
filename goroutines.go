package main

import (
	"fmt"
	"sync"
)

func Add(a, b int) int {
	return a + b
}

func Subtract(a, b int) int {
	return a - b
}

func Multiply(a, b int) int {
	return a * b
}

func Divide(a, b int) int {
	return a / b
}

type Task struct {
	operation string
	a, b      int
}

func calculate(i int, ch chan Task, wg *sync.WaitGroup) {
	for task := range ch {
		switch task.operation {
		case "add":
			fmt.Println("Result of addition is", Add(task.a, task.b), "done by goroutine", i)
		case "subtract":
			fmt.Println("Result of subtraction is", Subtract(task.a, task.b), "done by goroutine", i)
		case "multiply":
			fmt.Println("Result of multiplication is", Multiply(task.a, task.b), "done by goroutine", i)
		case "divide":
			fmt.Println("Result of division is", Divide(task.a, task.b), "done by goroutine", i)
		}
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan Task)
	for i := 0; i < 10; i++ {
		go calculate(i, ch, &wg)
		wg.Add(1)
	}

	ch <- Task{"add", 10, 5}
	ch <- Task{"subtract", 10, 5}
	ch <- Task{"multiply", 10, 5}
	ch <- Task{"divide", 10, 5}

	close(ch)
	wg.Wait()

}
