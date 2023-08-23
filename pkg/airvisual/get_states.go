package airvisual

import (
	"log"
)

type State struct {
	State string `json:"state"`
}

func (c Client) GetStates(country string) (data []State, err error) {
	params := map[string]string{
		"country": country,
		"key":     "c0c7b2e0-69a4-4b69-8a92-70792efae702",
	}

	resp, err := c.httpClient.Get("/states", params)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err = resp.Body.Close(); err != nil {
			log.Printf("Error unable to close response's body: %v\n", err)
		}
	}()

	if resp.StatusCode != 200 {
		return nil, nil
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
