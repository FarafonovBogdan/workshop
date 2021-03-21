package jokes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"workshop/internal/api"
)

const getJokePath = "api?format=json"

type JokeClient struct {
	url string
}

func (jc *JokeClient) GetJoke() (*api.JokeResponse, error) {
	urlPath := jc.url + getJokePath
	r, err := http.Get(urlPath)
	if err != nil {
		return nil, nil
	} else if r.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request error: %v", err)
	}

	var data api.JokeResponse

	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
