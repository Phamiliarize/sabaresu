package util

import (
	"encoding/json"
	"net/http"
)

func Body(r *http.Request) (map[string]interface{}, error) {
	switch r.Method {
	case "POST":
		var body map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			return nil, err
		}
		return body, nil
	case "PUT":
		var body map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			return nil, err
		}
		return body, nil
	case "PATCH":
		var body map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			return nil, err
		}
		return body, nil
	default:
		break
	}

	return nil, nil
}
