package main

import (
	"fmt"
)

type httpError struct {
	Err     error
	Message string
	Code    int
}

func (e *httpError) Error() string {
	return fmt.Sprintf("HttpError[%s] Message[%s] Code[%d]", e.Err, e.Message, e.Code)
}
