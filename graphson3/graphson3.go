package graphson3

import (
	"errors"
	"math"

	"github.com/buger/jsonparser"
)

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
