package main

import (
	"encoding/json"
	"fmt"
)

type ComicResponse struct {
	NumberOfPageResults int `json:"number_of_page_results"`
	Results             []struct {
		Aliases string `json:"aliases"`
		Volumes []struct {
			Name string `json:"name"`
			ID   int    `json:"id"`
			URL  string `json:"site_detail_url"`
		} `json:"volumes"`
	} `json:"results"`
}

func main() {
	apiKey := getEnv("COMICVINE_KEY")

	result := getComics(getFullUrl(apiKey, "4010-4212"))

	var comicResponseStruct ComicResponse
	json.Unmarshal(result, &comicResponseStruct)
	fmt.Printf("API response %+v\n", comicResponseStruct)
}
