# GraphSON

[![GoDoc Widget]][GoDoc]

GraphSON Go is a Go library for the parsing and interpretation of GraphSON structured JSON. See the [Gremlin IO Reference](http://tinkerpop.apache.org/docs/3.4.2/dev/io/#graphson) for more information on GraphSON and its position in the Gremlin query hierarchy. 

This library is meant to be used along your own implementation of Gremlin's HTTP API interface functionality. 


### Setup 

In order to use the parser you must import both this package and the package of the desired GraphSON version. The package handles the registration of the imported versions automatically.
```
import (
    "github.com/dnoberon/graphson"
    _ "github.com/dnoberon/graphson/graphson3"
 )

```

If you have correctly added your packages you can initialize the version 3 GraphSON parser.
```
parser := graphsonNewParser("v3")
```
<br>

### Usage

While the `GraphSONParser` interface provides methods for parsing Vertex, Edge, Vertex Property, and Property data types it is suggested that the user of this package only use the `Parse` method and then the subsequent methods on the return `ValuePair` type.

```
const timestamp = `{
  "@type" : "g:Timestamp",
  "@value" : 1481750076295
}`

valuePair, err := parser.Parse([]byte(timestamp))

fmt.Println(valuePair.Type) // graphson.Timestamp
validGoTimeStruct := valuePair.AsTimestamp()
```



[GoDoc]: https://godoc.org/github.com/DnOberon/graphson
[GoDoc Widget]: https://godoc.org/github.com/DnOberon/graphson?status.svg
