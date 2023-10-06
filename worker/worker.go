package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"young-astrologer-service/storage"
)

type APODResponse struct {
	Date        string `json:"date"`
	Title       string `json:"title"`
	Explanation string `json:"explanation"`
}

func fetchAPODData(apiKey string, db *sql.DB) {
	today := time.Now().Format("2006-01-02")
	url := fmt.Sprintf("https://api.nasa.gov/planetary/apod?api_key=%s&date=%s", apiKey, today)

	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var apodData APODResponse
	err = json.Unmarshal(body, &apodData)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Сохранение данных в базе данных
	insertStatement := `
		INSERT INTO apod_metadata (date, title, explanation)
		VALUES ($1, $2, $3)
		ON CONFLICT (date) DO NOTHING
	`

	_, err = db.Exec(insertStatement, apodData.Date, apodData.Title, apodData.Explanation)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Данные успешно получены и сохранены в базе данных.")
}

func main() {
	storage.Init()
	storage.CreateTables()

	apiKey := os.Getenv("APOD_API_KEY")
	if apiKey == "" {
		fmt.Println("APOD_API_KEY не установлен.")
		os.Exit(1)
	}

	fetchAPODData(apiKey, storage.DB)
}
