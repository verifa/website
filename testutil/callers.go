package testutil

import (
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

// Callers prints the stack trace of everything up til the line where Callers()
// was invoked.
func Callers() string {
	var pc [50]uintptr
	n := runtime.Callers(
		2,
		pc[:],
	) //nolint:gomnd    // skip runtime.Callers + Callers
	callsites := make([]string, 0, n)
	frames := runtime.CallersFrames(pc[:n])

	for frame, more := frames.Next(); more; frame, more = frames.Next() {
		callsites = append(callsites, frame.File+":"+strconv.Itoa(frame.Line))
	}

	callsites = callsites[:len(callsites)-1] // skip testing.tRunner
	if len(callsites) == 1 {
		return ""
	}

	var b strings.Builder

	for i := len(callsites) - 1; i >= 0; i-- {
		if b.Len() > 0 {
			b.WriteString(" -> ")
		}

		b.WriteString(filepath.Base(callsites[i]))
	}

	return "\n" + b.String() + ":"
}
