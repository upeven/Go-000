package main

import (
	"fmt"

	"./service"
)

func main() {
	amount, err := service.GetUserPaidAmount(100)
	if err != nil {
		// 500 return
	}
	// 200 return
	fmt.Println(amount)
}
