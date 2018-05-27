package httpjs

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

var myClient = &http.Client{Timeout: 60 * time.Second}

func GetJS(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	if r.StatusCode >= 400 {
		msg := fmt.Sprintf("request: %d", r.StatusCode)
		return errors.New(msg)
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}
