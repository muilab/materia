package utils

import "time"

// DateTimeUTC returns the current UTC time in RFC3339 format.
func DateTimeUTC() string {
	return time.Now().UTC().Format(time.RFC3339)
}

// PanicOnError will panic if the error is not nil.
func PanicOnError(err error) {
	if err != nil {
		panic(err)
	}
}
