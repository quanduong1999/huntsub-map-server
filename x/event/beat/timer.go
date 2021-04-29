package beat

import (
	"huntsub/huntsub-map-server/x/event"
)

var daily = event.NewHub(8)

func OnNewDay() *event.Line {
	return daily.NewLine()
}
