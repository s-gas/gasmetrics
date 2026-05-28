package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/s-gas/hauslab/services/gasmetrics/gasmetrics-bot/internal/reading"
	"github.com/s-gas/hauslab/services/gasmetrics/gasmetrics-bot/internal/telegram"
	"log"
	"net/http"
	"os"
)

func main() {
	bot, err := telegram.CreateBot()
	if err != nil {
		log.Fatal(err)
	}
	port := os.Getenv("PORT")
	url := fmt.Sprintf("http://gasmetrics-server:%s/readings", port)
	contentType := "application/json"

	updates := telegram.GetUpdates(bot)

	for update := range updates {
		r, err := reading.Parse(update.Message.Text)
		if err != nil {
			log.Println(err)
			continue
		}
		body, err := json.Marshal(r)
		if err != nil {
			log.Println(err)
			continue
		}
		resp, err := http.Post(url, contentType, bytes.NewReader(body))
		if err != nil {
			log.Println(err)
			continue
		}
		defer resp.Body.Close()

		var msg string
		if resp.StatusCode == http.StatusCreated {
			msg = "Entry added successfully"
		} else {
			msg = fmt.Sprintf("Failed to add entry: status code: %v", resp.StatusCode)
		}
		log.Println(msg)
		telegram.Reply(msg, bot, update)
	}
}
