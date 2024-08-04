package yelp

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const yelpBaseURL = "https://api.yelp.com/v3"

var yelpAPIKey = os.Getenv("YELP")

type SearchResponse struct {
	Businesses []Business `json:"businesses"`
}

type Business struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Rating float64 `json:"rating"`
	// Add other fields as needed
}

func SearchBusinesses(zipCode int) ([]Business, error) {
	url := fmt.Sprintf("%s/businesses/search?location=%d&open_now=true", yelpBaseURL, zipCode)
	return makeYelpRequest(url)
}

func GetBusiness(storeID string) (*Business, error) {
	url := fmt.Sprintf("%s/businesses/%s", yelpBaseURL, storeID)
	businesses, err := makeYelpRequest(url)
	if err != nil {
		return nil, err
	}
	if len(businesses) == 0 {
		return nil, fmt.Errorf("no business found")
	}
	return &businesses[0], nil
}

func makeYelpRequest(url string) ([]Business, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+yelpAPIKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var searchResp SearchResponse
	err = json.Unmarshal(body, &searchResp)
	if err != nil {
		return nil, err
	}

	return searchResp.Businesses, nil
}
