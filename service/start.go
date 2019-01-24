package service

import (
	"log"

	"../authenticate"
	trigger "../autoupdate-trigger"
	"../communication"
)

// Start starts the service to penetrate OpenSlides
func Start(host string, username string, password string, activeSeconds int, inactiveSeconds int, actionLevel string) {
	if !communication.TestConnectivity(host) {
		log.Panic("The URL " + host + " is not the API Endpoint")
	}
	var login = authenticate.HandleLogin(username, password, host)
	trigger.Run(login.Value, host, activeSeconds, inactiveSeconds, actionLevel)
}
