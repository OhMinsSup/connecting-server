package lib

import (
	"encoding/json"
	"fmt"
	"os"
)

// defaultLog manually encodes the log to STDERR, providing a basic, default logging implementation
// before log is fully configured.
func defaultLog(level, msg string, err error) {
	log := struct {
		Level   string  `json:"level"`
		Message string  `json:"msg"`
		Error   error   `json:"error"`
	}{
		level,
		msg,
		err,
	}

	if b, err := json.Marshal(log); err != nil {
		fmt.Fprintf(os.Stderr, `{"level":"error","msg":"failed to encode log message","error":"error message"}%s`, "\n")
	} else {
		fmt.Fprintf(os.Stderr, "%s\n", b)
	}
}

func DefaultDebugLog(msg string, err error) {
	defaultLog("debug", msg, err)
}

func DefaultInfoLog(msg string, err error) {
	defaultLog("info", msg, err)
}

func DefaultWarnLog(msg string, err error) {
	defaultLog("warn", msg, err)
}

func DefaultErrorLog(msg string, err error) {
	defaultLog("error", msg, err)
}

func DefaultCriticalLog(msg string, err error) {
	// We map critical to error in zap, so be consistent.
	defaultLog("error", msg, err)
}
