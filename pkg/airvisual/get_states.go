package airvisual

import (
	"log"
	"net/http"
	"os"
)

type State struct {
	State string `json:"state"`
}

func (c Client) GetStates(country string) (data []State, err error) {
	params := map[string]string{
		"country": country,
		"key":     os.Getenv("AIR_VISUAL_KEY"),
	}

	resp, err := c.httpClient.Get("/states", params)
	if err != nil {
		return nil, err
	}
	defer func() {
		if defErr := resp.Body.Close(); defErr != nil {
			log.Printf("Error unable to close response's body: %v\n", defErr)
		}
	}()

	statusCode := resp.StatusCode
	if statusCode != http.StatusOK {
		if statusCode == http.StatusBadRequest {
			return nil, StateNotSupportedError
		}
		return nil, RateLimitError
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
