package main

import (
	"fmt"
	"io"
	"net/http"
)

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
