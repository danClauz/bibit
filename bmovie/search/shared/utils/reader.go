package utils

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

func ReadAll(body io.Reader, out interface{}) error {
	b, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(b, out); err != nil {
		return err
	}

	return nil
}

