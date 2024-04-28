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

func main() {
	apiKey := getEnv("COMICVINE_KEY")
	url := "https://comicvine.gamespot.com/api/volumes/?api_key="
	query := "&format=json&sort=name:asc&filter=name:Walking%20Dead"
	reqUrl := url + apiKey + query
	res, err := http.Get(reqUrl)
	if err != nil {
		fmt.Printf("error %s \n", err)
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("error %s \n", err)
	}

	fmt.Printf("%s \n", resBody)
}
