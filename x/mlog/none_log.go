package mlog

//NoLog :
var NoLog = &noneLog{}

type noneLog struct {
}

//------------------------------	INFO	---------------------------->
func (n *noneLog) Debugf(level int, format string, args ...interface{}) {
}

func (n *noneLog) Debugln(level int, args ...interface{}) {

}

//------------------------------	INFO	---------------------------->
func (n *noneLog) Infof(level int, format string, args ...interface{}) {
}

func (n *noneLog) Infoln(level int, args ...interface{}) {

}

//------------------------------	ERROR	---------------------------->
func (n *noneLog) Error(args ...interface{}) {

}

func (n *noneLog) Errorf(format string, args ...interface{}) {
}

func (n *noneLog) Errorln(args ...interface{}) {

}

func (n *noneLog) ErrorDepth(level int, args ...interface{}) {

}

func (n *noneLog) ErrorStack(skip, depth int, args ...interface{}) {

}

func (n *noneLog) ErrorFullStack(args ...interface{}) {

}

//------------------------------	WARNING	---------------------------->

func (n *noneLog) Warning(args ...interface{}) {

}

func (n *noneLog) Warningf(format string, args ...interface{}) {
}

func (n *noneLog) Warningln(args ...interface{}) {

}

func (n *noneLog) WarningDepth(level int, args ...interface{}) {

}

func (n *noneLog) WarningStack(skip, depth int, args ...interface{}) {

}

func (n *noneLog) WarningFullStack(args ...interface{}) {

}

//------------------------------	FATAL	---------------------------->

func (n *noneLog) Fatal(args ...interface{}) {
}

func (n *noneLog) Fatalf(format string, args ...interface{}) {
}

func (n *noneLog) Fatalln(args ...interface{}) {
}

func (n *noneLog) FatalDepth(level int, args ...interface{}) {
}

func (n *noneLog) FatalStack(skip, depth int, args ...interface{}) {

}

func (n *noneLog) FatalFullStack(args ...interface{}) {

}

func (n *noneLog) Off() IMLog {
	return n
}

func (n *noneLog) EnableDebug() IMLog {
	return n
}
