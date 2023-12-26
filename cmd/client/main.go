package main

import (
	"fmt"

	"github.com/LOLONEKIT2019/prof_of_work/internal/client"
)

func main() {
	err := client.Start()
	if err != nil {
		fmt.Println("initialize client", err)
	}
}
