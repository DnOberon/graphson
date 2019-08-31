package graphson

import (
	"errors"
	"fmt"
	"math"
	"strings"

	"github.com/buger/jsonparser"
)

type ParsingError struct {
	Message   interface{}
	Operation string
	Field     string
}

func (e ParsingError) Error() string {
	return fmt.Sprintf("%v: field: %s operation: %s", e.Message, e.Field, e.Operation)
}

type ParsingErrors []ParsingError

func (pe ParsingErrors) combine() error {
	errorMessages := []interface{}{}
	operations := []string{}
	fields := []string{}

	if len(errorMessages) == 0 && len(operations) == 0 && len(fields) == 0 {
		return nil
	}

	for _, err := range pe {
		errorMessages = append(errorMessages, err.Message)
		operations = append(operations, err.Operation)
		fields = append(fields, err.Field)
	}

	return ParsingError{fmt.Sprintf("multiple parsing errors: [%v]", errorMessages),
		strings.Join(operations, ","),
		strings.Join(fields, ",")}
}

func parsedToType(in []byte, vt jsonparser.ValueType) (interface{}, error) {

	switch vt {
	case jsonparser.String:
		return jsonparser.ParseString(in)

	case jsonparser.Number:
		n, err := jsonparser.ParseFloat(in)
		if err != nil {
			return nil, err
		}

		if math.Trunc(n) == n {
			return int64(n), nil
		}

		return n, nil
	case jsonparser.Boolean:
		return jsonparser.ParseBoolean(in)

	case jsonparser.Unknown:
		return in, errors.New("unknown type or invalid data")

	case jsonparser.Null:
		return in, errors.New("unknown type or invalid data")
	}

	return in, nil
}
