package main

import (
	"os"
	"log"
	"fmt"
	"context"
	"net/http"

	"github.com/s-gas/hauslab/services/gasmetrics/gasmetrics-server/internal/handler"
	"github.com/s-gas/hauslab/services/gasmetrics/gasmetrics-server/internal/postgres"
)

func main() {
	ctx := context.Background()
	conn, err := postgres.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(context.Background())

	server := &handler.Server{Ctx: ctx, Conn: conn}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /readings", server.GetReadings)
	mux.HandleFunc("POST /readings", server.PostReading)
	mux.HandleFunc("DELETE /readings/{id}", server.DeleteReading)

	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%v", port)
	fmt.Printf("gasmetrics-server listening at %s%s\n", addr, "/readings")
	if err = http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}
