package airvisual

import (
	"log"
	"os"
)

type City struct {
	City string `json:"city"`
}

func (c Client) GetCities(country, state string) (data []City, err error) {
	params := map[string]string{
		"country": country,
		"state":   state,
		"key":     os.Getenv("AIR_VISUAL_KEY"),
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

	if resp.StatusCode != 200 {
		return []City{}, nil
	}

	var baseResponse BaseResponse
	if err = resp.Decode(&baseResponse); err != nil {
		return nil, err
	}

	if err = baseResponse.Decode(&data); err != nil {
		return nil, err
	}

	return
}
