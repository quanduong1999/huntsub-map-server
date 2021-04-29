package conf

import (
	"flag"
	"fmt"
	"os"
	"path"
)

/* glog
flag.BoolVar(&logging.toStderr, "logtostderr", false, "log to standard error instead of files")
flag.BoolVar(&logging.alsoToStderr, "alsologtostderr", false, "log to standard error as well as files")
flag.Var(&logging.verbosity, "v", "log level for V logs")
flag.Var(&logging.stderrThreshold, "stderrthreshold", "logs at or above this threshold go to stderr")
flag.Var(&logging.vmodule, "vmodule", "comma-separated list of pattern=N settings for file-filtered logging")
flag.Var(&logging.traceLocation, "log_backtrace_at", "when logging hits line file:N, emit a stack trace")
*/

type LogConfig struct {
	LogDir      string
	Verbosity   int
	LogToStdErr bool
}

func (lc LogConfig) Init() {
	var b = "false"
	if lc.LogToStdErr {
		b = "true"
	}
	flag.Lookup("alsologtostderr").Value.Set(b)
	if len(lc.LogDir) > 1 {
		var logdir = path.Join(".", lc.LogDir)
		if err := os.MkdirAll(logdir, os.ModeAppend); err != nil {
			fmt.Printf("Create log dir [%v] failed %v", logdir, err)
		}
		flag.Lookup("log_dir").Value.Set(logdir)
	}
	flag.Parse()
}

func (lc LogConfig) String() string {
	return fmt.Sprintf(
		"log:dir=%v,stderr=%v,verbosisty=%v",
		lc.LogDir, lc.LogToStdErr, lc.Verbosity,
	)
}
