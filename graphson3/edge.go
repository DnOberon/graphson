package graphson3

import (
	"strings"

	"github.com/dnoberon/graphson"

	"github.com/buger/jsonparser"
)

const edgeTypename = "g:Edge"
const propertyTypeName = "g:Property"

type Edge struct {
	ID         interface{}         `json:"id"` // ID as interface{}, different providers use different ID types
	Label      string              `json:"label"`
	InVLabel   string              `json:"inVLabel"`
	OutVLabel  string              `json:"outVLabel"`
	InV        interface{}         `json:"inV"`
	OutV       interface{}         `json:"outV"`
	Properties map[string]Property `json:"properties"`
}

type Property struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}

// ParseEdge expects the input to be valid JSON and to be a single Edge record. See either the testing file for sample
// edge json records or http://tinkerpop.apache.org/docs/3.4.2/dev/io/#_edge_3.
func ParseEdge(in []byte) (e Edge, err error) {
	e.Properties = map[string]Property{}

	if typename, err := jsonparser.GetString(in, "@type"); err != nil || typename != edgeTypename {
		return e, graphson.ParsingError{err, "@type", "parseEdge"}
	}

	// value location mapping on original json record, using the jsonparser package to avoid as much reflection as we can
	var paths = [][]string{
		{"@value", "id", "@value"},
		{"@value", "label"},
		{"@value", "inVLabel"},
		{"@value", "outVLabel"},
		{"@value", "inV", "@value"},
		{"@value", "outV", "@value"},
		{"@value", "properties"},
	}

	parsingErrors := graphson.ParsingErrors{}

	jsonparser.EachKey(in, func(idx int, value []byte, vt jsonparser.ValueType, err error) {
		currentError := graphson.ParsingError{nil, strings.Join(paths[idx], " "), "parseEdge"}

		if err != nil {
			currentError.Message = err.Error()
			parsingErrors = append(parsingErrors, currentError)
			return
		}

		switch idx {
		case 0: // @value -> id -> @value
			id, err := parsedToType(value, vt)
			if err != nil {
				currentError.Message = err.Error()
				break
			}

			e.ID = id

		case 1: // @value -> label
			label, err := jsonparser.ParseString(value)
			if err != nil {
				currentError.Message = err.Error()
				break
			}

			e.Label = label

		case 2: // @value -> inVLabel
			label, err := jsonparser.ParseString(value)
			if err != nil {
				currentError.Message = err.Error()
				break
			}

			e.InVLabel = label

		case 3: // @value -> outVLabel
			label, err := jsonparser.ParseString(value)
			if err != nil {
				currentError.Message = err.Error()
				break
			}

			e.OutVLabel = label

		case 4: // @value -> inV -> @value
			v, err := parsedToType(value, vt)
			if err != nil {
				currentError.Message = err.Error()
				break
			}

			e.InV = v

		case 5: // @value -> outV -> @value
			v, err := parsedToType(value, vt)
			if err != nil {
				currentError.Message = err.Error()
				break
			}

			e.OutV = v

		case 6: // @value -> properties
			err = jsonparser.ObjectEach(value, func(key []byte, prop []byte, dataType jsonparser.ValueType, offset int) error {
				propertyName, err := jsonparser.ParseString(key)
				if err != nil {
					currentError.Message = err.Error()
					return err
				}

				parsedProperty, err := ParseProperty(prop)
				if err != nil {
					currentError.Message = err.Error()
					return err
				}

				e.Properties[propertyName] = parsedProperty

				return nil
			})

			if err != nil {
				currentError.Message = err.Error()
				break
			}

		}
	}, paths...)

	return e, nil
}

// ParseProperty expects the input to be valid JSON and to be a single Property record. See either the testing file for sample
// property json records or http://tinkerpop.apache.org/docs/3.4.2/dev/io/#_property_3.
func ParseProperty(in []byte) (property Property, err error) {
	if typeName, err := jsonparser.GetString(in, "@type"); err != nil || typeName != propertyTypeName {
		return property, graphson.ParsingError{err, "@type", "parseProperty"}
	}

	// value location mapping on original json record, using the jsonparser package to avoid as much reflection as we can
	var paths = [][]string{
		{"@value", "key"},
		{"@value", "value", "@value"},
	}

	parsingErrors := graphson.ParsingErrors{}

	jsonparser.EachKey(in, func(idx int, value []byte, vt jsonparser.ValueType, err error) {
		currentError := graphson.ParsingError{nil, strings.Join(paths[idx], " "), "parseVertex"}

		if err != nil {
			currentError.Message = err.Error()
			parsingErrors = append(parsingErrors, currentError)
			return
		}

		switch idx {
		case 0: // @value -> @value -> key
			key, e := jsonparser.ParseString(value)
			if e != nil {
				currentError.Message = e.Error()
				break
			}

			property.Key = key

		case 1: // @value -> @value -> value
			val, e := parsedToType(value, vt)
			if e != nil {
				currentError.Message = e.Error()
				break
			}

			property.Value = val
		}

		if currentError.Message != nil {
			parsingErrors = append(parsingErrors, currentError)
		}

	}, paths...)

	return property, parsingErrors.Combine()
}
