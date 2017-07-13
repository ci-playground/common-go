package errors

import (
	errs "github.com/pkg/errors"
)

type stackTracer interface {
	StackTrace() errs.StackTrace
}

// Stack returns a stack trace if a given error has it, or nil otherwise.
func Stack(err error) errs.StackTrace {
	tracer, _ := err.(stackTracer)
	if tracer == nil {
		return nil
	}
	return tracer.StackTrace()
}
