package weather

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const BaseUrl = "https://api.openweathermap.org/data/2.5/weather"

type OpenWeatherMapProvider struct {
	apiKey string
}

func (p *OpenWeatherMapProvider) Get(loc string) (Data, error) {
	queryUrl, err := url.Parse(BaseUrl)
	if err != nil {
		return Data{}, fmt.Errorf("failed to parse base url: %v", err)
	}

	// Code *inspired* by https://stackoverflow.com/questions/30652577/go-doing-a-get-request-and-building-the-querystring
	q := queryUrl.Query()
	q.Add("q", loc)
	q.Add("appid", p.apiKey)
	queryUrl.RawQuery = q.Encode()

	req, err := http.NewRequest(http.MethodGet, queryUrl.String(), bytes.NewReader([]byte{}))
	if err != nil {
		return Data{}, fmt.Errorf("failed to create net/http request: %v", err)
	}

	rep, err := http.DefaultClient.Do(req)
	if err != nil {
		return Data{}, fmt.Errorf("failed to make http request: %v", err)
	}
	defer rep.Body.Close()

	body, err := io.ReadAll(rep.Body)
	if err != nil {
		return Data{}, fmt.Errorf("failed to read body into byte array: %v", err)
	}

	if rep.StatusCode != http.StatusOK {
		return Data{}, fmt.Errorf("got non-200 status code: %d", rep.StatusCode)
	}

	var data Data
	if err := json.Unmarshal(body, &data); err != nil {
		return Data{}, fmt.Errorf("failed to unmarshal json: %v", err)
	}

	return data, nil
}

// NewOpenWeatherMapProvider creates an open weather map provider that is used to get the weather details
func NewOpenWeatherMapProvider(apiKey string) *OpenWeatherMapProvider {
	return &OpenWeatherMapProvider{apiKey: apiKey}
}
