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

// getValueType examines a GraphSON 3 value/type pair and returns the correct value
func getValueType(in []byte) (valueType, error) {
	typeName, err := jsonparser.GetString(in, "@type")
	if err != nil {
		if len(in) != 0 {
			raw := string(in)
			if raw == "true" || raw == "false" {
				return Boolean, nil
			}

			return String, nil
		}
	}

	return valueTypeFromString(typeName), nil
}

type valueType int

const (
	String = valueType(iota)
	Boolean
	Class
	Date
	Double
	Float
	Int64
	Int32
	List
	Map
	Timestamp
	Set
	UUID
	Vertex
	VertexProperty
	Edge
	EdgeProperty
	Unknown
)

func (vt valueType) string() string {
	switch vt {
	case String:
		return "String"
	case Boolean:
		return "Boolean"
	case Class:
		return "g:Class"
	case Date:
		return "g:Date"
	case Double:
		return "g:Double"
	case Float:
		return "g:Float"
	case Int32:
		return "g:Int32"
	case Int64:
		return "g:Int64"
	case List:
		return "g:List"
	case Map:
		return "g:Map"
	case Timestamp:
		return "g:Timestamp"
	case Set:
		return "g:Set"
	case UUID:
		return "g:UUID"
	default:
		return "unknown"
	}
}

func valueTypeFromString(raw string) valueType {
	switch raw {
	case "g:Class":
		return Class
	case "g:Date":
		return Date
	case "g:Double":
		return Double
	case "g:Float":
		return Float
	case "g:Int32":
		return Int32
	case "g:Int64":
		return Int64
	case "g:List":
		return List
	case "g:Map":
		return Map
	case "g:Timestamp":
		return Timestamp
	case "g:Set":
		return Set
	case "g:UUID":
		return UUID
	case "g:Vertex":
		return Vertex
	default:
		return Unknown
	}
}
