package main

import (
	"os"
	"log"
	"fmt"
	"context"
	"net/http"
	"encoding/json"
	"github.com/s-gas/hauslab/services/gasmetrics/gasmetrics-server/internal/handler"
	"github.com/s-gas/hauslab/services/gasmetrics/gasmetrics-server/internal/postgres"
)

type Reading struct {
	Value int `json:"value"`
}

func main() {
	ctx := context.Background()
	conn, err := postgres.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(context.Background())

	var reading Reading
	endpoint := "/readings"
	http.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			log.Println("error: method not allowed")
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		if err := json.NewDecoder(r.Body).Decode(&reading); err != nil {
			log.Println("error: invalid body")
			http.Error(w, "invalid body", http.StatusBadRequest)
			return
		}
		if err := handler.StoreValue(ctx, conn, reading.Value); err != nil {
			log.Println("error:", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	})

	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%v", port)
	fmt.Printf("gasmetrics-server listening at %s%s\n", addr, endpoint)
	err = http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal(err)
	}
}
