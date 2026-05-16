package main

import (
	"fmt"
	"os"
	"log"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("usage: gasmetrics-cli <value>")
	}
	fmt.Println("hello")
}
