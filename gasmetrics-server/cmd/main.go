package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

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

	mux.HandleFunc("GET /", server.GetReadings)
	mux.HandleFunc("GET /stats", server.GetStats)
	mux.HandleFunc("POST /", server.PostReading)
	mux.HandleFunc("DELETE /{id}", server.DeleteReading)

	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%v", port)
	log.Printf("gasmetrics-server listening at port %s\n", addr)
	if err = http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}
