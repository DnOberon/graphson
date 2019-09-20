package graphson3

import (
	"github.com/buger/jsonparser"
	"github.com/dnoberon/graphson"
)

// TODO: some kind of recursive parsing func for handling nested items in a set/map

// Parse accepts a valid @type/@value pair and returns the parsed object. Additional operations can be used to discover type
func Parse(in []byte) (interface{}, error) {
	typeName, err := getValueType(in)
	if err != nil {
		return nil, err
	}

	switch typeName {
	case Vertex:
		return ParseVertex(in)
	case VertexProperty:
		return ParseVertexProperty(in)
	case Edge:
		return ParseEdge(in)
	case EdgeProperty:
		return ParseProperty(in)
	case Set:
		return parseSet(in)
	case String:
		return string(in), nil
	case Boolean:
		return string(in) == "true" || string(in) == "1", nil
	case Int32:
		return parseInt32(in)
	case Int64:
		return parseInt64(in)
	case Float:
		return parseFloat64(in)
	case Double:
		return parseFloat32(in)
	}

	return nil, nil
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

		v, err := Parse(value)
		if err != nil {
			currentError.Message = err.Error()
			parsingErrors = append(parsingErrors, currentError)
			return
		}

		out = append(out, v)
	})

	return out, parsingErrors.Combine()
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
