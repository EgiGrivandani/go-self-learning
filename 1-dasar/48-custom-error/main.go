package main

import "fmt"

type ApiError struct {
	Code    int
	Status  string
	Message string
}

func (e *ApiError) Error() string {
	return fmt.Sprintf("[%d %s] %s", e.Code, e.Status, e.Message)
}
func main() {
	err := &ApiError{
		Code:    404,
		Status:  "Not Found",
		Message: fmt.Sprintf("data dengan ID %d tidak ditemukan", 12),
	}

	fmt.Println(err)
}
