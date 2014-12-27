package transport

import (
	"net/http"
	"net/url"
	"fmt"
	"strings"
	"io/ioutil"
	"encoding/json"
)

const (
	GooglePlacesAPIURL = "https://maps.googleapis.com/maps/api/place/nearbysearch/json"
)

type Place struct {
	Geometry struct {
		Location struct {
			Latitude float64 `json:"lat"`
			Longitude float64 `json:"lng"`
		} `json:"location"`
	} `json:"geometry"`
	Icon string `json:"icon"`
	ID string `json:"id"`
	Name string `json:"name"`
	PlaceID string `json:"place_id"`
	Reference string `json:"reference"`
	Scrope string `json:"scope"`
	Types []string `json:"types"`
	Vicinity string `json:"vicinity"`
	Photos []struct {
		Height int
		Width int
		HTMLAttributions []string `json:"html_attributions"`
		Reference string `json:"photo_reference"`
	} `json:"photos"`
	OpeningHours struct {
		OpenNow bool `json:"open_now"`
		WeekdayText []string `json:"weekday_text"`
	} `json:"opening_hours"`
	Rating float64 `json:"rating"`
	PriceLevel int `json:"price_level"`
	AlternativeIDs []struct {
		PlaceID string `json:"place_id"`
		Scope string `json:"scope"`
	} `json:"alt_ids"`
}

type Info struct {
	HTMLAttributions []string `json:"html_attributions"`
	NextPageToken string `json:"next_page_token"`
	Places []Place `json:"results"`
	Status string `json:"status"`
}

func GetClosestPlaces(apiKey string, lat, lng float64, types []string) (info *Info, err error) {
	info = new(Info)

	url := fmt.Sprintf("%v?key=%v&location=%v,%v&rankby=distance&types=%v", GooglePlacesAPIURL, apiKey, lat, lng, url.QueryEscape(strings.Join(types, "|")))
	response, err := http.Get(url)

	if err != nil {
		return
	}

	data, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return
	}

	err = json.Unmarshal(data, info)

	return
}
