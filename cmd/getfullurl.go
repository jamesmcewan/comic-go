package main

func getFullUrl(apiKey string, publisherId string) string {
	url := "https://comicvine.gamespot.com/api/publisher/"
	query := "&format=json&sort=name:asc"

	return url + publisherId + "/?api_key=" + apiKey + query
}
