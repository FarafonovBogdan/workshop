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

func NewJokeClient(baseUrl string) *JokeClient {
	return &JokeClient{
		url: baseUrl,
	}
}

func (jc *JokeClient) GetJoke() (*api.JokeResponse, error) {
	urlPath := jc.url + getJokePath
	r, err := http.Get(urlPath)
	if err != nil {
		return nil, err
	} else if r.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request status: %s", http.StatusText(r.StatusCode))
	}

	var data api.JokeResponse

	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
