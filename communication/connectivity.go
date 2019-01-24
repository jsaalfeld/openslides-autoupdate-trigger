package communication

import (
	"net/http"
)

// TestConnectivity checks if a URL is reachable
func TestConnectivity(url string) bool {
	_, err := http.Get(url)
	if err != nil {
		return false
	}
	return true
}
