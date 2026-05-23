package handler

import (
	"log"
	"net/http"
	"encoding/json"
	"github.com/s-gas/hauslab/services/gasmetrics/gasmetrics-server/internal/postgres"
)

func (server *Server) PostReading(w http.ResponseWriter, r *http.Request) {
	ctx := server.Ctx
	conn := server.Conn
	var reading postgres.Reading
	if err := json.NewDecoder(r.Body).Decode(&reading); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}
	lastEntry, err := postgres.GetPreviousEntry(ctx, conn, reading)
	if err != nil {
		log.Println("error:", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	nextEntry, err := postgres.GetNextEntry(ctx, conn, reading)
	if err != nil {
		log.Println("error:", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = validate(reading.Value, lastEntry, nextEntry)
	if err != nil {
		log.Println("error:", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = postgres.AddEntry(ctx, conn, reading)
	if err != nil {
		log.Println("error:", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
