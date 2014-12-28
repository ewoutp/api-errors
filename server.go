package apierrors

import (
	"encoding/json"
	"net/http"

	"github.com/google/google-api-go-client/googleapi"
)

// Send an error response conform the googleapi error format.
func WriteApiError(resp http.ResponseWriter, reason Reason, message string, httpStatusCode int) error {
	err := map[string]interface{}{
		"error": googleapi.Error{
			Message: message,
			Errors: []googleapi.ErrorItem{
				googleapi.ErrorItem{
					Reason:  string(reason),
					Message: message,
				},
			},
		},
	}
	resp.Header().Add("Content-Type", "application/json")
	resp.WriteHeader(httpStatusCode)
	return json.NewEncoder(resp).Encode(err)
}
