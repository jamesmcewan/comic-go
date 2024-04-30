package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type ComicVineResponse struct {
	Results struct {
		Aliases string `json:"aliases"`
		Volumes []struct {
			Name          string `json:"name"`
			SiteDetailURL string `json:"site_detail_url"`
		} `json:"volumes"`
	} `json:"results"`
}

func main() {
	apiKey := getEnv("COMICVINE_KEY")
	publisherId := os.Args[1]
	fmt.Println("Searching for Publisher ID:", publisherId)
	fullURL := getFullUrl(apiKey, publisherId)

	client := http.Client{}
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:47.0) Gecko/20100101 Firefox/47.0")

	response, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		fmt.Println("Response OK")
		var comicVineResp ComicVineResponse
		err := json.NewDecoder(response.Body).Decode(&comicVineResp)
		if err != nil {
			fmt.Println("Error decoding JSON:", err)
			return
		}

		fmt.Println("Result found")
		fmt.Println("--------------------------------")
		fmt.Printf("Name: %v\n\n", comicVineResp.Results.Aliases)
		fmt.Println("--------------------------------")
		for _, volume := range comicVineResp.Results.Volumes {
			fmt.Printf("%v\n", volume.Name)
			fmt.Printf("%v\n", volume.SiteDetailURL)
			fmt.Println("--------------------------------")
		}
	} else {
		fmt.Printf("Request failed with status code: %d", response.StatusCode)
	}
}
