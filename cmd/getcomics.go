package main

import (
	"io"
	"log"
	"net/http"
)

func getComics(fullUrl string) []byte {
	res, err := http.Get(fullUrl)
	if err != nil {
		log.Fatalln(err)
	}

	defer res.Body.Close()

	resBody, _ := io.ReadAll(res.Body)

	return resBody
}
