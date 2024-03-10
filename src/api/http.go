package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

func GetRequest(getURL string, params url.Values) (result map[string]interface{}, err error) {
	BASE_URL := os.Getenv("TELEGRAM_BASEURL")
	BASE_URL += os.Getenv("TELEGRAM_TOKEN")

	// Construct the URL with query parameters
	requestURL, err := url.Parse(fmt.Sprintf("%s%s", BASE_URL, getURL))
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}

	requestURL.RawQuery = params.Encode()
	fmt.Println("============ request url", requestURL.String())

	request, err := http.NewRequest("GET", requestURL.String(), nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}

	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	// Close the response body.
	defer resp.Body.Close()

	// Read the response body.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// If body is empty.
	if len(body) == 0 {
		return nil, errors.New("empty body")
	}

	// Convert the body from json to struct format.
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return
}

func PostRequest(postURL string, body map[string]interface{}) (result map[string]interface{}, err error) {
	BASE_URL := os.Getenv("TELEGRAM_BASEURL")

	reqBody, err := json.Marshal(&body)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("GET", fmt.Sprintf("%s%s", BASE_URL, postURL), bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	client := &http.Client{}

	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	// Close the response body.
	defer resp.Body.Close()

	// Read the response body.
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// If body is empty.
	if len(respBody) == 0 {
		return nil, errors.New("empty body")
	}

	// Convert the body from json to struct format.
	err = json.Unmarshal(respBody, &result)
	if err != nil {
		return nil, err
	}

	return
}
