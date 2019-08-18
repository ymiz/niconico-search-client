package common

import (
	"encoding/json"
	"net/http"
	"net/url"
)

func Get(endPoint, userAgent string, values url.Values) (*Response, error) {
	method := "GET"
	params := ""
	if values != nil {
		params = "?" + values.Encode()
	}
	uri := endPoint + params
	req, err := http.NewRequest(method, uri, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", userAgent)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	var response Response
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
