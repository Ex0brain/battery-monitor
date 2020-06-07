package notification

import (
	"fmt"
	"github.com/gen2brain/beeep"
	"github.com/maateen/battery-monitor/internal/monitor"
	log "github.com/sirupsen/logrus"
)

func Notify(message monitor.Message) {
	title := fmt.Sprintf("Battery is %s", message.State)
	body := fmt.Sprintf("Current charge level: %d%%", message.Charge)
	icon := fmt.Sprintf("assets/%s.png", message.State)

	err := beeep.Notify(title, body, icon)
	if err != nil {
		log.Fatal(err)
	}
}