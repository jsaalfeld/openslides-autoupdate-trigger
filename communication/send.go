package communication

import (
	"bytes"
	"log"
	"net/http"
)

// SendPOST sends a POST Request with json as content to given url
func SendPOST(url string, json string) *http.Response {
	request, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(json)))
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		log.Panic(1, "Couldn't send POST request to: "+url)
		return nil
	}
	resp.Body.Close()
	return resp
}
