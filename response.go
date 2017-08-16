package nfs

import (
	"encoding/json"
)

func readResponseArray(resp string) ([]string, error) {
	// Parse json array
	var arr = make([]string, 0)
	err := json.Unmarshal([]byte(resp), &arr)
	if err != nil {
		return nil, err
	}
	return arr, nil
}
