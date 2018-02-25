package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("ArgmentError\nUsage: gdig example.com.")
		os.Exit(0)
	}
	addr := os.Args[1]
	res, err := net.LookupHost(addr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}
