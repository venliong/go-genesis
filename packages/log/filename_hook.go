package log

import (
	"path"
	"runtime"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

// ContextHook storing nothing but behavior
type ContextHook struct{}

// Levels returns all log levels
func (hook ContextHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

// Fire the log entry
func (hook ContextHook) Fire(entry *logrus.Entry) error {
	pc := make([]uintptr, 3, 3)
	cnt := runtime.Callers(6, pc)

	for i := 0; i < cnt; i++ {
		fu := runtime.FuncForPC(pc[i] - 1)
		name := fu.Name()
		if !strings.Contains(name, "github.com/sirupsen/logrus") {
			file, line := fu.FileLine(pc[i] - 1)
			entry.Data["file"] = path.Base(file)
			entry.Data["func"] = path.Base(name)
			entry.Data["line"] = line
			entry.Data["time"] = time.Now().Format(time.RFC3339)
			break
		}
	}
	return nil
}
