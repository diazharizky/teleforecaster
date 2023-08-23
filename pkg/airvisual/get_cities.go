package airvisual

import (
	"log"
)

type City struct {
	City string `json:"city"`
}

func (c Client) GetCities(country, state string) (data []City, err error) {
	params := map[string]string{
		"country": country,
		"state":   state,
		"key":     "c0c7b2e0-69a4-4b69-8a92-70792efae702",
	}

	resp, err := c.httpClient.Get("/cities", params)
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
