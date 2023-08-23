package airvisual

import (
	"github.com/diazharizky/teleforecaster/pkg/httpclient"
)

type Client struct {
	httpClient httpclient.Client
}

func New() Client {
	return Client{
		httpClient: httpclient.New(httpclient.ClientConfig{
			BaseURL: "http://api.airvisual.com/v2",
		}),
	}
}
