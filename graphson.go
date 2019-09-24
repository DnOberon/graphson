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
// currently covered. Visit http://tinkerpop.apache.org/docs/3.4.2/dev/io/#graphson-3d0 for more information
package graphson

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
