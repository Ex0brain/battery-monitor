package monitor

import (
	"github.com/distatus/battery"
	"github.com/maateen/battery-monitor/config"
	log "github.com/sirupsen/logrus"
)

// Message holds the information sent to channel
type Message struct {
	State        string
	Charge       int
	Notification string
}

// Monitor watches monitor and sends information to a channel.
func Monitor(messageChannel chan Message) {
	var lastState string
	var lastCharge int
	var lastNotification string
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
		currentState := states[bat.State]
		currentCharge := int(bat.Current / bat.Full * 100)

		// only one battery is supported
		if stateChanges(lastState, currentState) {
			lastState = currentState
			lastCharge = currentCharge
			lastNotification = currentState

			message := Message{
				State:        lastState,
				Charge:       lastCharge,
				Notification: lastNotification,
			}
			log.WithFields(log.Fields{
				"State":        lastState,
				"Charge":       lastCharge,
				"Notification": lastNotification,
			}).Info()

			messageChannel <- message // send message to channel
		} else if (currentState == "Discharging") && (lastCharge != currentCharge) {
			cfg := config.GetConfig()

			if currentCharge <= cfg.Critical {
				lastState = currentState
				lastCharge = currentCharge
				lastNotification = "Critical"
			} else if (currentCharge <= cfg.Low) && (lastNotification != "Low") {
				lastState = currentState
				lastCharge = currentCharge
				lastNotification = "Low"
			} else if (currentCharge <= cfg.First) && (lastNotification != "First") {
				lastState = currentState
				lastCharge = currentCharge
				lastNotification = "First"
			} else if (currentCharge <= cfg.Second) && (lastNotification != "Second") {
				lastState = currentState
				lastCharge = currentCharge
				lastNotification = "Second"
			} else if (currentCharge <= cfg.Third) && (lastNotification != "Third") {
				lastState = currentState
				lastCharge = currentCharge
				lastNotification = "Third"
			}

			message := Message{
				State:        lastState,
				Charge:       lastCharge,
				Notification: lastNotification,
			}
			log.WithFields(log.Fields{
				"State":        lastState,
				"Charge":       lastCharge,
				"Notification": lastNotification,
			}).Info()

			messageChannel <- message // send message to channel
		}
	}
}

func stateChanges(lastState string, currentState string) bool {
	if lastState != currentState {
		return true
	}
	return false
}
