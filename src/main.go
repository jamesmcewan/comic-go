package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func getEnv(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("env file missing")
	}

	return os.Getenv(key)
}

func getFullUrl(apiKey string) string {
	url := "https://comicvine.gamespot.com/api/volumes/?api_key="
	query := "&format=json&sort=name:asc&filter=name:Walking%20Dead"

	return url + apiKey + query
}

func getComics(fullUrl string) []byte {
	res, err := http.Get(fullUrl)
	if err != nil {
		fmt.Println("oh no")
	}
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("oh no")
	}

	return resBody
}

func main() {
	apiKey := getEnv("COMICVINE_KEY")

	result := getComics(getFullUrl(apiKey))
	fmt.Printf("%s \n", result)
}
