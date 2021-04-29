package runtime

import (
	"runtime"
)

func MaxProc() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}
