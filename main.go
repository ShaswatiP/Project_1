package main

import (
	"fmt"
	"server"
)

func main() {
	fmt.Println("Connected at port 9090")
	srv := server.Service_one()
	err := srv.Run(":9090")
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected at port 8000")
	srv := server.Service_two()
	err := srv.Run(":8000")
	if err != nil {
		panic(err)
	}
}
