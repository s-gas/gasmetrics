package handler

import (
	"log"
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/s-gas/hauslab/services/gasmetrics/gasmetrics-server/internal/postgres"
)

func (server *Server) GetReadings(w http.ResponseWriter, r *http.Request) {
	ctx := server.Ctx
	conn := server.Conn
	limit := 10
	if l := r.URL.Query().Get("limit"); l != "" {
		var err error
		limit, err = strconv.Atoi(l)
		if err != nil {
			http.Error(w, "invalid limit", http.StatusBadRequest)
			return
		}
	}
	entries, err := postgres.GetEntries(ctx, conn, limit)
	if err != nil {
		log.Println("error:", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(entries)
}

func (server *Server) GetStats(w http.ResponseWriter, r *http.Request) {
	ctx := server.Ctx
	conn := server.Conn
	avg, err := postgres.GetAverage(ctx, conn)
	if err != nil {
		log.Println("error:", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(avg)
}
