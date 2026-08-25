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

func DisplayNumber(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Display Number : ", i)
}

func TestGoroutine(t *testing.T) {
	var wg sync.WaitGroup

	wg.Add(1)

	go sayHello(&wg)
	fmt.Println("new hello")

	wg.Wait()
	

}

func TestGoroutine2(t *testing.T) {
	var wg sync.WaitGroup

	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go DisplayNumber(i, &wg)
	}

	wg.Wait()



}
