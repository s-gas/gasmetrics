package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"net/http"
	"github.com/s-gas/hauslab/services/gasmetrics/gasmetrics-bot/internal/telegram"
)

func main() {
	bot, err := telegram.CreateBot()
	if err != nil {
		log.Fatal(err)
	}
	port := os.Getenv("PORT")
	url := fmt.Sprintf("http://gasmetrics:%s/readings", port)
	contentType := "application/json"

	updates := telegram.GetUpdates(bot)

	for update := range updates {
		value := update.Message.Text
		body := strings.NewReader(fmt.Sprintf(`{"value": %s}`, value))
		resp, err := http.Post(url, contentType, body)
		if err != nil {
			log.Println(err)
		}
		defer resp.Body.Close()
	
		var msg string
		if resp.StatusCode == http.StatusOK {
			msg = "Entry added successfully"
		} else {
			msg = "Failed to add entry"
		}
		log.Println(msg)
		telegram.Reply(msg, bot, update)
	}
}
