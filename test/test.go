package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func main() {
	url := "https://sso.common.cloud.hpe.com/as/token.oauth2"

	headers := map[string]string{
		"Accept":       "*/*",
		"Content-Type": "application/x-www-form-urlencoded",
	}

	payload := map[string]string{
		"grant_type":    "client_credentials",
		"client_id":     "c1eaf55c-a6e1-42ca-9bd4-ebb6eff6b3c4",
		"client_secret": "b03bb5b28db111ed9deb8a46f282d903",
	}
	payloadStr := ""
	for key, value := range payload {
		payloadStr += key + "=" + value + "&"
	}
	payloadStr = payloadStr[:len(payloadStr)-1]

	client := &http.Client{}

	// Create a POST request with the URL, headers, and payload
	req, err := http.NewRequest("POST", url, bytes.NewBufferString(payloadStr))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	responseBody := new(bytes.Buffer)
	_, err = responseBody.ReadFrom(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	// Print the response status code and body
	fmt.Println("Response Status Code:", resp.Status)
	fmt.Println("Response Body:", responseBody.String())
}
