package error

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
)

type Error struct {
	ctx      context.Context
	code     Code
	messages []string
	err      error
}

func New(ctx context.Context, code Code, message string, err error) error {
	if _, ok := err.(*Error); ok {
		return err
	}

	if _, ok := mappingCodeValue[code]; !ok {
		code = CodeUnknown
	}

	newErr := &Error{
		ctx:      context.Background(),
		code:     code,
		messages: []string{message},
		err:      errors.New(mappingCodeValue[code]),
	}

	if ctx != nil {
		newErr.ctx = ctx
	}
	if err != nil {
		newErr.err = err
	}

	log.Println(fmt.Sprintf("[!%v] %s: %s (%s)", newErr.code, newErr.GetTitle(), newErr.messages[0], newErr.err.Error()))

	return newErr
}

func Cast(err error) *Error {
	if err == nil {
		return nil
	}

	if c, ok := err.(*Error); ok {
		return c
	}

	return New(context.Background(), CodeUnknown, "", err).(*Error)
}

func (e *Error) AppendMessage(message string) {
	e.messages = append(e.messages, message)
}

func (e *Error) Error() string {
	if e.err == nil {
		if len(e.messages) > 0 {
			return e.messages[0]
		}
		return "unknown error"
	}
	return e.err.Error()
}

func (e *Error) GetCode() int {
	return int(e.code)
}

func (e *Error) GetHTTPCodeEquivalent() int {
	if e.code < 1000 {
		return int(e.code)
	}
	return http.StatusInternalServerError
}

func (e *Error) GetTitle() string {
	return mappingCodeValue[e.code]
}

func (e *Error) GetError() error {
	return e.err
}

func (e *Error) IsInternalServerError() bool {
	return e.code >= minCustomErrorCode
}

func (e *Error) GetMessages() []string {
	return e.messages
}
