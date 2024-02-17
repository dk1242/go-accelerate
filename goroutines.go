package main

import (
	"fmt"
	"sync"

	"github.com/Priyanka488/calculator/api"
)

type Task struct {
	operation string
	a, b      int
}

func calculate(i int, ch chan Task, wg *sync.WaitGroup) {
	for task := range ch {
		switch task.operation {
		case "add":
			fmt.Println("Result of addition is", api.Add(task.a, task.b), "done by goroutine", i)
		case "subtract":
			fmt.Println("Result of subtraction is", api.Subtract(task.a, task.b), "done by goroutine", i)
		case "multiply":
			fmt.Println("Result of multiplication is", api.Multiply(task.a, task.b), "done by goroutine", i)
		case "divide":
			fmt.Println("Result of division is", api.Divide(task.a, task.b), "done by goroutine", i)
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
