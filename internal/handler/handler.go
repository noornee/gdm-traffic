package handler

import (
	"fmt"
	"io"
	"net/http"

	"github.com/noornee/gdm-traffic/internal/utils"
)

// Sends a request to the Google Maps Distance Matrix API to retrieve traffic data.
func FetchDistanceMatrixData(origin, destination, api_key string) ([]byte, error) {
	URL := fmt.Sprintf(
		"https://maps.googleapis.com/maps/api/distancematrix/json?origins=%s&destinations=%s&key=%s&departure_time=now&traffic_model=best_guess",
		origin, destination, api_key,
	)

	utils.InfoLog.Println("Fetching data from google maps api...")

	resp, err := http.Get(URL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// check status code
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Error fetching data. response status is %s", resp.Status)
	}

	// read response body
	body, err := io.ReadAll(resp.Body)

	return body, nil
}
