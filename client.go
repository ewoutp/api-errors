package apierrors

import (
	"google.golang.org/api/googleapi"
)

// Is the given error a google api error?
func IsApiError(err error) bool {
	_, ok := err.(*googleapi.Error)
	return ok
}

// Is the given error a google api error with given reason?
func IsApiErrorWithReason(err error, reason Reason) bool {
	aerr, ok := err.(*googleapi.Error)
	if !ok {
		return false
	}
	for _, e := range aerr.Errors {
		if e.Reason == string(reason) {
			return true
		}
	}
	return false
}
