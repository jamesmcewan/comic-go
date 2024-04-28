package main

func getFullUrl(apiKey string) string {
	url := "https://comicvine.gamespot.com/api/volumes/?api_key="
	query := "&format=json&sort=name:asc&filter=name:Walking%20Dead"

	return url + apiKey + query
}
