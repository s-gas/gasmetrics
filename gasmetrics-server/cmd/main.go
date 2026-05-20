package main

import (
	"os"
	"log"
	"fmt"
	"strconv"
	"context"
	"net/http"
	"encoding/json"
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

	mux := http.NewServeMux()

	mux.HandleFunc("GET /readings", func(w http.ResponseWriter, r *http.Request) {
		limit := 10
		if l := r.URL.Query().Get("limit"); l != "" {
			var err error
			limit, err = strconv.Atoi(l)
			if err != nil {
				http.Error(w, "invalid limit", http.StatusBadRequest)
				return
			}
		}
		readings, err := handler.Get(ctx, conn, limit)
		if err != nil {
			log.Println("error:", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(readings)
	})

	mux.HandleFunc("POST /readings", func(w http.ResponseWriter, r *http.Request) {
		var reading postgres.Reading
		if err := json.NewDecoder(r.Body).Decode(&reading); err != nil {
			http.Error(w, "invalid body", http.StatusBadRequest)
			return
		}
		if err := handler.Post(ctx, conn, reading.Value); err != nil {
			log.Println("error:", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
	})

	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%v", port)
	fmt.Printf("gasmetrics-server listening at %s%s\n", addr, "/readings")
	if err = http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}
