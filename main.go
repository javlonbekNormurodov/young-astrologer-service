package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	api "young-astrologer-service/api"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := mux.NewRouter()
	r.HandleFunc("/api/album", api.GetAllAPODRecords).Methods("GET")
	r.HandleFunc("/api/record/{date}", api.GetAPODRecordForDate).Methods("GET")

	fmt.Printf("Сервер запущен на порту %s\n", port)
	http.ListenAndServe(":"+port, r)
}
