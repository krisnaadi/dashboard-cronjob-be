package integration

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"reflect"

	"github.com/tech-djoin/wallet-djoin-service/internal/pkg/writer"
)

// Example values for URL, request body, and headers
// url := "https://api.example.com/hello"
//
//	body := map[string]interface{}{
//		"key1": "value1",
//		"key2": 123,
//		"key3": true,
//	}
//
//	headers := map[string]string{
//		"Content-Type":  "application/json",
//		"Authorization": "Bearer your-token",
//	}
//
// Call connectToGlobalAPI with example values
// resp, err := connectToGlobalAPI(url, body, headers)
//
//	if err != nil {
//		Handle error
//		return
//	}
//
// connectToGlobalAPI sends a POST request to the global API endpoint and returns the API response.
func SendRequestToPpobServer(url string, method string, body map[string]interface{}, headers map[string]string) (writer.Response, error) {
	reqBody, err := json.Marshal(body)
	if err != nil {
		return writer.Response{}, err
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(reqBody))
	if err != nil {
		return writer.Response{}, err
	}

	// Set headers
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	clientResponse := &http.Client{}
	resp, err := clientResponse.Do(req)
	if err != nil {
		return writer.Response{}, err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return writer.Response{}, err
	}

	var jsonResponse map[string]interface{}
	err = json.Unmarshal(respBody, &jsonResponse)
	if err != nil {
		return writer.Response{}, err
	}

	status := false
	statusResponse := reflect.ValueOf(jsonResponse["status"]).Kind() == reflect.Bool
	if statusResponse {
		status = jsonResponse["status"].(bool)
	}

	statusResponse = reflect.ValueOf(jsonResponse["status"]).Kind() == reflect.String
	if statusResponse {
		if jsonResponse["status"] == "success" {
			status = true
		}
	}

	// Assert the "message" value as a string
	message, ok := jsonResponse["message"].(string)
	if !ok {
		return writer.Response{}, errors.New("Failed to parse message")
	}

	if !status {
		return writer.Response{
			Status:  false,
			Message: message,
		}, nil
	}

	// Create an APIResponse object from the JSON data
	apiResponse := writer.Response{
		Status:  status,
		Message: message,
		Data:    jsonResponse["data"],
		// Set other fields as needed
	}

	return apiResponse, nil
}
