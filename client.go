package apierrors

import (
	"github.com/google/google-api-go-client/googleapi"
)

// Is the given error a google api error with given reason?
func IsApiError(err error, reason Reason) bool {
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
