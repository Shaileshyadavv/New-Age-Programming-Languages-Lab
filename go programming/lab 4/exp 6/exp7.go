package main
import (
	"fmt"
	"sync"
	"time"
)

type Result struct {
	Task  string
	Value []int
}

func squares(n int, ch chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("[squares]   started")
	time.Sleep(300 * time.Millisecond) 
	out := make([]int, 0, n)
	for i := 1; i <= n; i++ {
		out = append(out, i*i)
	}
	fmt.Println("[squares]   completed")
	ch <- Result{"Squares", out}
}

func cubes(n int, ch chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("[cubes]     started")
	time.Sleep(100 * time.Millisecond)
	out := make([]int, 0, n)
	for i := 1; i <= n; i++ {
		out = append(out, i*i*i)
	}
	fmt.Println("[cubes]     completed")
	ch <- Result{"Cubes", out}
}

func fibonacci(n int, ch chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("[fibonacci] started")
	time.Sleep(200 * time.Millisecond)
	out := make([]int, 0, n)
	a, b := 0, 1
	for i := 0; i < n; i++ {
		out = append(out, a)
		a, b = b, a+b
	}
	fmt.Println("[fibonacci] completed")
	ch <- Result{"Fibonacci", out}
}

func main() {
	const n = 10
	results := make(chan Result, 3) 
	var wg sync.WaitGroup

	fmt.Println("main: launching goroutines")
	wg.Add(3)
	go squares(n, results, &wg)
	go cubes(n, results, &wg)
	go fibonacci(n, results, &wg)

	go func() {
		wg.Wait()
		close(results)
	}()

	for r := range results {
		fmt.Printf("main: received %-9s -> %v\n", r.Task, r.Value)
	}
	fmt.Println("main: all tasks done")
}