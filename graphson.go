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

// TODO: Handle the following separately - List, Map, Site, Timestamp, Date
// TODO: I like the pattern of returning a concrete type alias and providing methods for getting values out of them, copy that
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
