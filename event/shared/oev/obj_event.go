package oev

import (
	"huntsub/huntsub-map-server/x/math"
	"time"
)

var objectOrigin = math.RandString("oev", 20)

type ObjectEvent struct {
	Action   ObjectActionName   `json:"act"`
	Category ObjectCategoryName `json:"cat"`
	Data     interface{}        `json:"data"`
	Origin   string             `json:"ogirin"`
	Ctime    int64              `json:"ctime"`
}

func NewObjectEvent(action ObjectActionName, data interface{}) *ObjectEvent {
	var o = &ObjectEvent{
		Action:   action,
		Category: GetCategory(data),
		Data:     data,
		Origin:   objectOrigin,
		Ctime:    time.Now().Unix(),
	}
	return o
}
