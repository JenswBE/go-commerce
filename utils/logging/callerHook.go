package logging

import (
	"fmt"
	"runtime"

	"github.com/rs/zerolog"
)

type CallerInfoHook struct{}

// Based on https://github.com/rs/zerolog/issues/22#issuecomment-1127295489
func (h CallerInfoHook) Run(e *zerolog.Event, l zerolog.Level, msg string) {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		e.Str("caller", fmt.Sprintf("%s:%d", file, line))
	}
}
