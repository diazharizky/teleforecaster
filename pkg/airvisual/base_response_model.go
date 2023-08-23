package airvisual

import (
	"encoding/json"
)

type BaseResponse struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

func (r BaseResponse) Decode(dest interface{}) error {
	jsonStr, err := json.Marshal(r.Data)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(jsonStr, dest); err != nil {
		return err
	}

	return nil
}
