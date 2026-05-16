package main

import (
	"context"
	"fmt"
	"log"

	"github.com/s-gas/hauslab/services/gasmetrics/gasmetrics/internal/handler"
	"github.com/s-gas/hauslab/services/gasmetrics/gasmetrics/internal/postgres"
	"github.com/s-gas/hauslab/services/gasmetrics/gasmetrics/internal/telegram"
)

func main() {
	ctx := context.Background()
	conn, err := postgres.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(context.Background())

	// add goroutine for starting HTTP server

	bot, err := telegram.CreateBot()
	if err != nil {
		log.Fatal(err)
	}
	updates := telegram.GetUpdates(bot)

	for update := range updates {
		err = handler.Update(ctx, conn, bot, update)
		if err != nil {
			fmt.Println(err)
		}
	}
}
