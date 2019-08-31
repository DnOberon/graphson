package graphson

import (
	"strings"

	"github.com/buger/jsonparser"
)

const vertexTypeName = "g:Vertex"
const vertexPropertyTypeName = "g:VertexProperty"

type Vertex struct {
	ID         interface{}                 `json:"id"`
	Label      string                      `json:"label"`
	Properties map[string][]VertexProperty `json:"properties"`
}

type VertexProperty struct {
	ID         interface{}            `json:"id"`
	Value      string                 `json:"value"`
	Label      string                 `json:"label"`
	Properties map[string]interface{} `json:"properties"`
}

func parseVertex(in []byte) (v Vertex, err error) {
	v.Properties = map[string][]VertexProperty{}

	if typeName, err := jsonparser.GetString(in, "@type"); err != nil || typeName != vertexTypeName {
		return v, ParsingError{err, "@type", "parseVertex"}
	}

	var paths = [][]string{
		{"@value", "label"},
		{"@value", "id", "@value"},
		{"@value", "properties"},
	}

	parsingErrors := ParsingErrors{}

	jsonparser.EachKey(in, func(idx int, value []byte, vt jsonparser.ValueType, err error) {
		currentError := ParsingError{nil, strings.Join(paths[idx], " "), "parseVertex"}

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
			id, e := parsedToType(in, vt)
			if e != nil {
				currentError.Message = e.Error()
				break
			}

			v.ID = id

		case 2: // @value -> properties (VertexProperty)
			_, e := jsonparser.ArrayEach(value, func(prop []byte, dataType jsonparser.ValueType, offset int, err error) {
				propertyName, e := jsonparser.ParseString(prop)
				if e != nil {
					currentError.Message = e.Error()
					return
				}

				property, _, _, e := jsonparser.Get(value, propertyName)
				if e != nil {
					currentError.Message = e.Error()
					return
				}

				parsedProperties, e := parseVertexProperties(property)
				if e != nil {
					currentError.Message = e.Error()
					return
				}

				v.Properties[propertyName] = append(v.Properties[propertyName], parsedProperties...)
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

	return v, parsingErrors.combine()
}

func parseVertexProperties(in []byte) (properties []VertexProperty, err error) {
	parsingErrors := ParsingErrors{}

	_, err = jsonparser.ArrayEach(in, func(prop []byte, dataType jsonparser.ValueType, offset int, err error) {
		currentError := ParsingError{nil, "parseVertexProperties", "properties"}

		propertyName, e := jsonparser.ParseString(prop)
		if e != nil {
			currentError.Message = e.Error()
			return
		}

		property, _, _, e := jsonparser.Get(in, propertyName)
		if e != nil {
			currentError.Message = e.Error()
			return
		}

		_, e = jsonparser.ArrayEach(property, func(p []byte, dataType jsonparser.ValueType, offset int, err error) {
			parsedProperty, e := parseVertexProperty(p)
			if e != nil {
				currentError.Message = e.Error()
				return
			}

			properties = append(properties, parsedProperty)
		})

		if e != nil {
			currentError.Message = e.Error()
			return
		}

	})

	if err != nil {
		return properties, ParsingError{err.Error(), "parseVertexProperties", "properties"}
	}

	return properties, parsingErrors.combine()
}

func parseVertexProperty(in []byte) (property VertexProperty, err error) {
	property.Properties = map[string]interface{}{}

	if typeName, err := jsonparser.GetString(in, "@type"); err != nil || typeName != vertexPropertyTypeName {
		return property, ParsingError{err, "@type", "parseVertexProperty"}
	}

	var paths = [][]string{
		{"@value", "label"},
		{"@value", "id", "@value"},
		{"@value", "value"},
		{"@value", "properties"},
	}

	parsingErrors := ParsingErrors{}

	jsonparser.EachKey(in, func(idx int, value []byte, vt jsonparser.ValueType, err error) {
		currentError := ParsingError{nil, strings.Join(paths[idx], " "), "parseVertex"}

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
			_, e := jsonparser.ArrayEach(value, func(prop []byte, dataType jsonparser.ValueType, offset int, err error) {
				propertyName, e := jsonparser.ParseString(prop)
				if e != nil {
					currentError.Message = e.Error()
					return
				}

				p, pType, _, e := jsonparser.Get(value, propertyName, "@value")
				if e != nil {
					currentError.Message = e.Error()
					return
				}

				property.Properties[propertyName], e = parsedToType(p, pType)
				if e != nil {
					currentError.Message = e.Error()
					return
				}
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

	return property, parsingErrors.combine()
}
