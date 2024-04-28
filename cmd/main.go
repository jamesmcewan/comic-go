package main

import "fmt"

func main() {
	apiKey := getEnv("COMICVINE_KEY")

	result := getComics(getFullUrl(apiKey))
	fmt.Printf("%s \n", result)
}
