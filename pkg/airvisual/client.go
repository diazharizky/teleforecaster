package airvisual

import (
	"log"

	"github.com/diazharizky/teleforecaster/pkg/httpclient"
)

type Client struct {
	httpClient httpclient.Client
}

func New() Client {
	c := Client{
		httpClient: httpclient.New(httpclient.ClientConfig{
			BaseURL: "http://api.airvisual.com/v2",
		}),
	}

	return c
}

func (c Client) GetCityData(city, state, country string) (data *CityData, err error) {
	params := map[string]string{
		"city":    city,
		"state":   state,
		"country": country,
		"key":     "c0c7b2e0-69a4-4b69-8a92-70792efae702",
	}

	resp, err := c.httpClient.Get("/city", params)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err = resp.Body.Close(); err != nil {
			log.Printf("Error unable to close response's body: %v\n", err)
		}
	}()

	var baseResponse BaseResponse
	if err = resp.Decode(&baseResponse); err != nil {
		return nil, err
	}

	if err = baseResponse.Decode(&data); err != nil {
		return nil, err
	}

	return
}
