package service

import (
	"log"
	"strconv"

	"../communication"
)

// Start starts the service to penetrate OpenSlides
func Start(host string, username string, password string, activeSeconds int, inactiveSeconds int, actionLevel string) {
	if !communication.TestConnectivity(host) {
		log.Panic("The URL " + host + " is not the API Endpoint")
	}
	var login = HandleLogin(username, password, host)
	log.Output(1, "Host: "+host)
	log.Output(1, "API URL: "+host)
	log.Output(1, "LoginAnswer: "+login.Value)
	log.Output(1, "Username: "+username)
	log.Output(1, "Password: "+password)
	log.Output(1, "Active Seconds: "+strconv.Itoa(activeSeconds))
	log.Output(1, "Inactive Seconds: "+strconv.Itoa(inactiveSeconds))
	log.Output(1, "Action Level: "+actionLevel)
}
