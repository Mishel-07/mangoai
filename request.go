package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func (m *Mango) DoRequest(endpoint string, method string, payload map[string]interface{}) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/%s", m.baseURL, endpoint)
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Error: Report https://github.com/Mishel-07/mangoai/issues")
	}

	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
