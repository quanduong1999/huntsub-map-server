package init2

import (
	"context"
	"db/mgo"
	"huntsub/huntsub-map-server/cache"
	"huntsub/huntsub-map-server/config"
	"huntsub/huntsub-map-server/emitcenter"
	"huntsub/huntsub-map-server/httpserver"
	"huntsub/huntsub-map-server/notification"
	"util/runtime"

	"github.com/golang/glog"
)

func initialize(ctx context.Context) {
	emitcenter.Start(ctx)
	notification.Start(ctx)
	cache.Start(ctx)
	mgo.Start(ctx)
}

func Start(ctx context.Context, p *config.ProjectConfig) {
	runtime.MaxProc()
	server = httpserver.NewProjectHttpServer(p)
	initialize(ctx)
}

func Wait() {
	defer beforeExit()
	cache.Wait()
	server.Wait()
	notification.Wait()
	emitcenter.Wait()
}

func beforeExit() {
	runtime.Recover()
	glog.Flush()
}
