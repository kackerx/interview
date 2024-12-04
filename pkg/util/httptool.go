package util

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

func Post(ctx context.Context, url string, payload any) ([]byte, error) {
	bs, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	data := strings.NewReader(string(bs))
	req, err := http.NewRequest(http.MethodPost, url, data)

	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer sk-IbpZq8zcmT1plBAWBa65148eA4F24fE092Ee1bEd2a0462E9")
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
