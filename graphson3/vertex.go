package graphson3

import (
	"strings"

	"github.com/dnoberon/graphson"

	"github.com/buger/jsonparser"
)

const vertexTypeName = "g:VertexRecord"
const vertexPropertyTypeName = "g:VertexPropertyRecord"

// VertexRecord mirrors the basic VertexRecord structure defined by GraphSON and Gremlin. "g:VertexRecord" as defined GraphSON 3.0 name.
type VertexRecord struct {
	ID         interface{}                       `json:"id"` // ID as interface{}, different providers use different ID types
	Label      string                            `json:"label"`
	Properties map[string][]VertexPropertyRecord `json:"properties"`
}

// VertexPropertyRecord mirrors the basic VertexRecord Property structure defined by GraphSON and Gremlin. "g:VertexPropertyRecord" as
// defined GraphSON 3.0 name.
type VertexPropertyRecord struct {
	ID         interface{}                   `json:"id"` // ID as interface{}, different providers use different ID types
	Value      string                        `json:"value"`
	Label      string                        `json:"label"`
	Properties map[string]graphson.ValuePair `json:"properties"`
}

// ParseVertex expects the input to be valid JSON and to be a single VertexRecord record. See either the testing file for sample
// vertex json records or http://tinkerpop.apache.org/docs/3.4.2/dev/io/#_vertex_3.
func (g GraphSONv3Parser) ParseVertex(in []byte) (v VertexRecord, err error) {
	v.Properties = map[string][]VertexPropertyRecord{}

	if typeName, err := jsonparser.GetString(in, "@type"); err != nil || typeName != vertexTypeName {
		return v, graphson.ParsingError{err, "@type", "parseVertex"}
	}

	// value location mapping on original json record, using the jsonparser package to avoid as much reflection as we can
	var paths = [][]string{
		{"@value", "label"},
		{"@value", "id", "@value"},
		{"@value", "properties"},
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
		case 0: // @value -> label
			label, e := jsonparser.ParseString(value)
			if e != nil {
				currentError.Message = e.Error()
				break
			}

			v.Label = label

		case 1: // @value -> label -> @value
			id, e := parsedToType(value, vt)
			if e != nil {
				currentError.Message = e.Error()
				break
			}

			v.ID = id

		case 2: // @value -> properties (VertexPropertyRecord)
			e := jsonparser.ObjectEach(value, func(key []byte, prop []byte, dataType jsonparser.ValueType, offset int) error {
				propertyName, e := jsonparser.ParseString(key)
				if e != nil {
					currentError.Message = e.Error()
					return e
				}

				parsedProperties, e := g.parseVertexProperties(prop)
				if e != nil {
					currentError.Message = e.Error()
					return e
				}

				v.Properties[propertyName] = append(v.Properties[propertyName], parsedProperties...)

				return nil
			})

			if e != nil {
				currentError.Message = e.Error()
				break
			}
		}

		if currentError.Message != nil {
			parsingErrors = append(parsingErrors, currentError)
		}

	}, paths...)

	return v, parsingErrors.Combine()
}

func (g GraphSONv3Parser) parseVertexProperties(in []byte) ([]VertexPropertyRecord, error) {
	properties := []VertexPropertyRecord{}
	parsingErrors := graphson.ParsingErrors{}

	_, err := jsonparser.ArrayEach(in, func(prop []byte, dataType jsonparser.ValueType, offset int, err error) {
		currentError := graphson.ParsingError{nil, "parseVertexProperties", "properties"}

		parsedProperty, e := g.ParseVertexProperty(prop)
		if e != nil {
			currentError.Message = e.Error()
			return
		}

		properties = append(properties, parsedProperty)

		if e != nil {
			currentError.Message = e.Error()
			return
		}

	})

	if err != nil {
		return properties, graphson.ParsingError{err.Error(), "parseVertexProperties", "properties"}
	}

	return properties, parsingErrors.Combine()
}

// ParseVertexProperty expects the input to be valid JSON and to be a single VertexRecord Property record. See either the testing file for sample
// vertex json records or http://tinkerpop.apache.org/docs/3.4.2/dev/io/#_vertexproperty_3.
func (g GraphSONv3Parser) ParseVertexProperty(in []byte) (property VertexPropertyRecord, err error) {
	property.Properties = map[string]graphson.ValuePair{}

	if typeName, err := jsonparser.GetString(in, "@type"); err != nil || typeName != vertexPropertyTypeName {
		return property, graphson.ParsingError{err, "@type", "parseVertexProperty"}
	}

	// value location mapping on original json record, using the jsonparser package to avoid as much reflection as we can
	var paths = [][]string{
		{"@value", "label"},
		{"@value", "id", "@value"},
		{"@value", "value"},
		{"@value", "properties"},
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
		case 0: // @value -> label
			label, e := jsonparser.ParseString(value)
			if e != nil {
				currentError.Message = e.Error()
				break
			}

			property.Label = label

		case 1: // @value -> id -> @value
			id, e := parsedToType(value, vt)
			if e != nil {
				currentError.Message = e.Error()
				break
			}

			property.ID = id

		case 2: // @value -> value
			pValue, e := jsonparser.ParseString(value)
			if e != nil {
				currentError.Message = e.Error()
				break
			}

			property.Value = pValue

		case 3: // @value -> properties
			e := jsonparser.ObjectEach(value, func(key []byte, prop []byte, dataType jsonparser.ValueType, offset int) error {
				propertyName, e := jsonparser.ParseString(key)
				if e != nil {
					currentError.Message = e.Error()
					return currentError
				}

				property.Properties[propertyName], e = g.Parse(prop)
				if e != nil {
					currentError.Message = e.Error()
					return currentError
				}

				return nil
			})

			if e != nil {
				currentError.Message = e.Error()
				break
			}

		}

		if currentError.Message != nil {
			parsingErrors = append(parsingErrors, currentError)
		}

	}, paths...)

	return property, parsingErrors.Combine()
}
