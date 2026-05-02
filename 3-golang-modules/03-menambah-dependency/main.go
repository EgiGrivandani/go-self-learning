package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	id1 := uuid.New()
	fmt.Println("UUID 1:", id1.String())
}
