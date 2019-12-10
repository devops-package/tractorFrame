package Guzzle

import (
	"context"
	"net/http"
	"time"
)

type Response struct {

	Duration time.Duration

	// StatusCode is the HTTP Status Code returned by the HTTP Response. Taken from resp.StatusCode.
	StatusCode int

	// Header stores the response headers as http.Header interface.
	Header http.Header

	// Cookies stores the parsed response cookies.
	Cookies []*http.Cookie

	// Expose the native Go http.Response object for convenience.
	RawResponse *http.Response

	// Expose the native Go http.Request object for convenience.
	RawRequest *http.Request

	// Expose original request Context for convenience.
	Context *context.Context

}
