package main

import (
	"fmt"

	"github.com/LOLONEKIT2019/prof_of_work/internal/server"
)

func main() {
	err := server.Start()
	if err != nil {
		fmt.Println("start server", err)
	}
}
