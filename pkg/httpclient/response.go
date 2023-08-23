package httpclient

import (
	"encoding/json"
	"io"
)

type Response struct {
	StatusCode int
	Body       io.ReadCloser
}

func (r Response) Decode(dest interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(dest); err != nil {
		return err
	}

	return nil
}
