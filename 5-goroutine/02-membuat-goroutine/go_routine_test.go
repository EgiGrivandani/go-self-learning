package membuat_goroutine

import (
	"fmt"
	"sync"
	"testing"
)

func sayHello(wg *sync.WaitGroup) {
	wg.Done()
	fmt.Println("hello world")
}

func TestGoroutine(t *testing.T) {
	var wg sync.WaitGroup

	wg.Add(1)

	go sayHello(&wg)
	fmt.Println("hello world2")

	wg.Wait()
}
