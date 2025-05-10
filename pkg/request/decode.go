package request

import (
	"encoding/json"
	"io"
)

func Decode[T any](body io.ReadCloser) (T, error) {
	var loginRequest T
	err := json.NewDecoder(body).Decode(&loginRequest)
	if err != nil {
		return loginRequest, err
	}
	return loginRequest, nil
}
