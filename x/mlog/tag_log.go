package mlog

import (
	"fmt"
	"runtime"
	"strings"

	"runtime/debug"

	"github.com/golang/glog"
)

type tagLog struct {
	ProjectPath string
	Tag         string
	Debug       bool
}

//NewTagLog : Create new tag log
func NewTagLog(tag string) IMLog {
	if IsSkip(tag) {
		return NoLog
	}
	return &tagLog{
		Tag:         "[" + tag + "] ",
		ProjectPath: SolutionDir,
	}
}

//------------------------------	Util	---------------------------->
func (t *tagLog) splitFilePath(path *string) {
	index := strings.Index(*path, t.ProjectPath)
	if index != -1 {
		*path = (*path)[len(t.ProjectPath):]
	}
}

func (t *tagLog) splitFuncPath(path *string) {
	index := strings.LastIndex(*path, ".")
	if index != -1 {
		*path = (*path)[index+1:]
	}
}

//------------------------------	INFO	---------------------------->
func (t *tagLog) Debugf(level int, format string, args ...interface{}) {
	if t.Debug {
		format = t.Tag + "  " + format
		format = fmt.Sprintf(format, args...)
		glog.InfoDepth(level+1, format)
	}
}

func (t *tagLog) Debugln(level int, args ...interface{}) {
	if t.Debug {
		args = append([]interface{}{t.Tag}, args...)
		glog.InfoDepth(level+1, args...)
	}
}

func (t *tagLog) EnableDebug() IMLog {
	t.Debug = true
	return t
}

//------------------------------	INFO	---------------------------->
func (t *tagLog) Infof(level int, format string, args ...interface{}) {
	format = t.Tag + "  " + format
	format = fmt.Sprintf(format, args...)
	glog.InfoDepth(level+1, format)
}

func (t *tagLog) Infoln(level int, args ...interface{}) {
	args = append([]interface{}{t.Tag}, args...)
	glog.InfoDepth(level+1, args...)
}

//------------------------------	ERROR	---------------------------->

func (t *tagLog) Error(args ...interface{}) {
	args = append([]interface{}{t.Tag}, args...)
	glog.ErrorDepth(1, args...)
}

func (t *tagLog) Errorf(format string, args ...interface{}) {
	format = t.Tag + "  " + format
	format = fmt.Sprintf(format, args...)
	glog.ErrorDepth(1, format)
}

func (t *tagLog) Errorln(args ...interface{}) {
	args = append([]interface{}{t.Tag}, args...)
	glog.ErrorDepth(1, args...)
}

func (t *tagLog) ErrorDepth(level int, args ...interface{}) {
	args = append([]interface{}{t.Tag}, args...)
	glog.ErrorDepth(level+1, args...)
}

func (t *tagLog) ErrorStack(skip, depth int, args ...interface{}) {
	var pc = make([]uintptr, depth)
	runtime.Callers(skip+1, pc)
	var msgs = make([]string, 0)
	var s string
	for _, p := range pc {
		if p != 0 {
			file, line := runtime.FuncForPC(p).FileLine(p)
			t.splitFilePath(&file)
			s = fmt.Sprintf("%s:%d:%s", file, line, runtime.FuncForPC(p).Name())
			msgs = append(msgs, s)
		} else {
			break
		}
	}
	s = ""
	for i := len(msgs) - 1; i >= 0; i-- {
		if i == 0 {
			s += msgs[i]
		}
		s += msgs[i] + "  ->  "
	}
	args = append([]interface{}{s}, args)
	t.Errorln(args...)
}

func (t *tagLog) ErrorFullStack(args ...interface{}) {
	var msg = string(debug.Stack())
	args = append([]interface{}{msg}, args)
	t.ErrorDepth(1, args...)
}

//------------------------------	WARNING	---------------------------->

func (t *tagLog) Warning(args ...interface{}) {
	args = append([]interface{}{t.Tag}, args...)
	glog.WarningDepth(1, args...)
}

func (t *tagLog) Warningf(format string, args ...interface{}) {
	format = t.Tag + "  " + format
	format = fmt.Sprintf(format, args...)
	glog.WarningDepth(1, format)
}

func (t *tagLog) Warningln(args ...interface{}) {
	args = append([]interface{}{t.Tag}, args...)
	args = append(args, "\r\n")
	glog.WarningDepth(1, args...)
}

func (t *tagLog) WarningDepth(level int, args ...interface{}) {
	args = append([]interface{}{t.Tag}, args...)
	glog.WarningDepth(level+1, args...)
}

func (t *tagLog) WarningStack(skip, depth int, args ...interface{}) {
	var pc = make([]uintptr, depth)
	runtime.Callers(skip+1, pc)
	var msgs = make([]string, 0)
	var s string
	for _, p := range pc {
		if p != 0 {
			file, line := runtime.FuncForPC(p).FileLine(p)
			t.splitFilePath(&file)
			s = fmt.Sprintf("%s:%d:%s", file, line, runtime.FuncForPC(p).Name())
			msgs = append(msgs, s)
		} else {
			break
		}
	}
	s = ""
	for i := len(msgs) - 1; i >= 0; i-- {
		if i == 0 {
			s += msgs[i]
		}
		s += msgs[i] + "  ->  "
	}
	t.Warningln(s)
}

func (t *tagLog) WarningFullStack(args ...interface{}) {
	var msg = string(debug.Stack())
	args = append([]interface{}{msg}, args)
	t.Warningln(args...)
}

//------------------------------	FATAL	---------------------------->

func (t *tagLog) Fatal(args ...interface{}) {
	args = append([]interface{}{t.Tag}, args...)
	glog.FatalDepth(1, args...)
}

func (t *tagLog) Fatalf(format string, args ...interface{}) {
	format = t.Tag + "  " + format
	format = fmt.Sprintf(format, args...)
	glog.FatalDepth(1, format)
}

func (t *tagLog) Fatalln(args ...interface{}) {
	args = append([]interface{}{t.Tag}, args...)
	args = append(args, "\r\n")
	glog.FatalDepth(1, args...)
}

func (t *tagLog) FatalDepth(level int, args ...interface{}) {
	args = append([]interface{}{t.Tag}, args...)
	glog.FatalDepth(level+1, args...)
}

func (t *tagLog) FatalStack(skip, depth int, args ...interface{}) {
	var pc = make([]uintptr, depth)
	runtime.Callers(skip+1, pc)
	var msgs = make([]string, 0)
	var s string
	for _, p := range pc {
		if p != 0 {
			file, line := runtime.FuncForPC(p).FileLine(p)
			t.splitFilePath(&file)
			s = fmt.Sprintf("%s:%d:%s", file, line, runtime.FuncForPC(p).Name())
			msgs = append(msgs, s)
		} else {
			break
		}
	}
	s = ""
	for i := len(msgs) - 1; i >= 0; i-- {
		if i == 0 {
			s += msgs[i]
		}
		s += msgs[i] + "  ->  "
	}
	t.Fatalln(s)
}

func (t *tagLog) FatalFullStack(args ...interface{}) {
	var msg = string(debug.Stack())
	args = append([]interface{}{msg}, args)
	t.Fatalln(args...)
}

/**************************/
func (t *tagLog) Off() IMLog {
	return &noneLog{}
}
