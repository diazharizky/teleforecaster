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
		"key":     "",
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
