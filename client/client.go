package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/agpelkey/microserv-prac/types"
)

type Client struct {
	endpoint string
}

func New(endpoint string) *Client {
	return &Client{
		endpoint: endpoint,
	}
}

func (c *Client) FetchPrice(ctx context.Context, ticker string) (*types.PriceResponse, error) {

	endpoint := fmt.Sprintf("%s?ticker=%s", c.endpoint, ticker)

	request, err := http.NewRequest("get", endpoint, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		httpErr := map[string]any{}
		if err := json.NewDecoder(request.Body).Decode(&httpErr); err != nil {
			return nil, err 
		}
		return nil, fmt.Errorf("Service responded with non OK status code: %s", httpErr)
	}

	priceResp := new(types.PriceResponse)
	if err := json.NewDecoder(resp.Body).Decode(priceResp); err != nil {
		return nil, err
	}

	return priceResp, nil
}
