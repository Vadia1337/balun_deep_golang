package main

import (
	bytes2 "bytes"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

type MultiError struct {
	errors []error
}

func (e *MultiError) Error() string {
	if len(e.errors) == 0 {
		return ""
	}

	var bytes bytes2.Buffer
	bytes.WriteString(fmt.Sprintf("%d errors occured:\n", len(e.errors)))

	for i := range e.errors {
		bytes.WriteString(fmt.Sprintf("\t* %s", e.errors[i].Error()))
	}

	bytes.WriteString("\n")

	return bytes.String()
}

func Append(err error, errs ...error) *MultiError {
	if err == nil && len(errs) == 0 {
		return nil
	}

	var multiErr *MultiError
	if errors.As(err, &multiErr) {
		multiErr.errors = append(multiErr.errors, errs...)
		return multiErr
	}

	erros := make([]error, 0, len(errs)+1)
	if err != nil {
		erros = append(erros, err)
	}

	erros = append(erros, errs...)

	return &MultiError{
		errors: erros,
	}
}

func TestMultiError(t *testing.T) {
	var err error
	err = Append(err, errors.New("error 1"))
	err = Append(err, errors.New("error 2"))

	expectedMessage := "2 errors occured:\n\t* error 1\t* error 2\n"
	assert.EqualError(t, err, expectedMessage)
}
