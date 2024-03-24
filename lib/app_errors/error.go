package apperrors

import (
	"errors"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
)

type Error struct {
	Err   error  `json:"error,omitempty"`
	Stack string `json:"stack,omitempty"`
	Args  []any  `json:"args,omitempty"`
}

func New(err error, args ...any) error {
	return &Error{
		Err:   err,
		Stack: stackTrace(),
		Args:  args,
	}
}

func (e *Error) Error() string {
	return e.Err.Error()
}

func Wrap(err error, args ...any) error {
	if err == nil {
		return nil
	}
	return New(err, args...)
}

func stackTrace() string {
	const skipCount = 4
	const depth = 16

	pcs := make([]uintptr, depth)

	frameCount := runtime.Callers(skipCount, pcs)
	frames := runtime.CallersFrames(pcs[:frameCount])

	var sb strings.Builder
	sb.WriteString("goroutine [running]:\n")

	for {
		frame, more := frames.Next()

		sb.WriteString(frame.Function)
		sb.WriteString("\n\t")
		sb.WriteString(frame.File)
		sb.WriteString(":")
		sb.WriteString(strconv.Itoa(frame.Line))
		sb.WriteString("\n")

		if !more {
			break
		}
	}

	return sb.String()
}

func IsError(err error, target error) bool {
	converted := &Error{}
	if ok := errors.As(err, &converted); ok {
		return errors.Is(converted.Err, target)
	}
	return errors.Is(err, target)
}

type MultipleErrorBuilder struct {
	Errors []error `json:"errors,omitempty"`
}

func NewMultipleErrorBuilder() *MultipleErrorBuilder {
	return &MultipleErrorBuilder{
		Errors: []error{},
	}
}
func (e *MultipleErrorBuilder) Append(err error) *MultipleErrorBuilder {
	e.Errors = append(e.Errors, err)
	return e
}
func (e *MultipleErrorBuilder) Build() error {
	if len(e.Errors) == 0 {
		return nil
	}
	if len(e.Errors) == 1 {
		return e.Errors[0]
	}

	return Wrap(errors.Join(e.Errors...), e.Errors)
}

func Print(err error) {
	converted := &Error{}
	if ok := errors.As(err, &converted); ok {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		fmt.Fprintln(os.Stderr, converted.Args)
		fmt.Fprint(os.Stderr, converted.Stack)
	} else {
		Print(Wrap(err))
	}
}

func Source(err error) error {
	converted := &Error{}
	if ok := errors.As(err, &converted); ok {
		return converted.Err
	}
	return err
}

func Handle(err error) {
	if err != nil {
		Print(err)
		os.Exit(1)
	}
}
