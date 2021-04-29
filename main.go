package main

import (
	"context"
	"huntsub/huntsub-map-server/config"
	"huntsub/huntsub-map-server/init2"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	init2.Start(ctx, config.ReadConfig())
	init2.Wait()
}
