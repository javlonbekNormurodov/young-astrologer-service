package api

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"young-astrologer-service/storage"

	"github.com/gorilla/mux"
)

type APODMetadata struct {
	Date        string `json:"date"`
	Title       string `json:"title"`
	Explanation string `json:"explanation"`
}

func GetAllAPODRecords(w http.ResponseWriter, r *http.Request) {
	var metadataList []APODMetadata
	rows, err := storage.DB.Query("SELECT date, title, explanation FROM apod_metadata ORDER BY date DESC")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var metadata APODMetadata
		err := rows.Scan(&metadata.Date, &metadata.Title, &metadata.Explanation)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		metadataList = append(metadataList, metadata)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(metadataList)
}

func GetAPODRecordForDate(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	date := params["date"]

	var metadata APODMetadata
	err := storage.DB.QueryRow("SELECT date, title, explanation FROM apod_metadata WHERE date = $1", date).Scan(&metadata.Date, &metadata.Title, &metadata.Explanation)
	if err != nil {
		if err == sql.ErrNoRows {
			http.NotFound(w, r)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(metadata)
}
