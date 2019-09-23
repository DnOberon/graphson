package graphson

type ValuePair struct {
	Type  ValueType
	Value interface{}
}

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

// VertexRecord mirrors the basic VertexRecord structure defined by GraphSON and Gremlin. "g:VertexRecord" as defined GraphSON 3.0 name.
type VertexRecord struct {
	ID         interface{}                       `json:"id"` // ID as interface{}, different providers use different ID types
	Label      string                            `json:"label"`
	Properties map[string][]VertexPropertyRecord `json:"properties"`
}

// VertexPropertyRecord mirrors the basic VertexRecord Property structure defined by GraphSON and Gremlin. "g:VertexPropertyRecord" as
// defined GraphSON 3.0 name.
type VertexPropertyRecord struct {
	ID         interface{}          `json:"id"` // ID as interface{}, different providers use different ID types
	Value      string               `json:"value"`
	Label      string               `json:"label"`
	Properties map[string]ValuePair `json:"properties"`
}

type EdgeRecord struct {
	ID         interface{}         `json:"id"` // ID as interface{}, different providers use different ID types
	Label      string              `json:"label"`
	InVLabel   string              `json:"inVLabel"`
	OutVLabel  string              `json:"outVLabel"`
	InV        interface{}         `json:"inV"`
	OutV       interface{}         `json:"outV"`
	Properties map[string]Property `json:"properties"`
}

type Property struct {
	Key   string    `json:"key"`
	Value ValuePair `json:"value"`
}
