package runtime

import (
	"github.com/golang/glog"
	"runtime/debug"
)

func Recover() {
	if r := recover(); r != nil {
		glog.Error(r, string(debug.Stack()))
	}
}
