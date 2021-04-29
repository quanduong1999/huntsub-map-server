package notification

import (
	"context"
	"huntsub/huntsub-map-server/event"
	"huntsub/huntsub-map-server/x/mlog"
)

var notificationLog = mlog.NewTagLog("notification_event")
var forceRefreshChan = make(chan struct{}, 8)

func ForceRefresh() {
	forceRefreshChan <- struct{}{}
}

func autoRefresh(ctx context.Context) {
	ready()
	refreshContext, rcCancel := context.WithCancel(ctx)
	defer rcCancel()
	oev, oevCancel := event.ObjectEventSource.OnEvent()
	defer oevCancel()

	for {
		select {
		case v := <-oev:
			handleEvent(v)
		case <-refreshContext.Done():
			return
		}
	}
}

func Start(ctx context.Context) {
	go autoRefresh(ctx)
}
