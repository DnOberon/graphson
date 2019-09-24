package graphson3

import (
	"time"

	"github.com/buger/jsonparser"
	"github.com/dnoberon/graphson"
)

// Parse accepts a valid @type/@value pair and returns the parsed object. Additional operations can be used to discover type
func (g GraphSONv3Parser) Parse(in []byte) (graphson.ValuePair, error) {
	typeName, err := getValueType(in)
	if err != nil {
		return graphson.ValuePair{}, err
	}

	var out interface{}

	switch typeName {
	case graphson.Vertex:
		out, err = g.ParseVertex(in)
	case graphson.VertexProperty:
		out, err = g.ParseVertexProperty(in)
	case graphson.Edge:
		out, err = g.ParseEdge(in)
	case graphson.EdgeProperty:
		out, err = g.ParseProperty(in)
	case graphson.Set:
		out, err = g.parseSet(in)
	case graphson.List:
		out, err = g.parseSet(in)
	case graphson.Class:
		out, err = g.parseClass(in)
	case graphson.String:
		out, err = string(in), nil
	case graphson.Boolean:
		out, err = string(in) == "true" || string(in) == "1", nil
	case graphson.Int32:
		out, err = g.parseInt32(in)
	case graphson.Int64:
		out, err = g.parseInt64(in)
	case graphson.Float:
		out, err = g.parseFloat64(in)
	case graphson.Double:
		out, err = g.parseFloat32(in)
	case graphson.UUID:
		out, err = g.parseUUID(in)
	case graphson.Date:
		out, err = g.parseTimestamp(in)
	case graphson.Timestamp:
		out, err = g.parseTimestamp(in)
	}

	return graphson.ValuePair{Type: typeName, Value: out}, err
}

// parseSet also applies to the g:List type, Formatting between the two types is exactly the same
func (g GraphSONv3Parser) parseSet(in []byte) ([]graphson.ValuePair, error) {
	var out []graphson.ValuePair

	vt, err := getValueType(in)
	if err != nil {
		return nil, err
	}

	if vt != graphson.Set && vt != graphson.List {
		return nil, graphson.ParsingError{Message: "provided input not a g:Set or g:List type", Operation: "parseSet", Field: "@type"}
	}

	value, dt, _, err := jsonparser.Get(in, "@value")
	if dt != jsonparser.Array {
		return nil, graphson.ParsingError{Message: "provided input not a valid g:Set or g:List type, bad array", Operation: "parseSet", Field: "@type"}
	}

	parsingErrors := graphson.ParsingErrors{}
	_, err = jsonparser.ArrayEach(value, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		currentError := graphson.ParsingError{Operation: "parseSet", Field: "@value"}

		if err != nil {
			currentError.Message = err.Error()
			parsingErrors = append(parsingErrors, currentError)
			return
		}

		vp, err := g.Parse(value)
		if err != nil {
			currentError.Message = err.Error()
			parsingErrors = append(parsingErrors, currentError)
			return
		}

		out = append(out, vp)
	})

	return out, parsingErrors.Combine()
}

func (g GraphSONv3Parser) parseFlatMap(in []byte) ([]graphson.ValuePair, error) {
	ordered := []graphson.ValuePair{}

	vt, err := getValueType(in)
	if err != nil {
		return nil, err
	}

	if vt != graphson.Map {
		return nil, graphson.ParsingError{Message: "provided input not a g:Map type", Operation: "parseMap", Field: "@type"}
	}

	value, dt, _, err := jsonparser.Get(in, "@value")
	if dt != jsonparser.Array {
		return nil, graphson.ParsingError{Message: "provided input not a valid g:Map type, bad array", Operation: "parseMap", Field: "@type"}
	}

	parsingErrors := graphson.ParsingErrors{}
	_, err = jsonparser.ArrayEach(value, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		currentError := graphson.ParsingError{Operation: "parseMap", Field: "@value"}

		if err != nil {
			currentError.Message = err.Error()
			parsingErrors = append(parsingErrors, currentError)
			return
		}

		vp, err := g.Parse(value)
		if err != nil {
			currentError.Message = err.Error()
			parsingErrors = append(parsingErrors, currentError)
			return
		}

		ordered = append(ordered, vp)

	})

	// because g:Map relies so heavily on element order it is suggested that results are always ignored if any error is present
	return ordered, nil
}

func (g GraphSONv3Parser) parseInt32(in []byte) (int, error) {
	vt, err := getValueType(in)
	if err != nil {
		return 0, err
	}

	if vt != graphson.Int32 {
		return 0, graphson.ParsingError{Message: "provided input not a g:Int32 type", Operation: "parseInt32", Field: "@type"}
	}

	value, err := jsonparser.GetInt(in, "@value")

	// There is a very real possibility that we unintentionally truncate the value if it is not really an int32
	return int(value), err
}

func (g GraphSONv3Parser) parseInt64(in []byte) (int64, error) {
	vt, err := getValueType(in)
	if err != nil {
		return 0, err
	}

	if vt != graphson.Int64 {
		return 0, graphson.ParsingError{Message: "provided input not a g:Int64 type", Operation: "parseInt64", Field: "@type"}
	}

	return jsonparser.GetInt(in, "@value")
}

func (g GraphSONv3Parser) parseFloat32(in []byte) (float32, error) {
	vt, err := getValueType(in)
	if err != nil {
		return 0, err
	}

	if vt != graphson.Double {
		return 0, graphson.ParsingError{Message: "provided input not a g:Float32 type", Operation: "parseFloat32", Field: "@type"}
	}

	value, err := jsonparser.GetFloat(in, "@value")

	return float32(value), err
}

func (g GraphSONv3Parser) parseFloat64(in []byte) (float64, error) {
	vt, err := getValueType(in)
	if err != nil {
		return 0, err
	}

	if vt != graphson.Float {
		return 0, graphson.ParsingError{Message: "provided input not a g:Float64 type", Operation: "parseFloat64", Field: "@type"}
	}

	return jsonparser.GetFloat(in, "@value")
}

func (g GraphSONv3Parser) parseTimestamp(in []byte) (time.Time, error) {
	vt, err := getValueType(in)
	if err != nil {
		return time.Time{}, err
	}

	if vt != graphson.Timestamp && vt != graphson.Date {
		return time.Time{}, graphson.ParsingError{Message: "provided input not a g:Timestamp g:Date type", Operation: "parseTimestamp", Field: "@type"}
	}

	value, err := jsonparser.GetInt(in, "@value")
	if err != nil {
		return time.Time{}, graphson.ParsingError{Message: err.Error(), Operation: "parseTimestamp", Field: "@value"}
	}

	return time.Unix(value/10000, 0), nil
}

func (g GraphSONv3Parser) parseClass(in []byte) (string, error) {
	vt, err := getValueType(in)
	if err != nil {
		return "", err
	}

	if vt != graphson.Class {
		return "", graphson.ParsingError{Message: "provided input not g:Class type", Operation: "parseClass", Field: "@type"}
	}

	return jsonparser.GetString(in, "@value")
}

func (g GraphSONv3Parser) parseUUID(in []byte) (string, error) {
	vt, err := getValueType(in)
	if err != nil {
		return "", err
	}

	if vt != graphson.UUID {
		return "", graphson.ParsingError{Message: "provided input not g:UUID type", Operation: "parseClass", Field: "@type"}
	}

	return jsonparser.GetString(in, "@value")
}
