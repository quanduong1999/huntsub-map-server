package mlog

const (
	SolutionDir = "G:/Solution/Golang/src/did-eco/did-lucky/huntsub-map-server"
)

//IMLog : Log interface
type IMLog interface {
	//Info Group
	Debugf(level int, format string, args ...interface{})
	Debugln(level int, args ...interface{})

	//Info Group
	Infof(level int, format string, args ...interface{})
	Infoln(level int, args ...interface{})

	//Error Group
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Errorln(args ...interface{})
	ErrorDepth(level int, args ...interface{})
	ErrorStack(skip, depth int, args ...interface{})
	ErrorFullStack(args ...interface{})

	//Warning Group
	Warning(args ...interface{})
	Warningf(format string, args ...interface{})
	Warningln(args ...interface{})
	WarningDepth(level int, args ...interface{})
	WarningStack(skip, depth int, args ...interface{})
	WarningFullStack(args ...interface{})

	//Fatal Group
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
	Fatalln(args ...interface{})
	FatalDepth(level int, args ...interface{})
	FatalStack(skip, depth int, args ...interface{})
	FatalFullStack(args ...interface{})

	// State
	Off() IMLog
	EnableDebug() IMLog
}
