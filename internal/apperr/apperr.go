package apperr

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
)

type Trace struct {
	File string
	Line int
	Func string
}

type AppError struct {
	Err    error
	Traces []Trace
}

func (e *AppError) Error() string {
	return e.Err.Error()
}

func Wrap(err error) *AppError {
	if err == nil {
		return nil
	}

	var appErr *AppError
	if AsAppError(err, &appErr) {
		appErr.Traces = append(appErr.Traces, getTrace())
		return appErr
	}

	return &AppError{
		Err:    err,
		Traces: []Trace{getTrace()},
	}
}

func (e *AppError) DetailMessage() string {
	var sBuilder strings.Builder
	sBuilder.WriteString(e.Error())

	if len(e.Traces) > 0 {
		sBuilder.WriteString("\n  └ trace:\n")
		for _, t := range e.Traces {
			sBuilder.WriteString("      ")
			sBuilder.WriteString(fmt.Sprintf("%s:%d (%s)\n", t.File, t.Line, t.Func))
		}
	}
	return sBuilder.String()
}

func getTrace() Trace {
	pc, file, line, ok := runtime.Caller(2)
	fn := "unknown"
	if ok {
		fn = runtime.FuncForPC(pc).Name()
	}
	return Trace{
		File: filepath.Base(file),
		Line: line,
		Func: fn,
	}
}

func AsAppError(err error, target **AppError) bool {
	appErr, ok := err.(*AppError)
	if !ok {
		return false
	}
	*target = appErr
	return true
}
