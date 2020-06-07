package monitor

import (
	"github.com/distatus/battery"
	log "github.com/sirupsen/logrus"
)

// Message holds the information sent to channel
type Message struct {
	State string
	Charge int
}

// Monitor watches monitor and sends information to a channel.
func Monitor(messageChannel chan Message) {
	var lastState battery.State
	var states = []string{
		"Unknown",
		"Empty",
		"Full",
		"Charging",
		"Discharging",
	}

	log.Info("Monitoring the battery.")

	for {
		batteries, _ := battery.GetAll()
		bat := batteries[0]

		// only one battery is supported
		if stateChanges(lastState, bat.State) {
			lastState = bat.State
			message := Message{
				State: states[bat.State],
				Charge: int(bat.Current/bat.Full*100),
			}
			log.WithFields(log.Fields{
				"State": message.State,
				"Charge": message.Charge,
			}).Info()

			messageChannel <- message // send message to channel
		}
	}
}

func stateChanges(lastState battery.State, currentState battery.State) bool {
	if lastState != currentState {
		return true
	} else {
		return false
	}
}