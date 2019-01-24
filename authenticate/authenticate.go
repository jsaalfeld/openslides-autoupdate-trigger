package authenticate

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"../communication"
)

// Login structure for user login
type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

const loginURL = "/apps/users/login/"

// HandleLogin handles the login into openslides
func HandleLogin(user string, secret string, apiURL string) *http.Cookie {
	//Trim of the trailing /rest, because the login is at tld/users/login
	url := strings.Trim(apiURL, "/rest") + loginURL
	login := Login{"admin", "admin"}
	loginJSON, err := json.Marshal(login)
	if err != nil {
		log.Panic("Couldn't convert Login Object to Json")
		return nil
	}
	resp := communication.SendPOST(url, string(loginJSON))
	if resp.StatusCode >= 500 && resp.StatusCode < 600 {
		log.Output(1, "Login failed on Serviside with: "+strconv.Itoa(resp.StatusCode))
		return HandleLogin(user, secret, url)
	}
	if resp.StatusCode != 200 {
		log.Fatal("Login failed with errorcode: " + strconv.Itoa(resp.StatusCode))
	}
	log.Output(1, "Login as user "+user+" was successfull")
	for _, cookie := range resp.Cookies() {
		if cookie.Name == "OpenSlidesSessionID" {
			return cookie
		}
	}
	return nil
}
