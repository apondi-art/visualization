// internal/api/client.go
package api

import (
	"encoding/json"
	"net/http"

	"learn.zone01kisumu.ke/git/quochieng/groupie-tracker/internal/models"
	// "groupie-tracker/internal/models"
)

const (
	ArtistsURL = "https://groupietrackers.herokuapp.com/api/artists"
)

func FetchJSON(url string, target interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(target)
}

func FetchAllData() ([]models.CombinedData, error) {
	var artists []models.Artist

	if err := FetchJSON(ArtistsURL, &artists); err != nil {
		return nil, err
	}

	combinedData := make([]models.CombinedData, len(artists))
	for i, artist := range artists {
		combinedData[i] = models.CombinedData{
			Artist: artist,
		}
	}

	return combinedData, nil
}
