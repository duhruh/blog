package errors

import (
	"errors"

	stack "github.com/go-errors/errors"
)

var (
	ErrorCouldNotFindBlogById  = errors.New("cannot find blog by id")
	ErrorCouldNotRetrieveBlogs = errors.New("could not retrieve all blogs")
)

func New(err error) error {
	return stack.New(err)
}

func StackTrace(err error) string {
	if err == nil {
		return "null"
	}
	stackErr := err.(*stack.Error)
	return stackErr.ErrorStack()
}
