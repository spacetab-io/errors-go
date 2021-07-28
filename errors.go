package errors

import (
	"fmt"
)

// HasErrors checks if error occurs in passed err.
func HasErrors(err interface{}) (hasErrors bool) {
	switch e := err.(type) {
	case []error:
		hasErrors = len(e) > 0
	case map[string]error:
		hasErrors = len(e) > 0
	default:
		hasErrors = e != nil
	}

	return
}

func Wrap(method, action string, err error) error {
	return fmt.Errorf("%s %s error: %w", method, action, err)
}
