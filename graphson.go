// Copyright [2019] [John Darrington]

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

// http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package graphson provides a parser for GraphSON structured json data. This interface is primarily consumed by the
// gremlin go language variant. It provides a GraphSON version agnostic interface to the user, though only GraphSON 3.0+ is
// currently covered. Visit http://tinkerpop.apache.org/docs/3.4.2/dev/io/#graphson-3d0 for more information. This is not
// intended to handle things like
package graphson

import (
	"sync"
	"time"
)

var (
	parsersMu sync.RWMutex
	parsers   = make(map[string]GraphSONParser)
)

// RegisterParser allows an outside package to register a GraphSONParser compatible type with this package.
func RegisterParser(name string, parser GraphSONParser) {
	parsersMu.Lock()
	defer parsersMu.Unlock()

	if parser == nil {
		panic("GraphSONParser is nil for provider " + name)
	}

	parsers[name] = parser
}

func NewParser(parserVersion string) GraphSONParser {
	return parsers[parserVersion]
}

// ValueType represents the GraphSON equivalent type of a value in a ValuePair type.
type ValueType int

const (
	String = ValueType(iota)
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

// GraphSONParser enforces a standard set of functions that a GraphSON parser must satisfy. It is up to the individual
// implementer to handle the parsing of any data types apart from Vertex, Vertex Property, Edge, and Property.
// *Note:* TinkerGraph and Path are absent from this list and interface as we believe they're considered legacy types.
type GraphSONParser interface {
	Parse(in []byte) (ValuePair, error)
	ParseVertex(in []byte) (VertexRecord, error)
	ParseVertexProperties(in []byte) ([]VertexPropertyRecord, error)
	ParseVertexProperty(in []byte) (VertexPropertyRecord, error)
	ParseEdge(in []byte) (EdgeRecord, error)
	ParseProperty(in []byte) (Property, error)
}

// VertexRecord mirrors the basic Vertex record structure defined by GraphSON and Gremlin.
type VertexRecord struct {
	ID         interface{}                       `json:"id"` // ID as interface{}, different providers use different ID types
	Label      string                            `json:"label"`
	Properties map[string][]VertexPropertyRecord `json:"properties"`
}

// VertexPropertyRecord mirrors the basic VertexRecord Property structure defined by GraphSON and Gremlin.
type VertexPropertyRecord struct {
	ID         interface{}          `json:"id"` // ID as interface{}, different providers use different ID types
	Value      string               `json:"value"`
	Label      string               `json:"label"`
	Properties map[string]ValuePair `json:"properties"`
}

// EdgeRecord mirrors the basic Edge record defined by GraphSON and Gremlin.
type EdgeRecord struct {
	ID         interface{}         `json:"id"` // ID as interface{}, different providers use different ID types
	Label      string              `json:"label"`
	InVLabel   string              `json:"inVLabel"`
	OutVLabel  string              `json:"outVLabel"`
	InV        interface{}         `json:"inV"`
	OutV       interface{}         `json:"outV"`
	Properties map[string]Property `json:"properties"`
}

// Property reflects a common GraphSON pattern of Key - Type Value data representation.
type Property struct {
	Key   string    `json:"key"`
	Value ValuePair `json:"value"`
}

// ValuePair contains data and information about the shape of that data. Methods exist for extracting concrete data types
// from a ValuePair's contained Value. All parsing of the Value should have taken place by the original parser and nothing
// but organization and type inference should happen after this data structure has been created.
type ValuePair struct {
	Type  ValueType
	Value interface{}
}

func (vp ValuePair) AsVertex() VertexRecord {
	if vp.Type != Vertex {
		return VertexRecord{}
	}

	return vp.Value.(VertexRecord)
}

func (vp ValuePair) AsVertexProperty() VertexPropertyRecord {
	if vp.Type != VertexProperty {
		return VertexPropertyRecord{}
	}

	return vp.Value.(VertexPropertyRecord)
}

func (vp ValuePair) AsEdge() EdgeRecord {
	if vp.Type != Edge {
		return EdgeRecord{}
	}

	return vp.Value.(EdgeRecord)
}

func (vp ValuePair) AsProperty() Property {
	if vp.Type != EdgeProperty {
		return Property{}
	}

	return vp.Value.(Property)
}

func (vp ValuePair) AsSet() []ValuePair {
	if vp.Type != Set {
		return nil
	}

	return vp.Value.([]ValuePair)
}

func (vp ValuePair) AsMap() map[ValuePair]ValuePair {
	if vp.Type != Map {
		return nil
	}

	return vp.Value.(map[ValuePair]ValuePair)
}

func (vp ValuePair) AsString() string {
	if vp.Type != String {
		return ""
	}

	return vp.Value.(string)
}

func (vp ValuePair) AsInt32() int {
	if vp.Type != Int32 {
		return 0
	}

	return vp.Value.(int)
}

func (vp ValuePair) AsInt64() int64 {
	if vp.Type != Int64 {
		return 0
	}

	return vp.Value.(int64)
}

func (vp ValuePair) AsFloat32() float32 {
	if vp.Type != Double {
		return 0
	}

	return vp.Value.(float32)
}

func (vp ValuePair) AsFloat64() float64 {
	if vp.Type != Float {
		return 0
	}

	return vp.Value.(float64)
}
func (vp ValuePair) AsTime() time.Time {
	if vp.Type != Timestamp {
		return time.Time{}
	}

	return vp.Value.(time.Time)
}

func (vp ValuePair) AsClAss() string {
	if vp.Type != Class {
		return ""
	}

	return vp.Value.(string)
}

func (vp ValuePair) AsUUID() string {
	if vp.Type != UUID {
		return ""
	}

	return vp.Value.(string)
}
