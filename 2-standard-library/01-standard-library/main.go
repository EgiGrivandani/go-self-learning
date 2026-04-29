package main

import (
	"fmt"
	"time"
)

func main() {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	t := time.Now().In(loc)
	fmt.Println("The time is:", t.Format(time.RFC822))
}
