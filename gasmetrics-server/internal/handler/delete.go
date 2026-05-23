package handler

import (
	"log"
	"net/http"
	"strconv"
	"github.com/s-gas/hauslab/services/gasmetrics/gasmetrics-server/internal/postgres"
)

func (server *Server) DeleteReading(w http.ResponseWriter, r *http.Request) {
	ctx := server.Ctx
	conn := server.Conn
	var reading postgres.Reading
	var err error
	reading.Id, err = strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	if err := postgres.DeleteEntry(ctx, conn, reading); err != nil {
		log.Println("error:", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
