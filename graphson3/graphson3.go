package graphson3

import (
	"errors"
	"math"

	"github.com/dnoberon/graphson"

	"github.com/buger/jsonparser"
)

type GraphSONv3Parser struct{}

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

// getValueType examines a GraphSON 3 value/type pair and returns the correct value
func getValueType(in []byte) (graphson.ValueType, error) {
	typeName, err := jsonparser.GetString(in, "@type")
	if err != nil {
		if len(in) != 0 {
			raw := string(in)
			if raw == "true" || raw == "false" {
				return graphson.Boolean, nil
			}

			return graphson.String, nil
		}
	}

	return valueTypeFromString(typeName), nil
}

func valueTypeFromString(raw string) graphson.ValueType {
	switch raw {
	case "g:Class":
		return graphson.Class
	case "g:Date":
		return graphson.Date
	case "g:Double":
		return graphson.Double
	case "g:Float":
		return graphson.Float
	case "g:Int32":
		return graphson.Int32
	case "g:Int64":
		return graphson.Int64
	case "g:List":
		return graphson.List
	case "g:Map":
		return graphson.Map
	case "g:Timestamp":
		return graphson.Timestamp
	case "g:Set":
		return graphson.Set
	case "g:UUID":
		return graphson.UUID
	case "g:Vertex":
		return graphson.Vertex
	default:
		return graphson.Unknown
	}
}
