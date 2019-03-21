package iampigeon

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

// Client ...
type Client struct {
	*http.Client
	Host   string
	APIKey string
}

func (cl *Client) buildRequest(method, endpoint string, body io.Reader) (*http.Request, error) {
	// TODO: do the right url parse
	url := cl.Host + endpoint
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Api-Key", cl.APIKey)

	return req, nil
}

// Request ...
func (cl *Client) Request(method, endpoint string, m *Message) (*http.Response, error) {
	body := new(bytes.Reader)

	if m != nil {
		dataBytes := make(map[string]interface{})
		dataBytes["message"] = m
		bodyBytes, err := json.Marshal(dataBytes)
		if err != nil {
			return nil, err
		}

		body = bytes.NewReader(bodyBytes)
		if err != nil {
			return nil, err
		}
	}

	req, err := cl.buildRequest(method, endpoint, body)
	if err != nil {
		return nil, err
	}

	res, err := cl.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
