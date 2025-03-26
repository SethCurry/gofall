package gofall

import (
	"errors"
	"fmt"
	"strings"
)

// APIError is the response from the Scryfall API when an error occurs.
// Most API methods in this package will return an APIError whenever possible.
type APIError struct {
	// The HTTP status code of the response
	Status int `json:"status"`

	// A machine-friendly code for the error condition
	Code string `json:"code"`

	Details string `json:"details"`

	// A computer-friendly string specifying more details about the error condition.
	// E.g. for a 404 it might return "ambiguous" if the request refers to multiple
	// cards potentially.
	//
	// This field can be empty.
	Type string `json:"type"`

	// A series of human-readable errors.
	Warnings []string `json:"warnings"`
}

func (a APIError) Error() string {
	return fmt.Sprintf("API Error: %s: %s", a.Code, strings.Join(a.Warnings, " | "))
}

// ErrNoBackFace is returned when a request asks for a card back,
// but the card has no back face on Scryfall.
var ErrNoBackFace = errors.New("no back face")

func matchAPIError(err *APIError) error {
	if err.Status == 422 {
		return ErrNoBackFace
	}

	return nil
}
