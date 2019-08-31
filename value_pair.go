package graphson

import (
	"reflect"
	"strings"

	"github.com/buger/jsonparser"
)

type ValuePair struct {
	ValueTypeName string `json:"@type"`
	Value         []byte ` json:"@value"`
	ValueType     reflect.Type
}

type ResultList []ValuePair

func parseValuePair(in []byte) (ValuePair, error) {
	v := ValuePair{}
	var paths = [][]string{
		{"@type"},
		{"@value"},
	}

	parsingErrors := ParsingErrors{}

	jsonparser.EachKey(in, func(idx int, value []byte, vt jsonparser.ValueType, err error) {
		currentError := ParsingError{nil, strings.Join(paths[idx], " "), "parseVertex"}

		switch idx {
		case 0:
			typeName, e := jsonparser.ParseString(value)
			if e != nil {
				currentError.Message = e.Error()
				break
			}

			v.ValueTypeName = typeName
		case 1:
			_, e := parsedToType(value, vt)
			if e != nil {
				currentError.Message = e.Error()
				break
			}

		}

		if currentError.Message != nil {
			parsingErrors = append(parsingErrors, currentError)
		}
	}, paths...)

	return v, nil
}
