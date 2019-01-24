package trigger

import (
	"log"
	"strconv"
)

const medium = 500
const high = 1000
const low = 50

// Run is the Entrypoint for the trigger
// It takes the session ID from successfull login
// stresses for activeSeconds
// rests for inactiveSeconds
// stresses on actionLevel
// will repeat until it is killed
func Run(openSlidesSessionID string, apiURL string, activeSeconds int, inactiveSeconds int, actionLevel string) {
	log.Output(1, "sessionID: "+openSlidesSessionID)
	log.Output(1, "activeSeconds: "+strconv.Itoa(activeSeconds))
	log.Output(1, "inactiveSeconds: "+strconv.Itoa(inactiveSeconds))
	log.Output(1, "Creating some Motions...")
	// Create about 5 Motions and then play with them ;)
}
