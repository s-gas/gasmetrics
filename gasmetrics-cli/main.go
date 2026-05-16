package main

import (
	"os"
	"fmt"
	"log"
	"strings"
	"net/http"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("usage: gasmetrics-cli <value>")
	}
	value := os.Args[1]
	url := "http://hauslab/readings"
	contentType := "application/json"
	body := strings.NewReader(fmt.Sprintf(`{"value": %s}`, value))
	resp, err := http.Post(url, contentType, body)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	
	if resp.StatusCode == http.StatusOK {
		log.Println("Entry added successfully")
	} else {
		log.Println("Failed to add entry")
	}
}
