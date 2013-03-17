package ip

import (
	"net/http"
	"io/ioutil"
)

func getContent(url string) ([]byte, error) {
	// Send the request
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	// Read the content into a byte array
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err 
	}

	// Return the bytes as a buffer
	return body, nil
}