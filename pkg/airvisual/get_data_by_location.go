package airvisual

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func (c Client) GetDataByLocation(lat, lng float32) (data *CityData, err error) {
	params := map[string]string{
		"lat": fmt.Sprintf("%f", lat),
		"lon": fmt.Sprintf("%f", lng),
		"key": os.Getenv("AIR_VISUAL_KEY"),
	}

	resp, err := c.httpClient.Get("/nearest_city", params)
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
