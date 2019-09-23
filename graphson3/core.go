package graphson3

import (
	"time"

	"github.com/buger/jsonparser"
	"github.com/dnoberon/graphson"
)

// Parse accepts a valid @type/@value pair and returns the parsed object. Additional operations can be used to discover type
func Parse(in []byte) (interface{}, valueType, error) {
	typeName, err := getValueType(in)
	if err != nil {
		return nil, 0, err
	}

	var out interface{}

	switch typeName {
	case Vertex:
		out, err = ParseVertex(in)
	case VertexProperty:
		out, err = ParseVertexProperty(in)
	case Edge:
		out, err = ParseEdge(in)
	case EdgeProperty:
		out, err = ParseProperty(in)
	case Set:
		out, err = parseSet(in)
	case List:
		out, err = parseSet(in)
	case Class:
		out, err = parseClass(in)
	case String:
		out, err = string(in), nil
	case Boolean:
		out, err = string(in) == "true" || string(in) == "1", nil
	case Int32:
		out, err = parseInt32(in)
	case Int64:
		out, err = parseInt64(in)
	case Float:
		out, err = parseFloat64(in)
	case Double:
		out, err = parseFloat32(in)
	case UUID:
		out, err = parseUUID(in)
	case Date:
		out, err = parseTimestamp(in)
	case Timestamp:
		out, err = parseTimestamp(in)
	}

	return out, typeName, err
}

// parseSet also applies to the g:List type, Formatting between the two types is exactly the same
func parseSet(in []byte) ([]interface{}, error) {
	var out []interface{}

	vt, err := getValueType(in)
	if err != nil {
		return nil, err
	}

	if vt != Set {
		return nil, graphson.ParsingError{Message: "provided input not a g:Set type", Operation: "parseSet", Field: "@type"}
	}

	value, dt, _, err := jsonparser.Get(in, "@value")
	if dt != jsonparser.Array {
		return nil, graphson.ParsingError{Message: "provided input not a valid g:Set type, bad array", Operation: "parseSet", Field: "@type"}
	}

	parsingErrors := graphson.ParsingErrors{}
	_, err = jsonparser.ArrayEach(value, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		currentError := graphson.ParsingError{Operation: "parseSet", Field: "@value"}

		if err != nil {
			currentError.Message = err.Error()
			parsingErrors = append(parsingErrors, currentError)
			return
		}

		v, _, err := Parse(value)
		if err != nil {
			currentError.Message = err.Error()
			parsingErrors = append(parsingErrors, currentError)
			return
		}

		out = append(out, v)
	})

	return out, parsingErrors.Combine()
}

func parseMap(in []byte) (map[interface{}]interface{}, error) {
	out := map[interface{}]interface{}{}

	vt, err := getValueType(in)
	if err != nil {
		return nil, err
	}

	if vt != Map {
		return nil, graphson.ParsingError{Message: "provided input not a g:Map type", Operation: "parseMap", Field: "@type"}
	}

	return out, nil
}

func parseInt32(in []byte) (int, error) {
	vt, err := getValueType(in)
	if err != nil {
		return 0, err
	}

	if vt != Int32 {
		return 0, graphson.ParsingError{Message: "provided input not a g:Int32 type", Operation: "parseInt32", Field: "@type"}
	}

	value, err := jsonparser.GetInt(in, "@value")

	// There is a very real possibility that we unintentionally truncate the value if it is not really an int32
	return int(value), err
}

func parseInt64(in []byte) (int64, error) {
	vt, err := getValueType(in)
	if err != nil {
		return 0, err
	}

	if vt != Int64 {
		return 0, graphson.ParsingError{Message: "provided input not a g:Int64 type", Operation: "parseInt64", Field: "@type"}
	}

	return jsonparser.GetInt(in, "@value")
}

func parseFloat32(in []byte) (float32, error) {
	vt, err := getValueType(in)
	if err != nil {
		return 0, err
	}

	if vt != Double {
		return 0, graphson.ParsingError{Message: "provided input not a g:Float32 type", Operation: "parseFloat32", Field: "@type"}
	}

	value, err := jsonparser.GetFloat(in, "@value")

	return float32(value), err
}

func parseFloat64(in []byte) (float64, error) {
	vt, err := getValueType(in)
	if err != nil {
		return 0, err
	}

	if vt != Float {
		return 0, graphson.ParsingError{Message: "provided input not a g:Float64 type", Operation: "parseFloat64", Field: "@type"}
	}

	return jsonparser.GetFloat(in, "@value")
}

func parseTimestamp(in []byte) (time.Time, error) {
	vt, err := getValueType(in)
	if err != nil {
		return time.Time{}, err
	}

	if vt != Timestamp && vt != Date {
		return time.Time{}, graphson.ParsingError{Message: "provided input not a g:Timestamp g:Date type", Operation: "parseTimestamp", Field: "@type"}
	}

	value, err := jsonparser.GetInt(in, "@value")
	if err != nil {
		return time.Time{}, graphson.ParsingError{Message: err.Error(), Operation: "parseTimestamp", Field: "@value"}
	}

	return time.Unix(value/10000, 0), nil
}

func parseClass(in []byte) (string, error) {
	vt, err := getValueType(in)
	if err != nil {
		return "", err
	}

	if vt != Class {
		return "", graphson.ParsingError{Message: "provided input not g:Class type", Operation: "parseClass", Field: "@type"}
	}

	return jsonparser.GetString(in, "@value")
}

func parseUUID(in []byte) (string, error) {
	vt, err := getValueType(in)
	if err != nil {
		return "", err
	}

	if vt != UUID {
		return "", graphson.ParsingError{Message: "provided input not g:UUID type", Operation: "parseClass", Field: "@type"}
	}

	return jsonparser.GetString(in, "@value")
}
