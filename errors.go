package graphson

import (
	"fmt"
	"strings"
)

// ParsingError is a self-contained way of handling internal errors. Best used with the .Unwrap() functionality in the
// error package from 1.13
type ParsingError struct {
	Message   interface{}
	Operation string
	Field     string
}

// Error wraps and satisfies the errors packagee
func (e ParsingError) Error() string {
	return fmt.Sprintf("%v: field: %s operation: %s", e.Message, e.Field, e.Operation)
}

// ParsingErrors is a an ease of use struct
type ParsingErrors []ParsingError

// Combine all current parsing errors in to a single, friendly standard error
func (pe ParsingErrors) Combine() error {
	errorMessages := []interface{}{}
	operations := []string{}
	fields := []string{}

	for _, err := range pe {
		errorMessages = append(errorMessages, err.Message)
		operations = append(operations, err.Operation)
		fields = append(fields, err.Field)
	}

	if len(errorMessages) == 0 && len(operations) == 0 && len(fields) == 0 {
		return nil
	}

	return ParsingError{fmt.Sprintf("multiple parsing errors: [%v]", errorMessages),
		strings.Join(operations, ","),
		strings.Join(fields, ",")}
}
