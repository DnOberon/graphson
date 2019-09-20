package graphson3

const valuePairSimple30 = `{"@type": "g:Int64","@value": 1}`

const vertex30 = `{
  "@type" : "g:VertexRecord",
  "@value" : {
    "id" : {
      "@type" : "g:Int32",
      "@value" : 1
    },
    "label" : "person",
    "properties" : {
      "name" : [ {
        "@type" : "g:VertexPropertyRecord",
        "@value" : {
          "id" : {
            "@type" : "g:Int64",
            "@value" : 1
          },
          "value" : "marko",
          "label" : "name"
        }
      } ],
      "location" : [ {
        "@type" : "g:VertexPropertyRecord",
        "@value" : {
          "id" : {
            "@type" : "g:Int64",
            "@value" : 6
          },
          "value" : "san diego",
          "label" : "location",
          "properties" : {
            "startTime" : {
              "@type" : "g:Int32",
              "@value" : 1997
            },
            "endTime" : {
              "@type" : "g:Int32",
              "@value" : 2001
            }
          }
        }
      }, {
        "@type" : "g:VertexPropertyRecord",
        "@value" : {
          "id" : {
            "@type" : "g:Int64",
            "@value" : 7
          },
          "value" : "santa cruz",
          "label" : "location",
          "properties" : {
            "startTime" : {
              "@type" : "g:Int32",
              "@value" : 2001
            },
            "endTime" : {
              "@type" : "g:Int32",
              "@value" : 2004
            }
          }
        }
      }, {
        "@type" : "g:VertexPropertyRecord",
        "@value" : {
          "id" : {
            "@type" : "g:Int64",
            "@value" : 8
          },
          "value" : "brussels",
          "label" : "location",
          "properties" : {
            "startTime" : {
              "@type" : "g:Int32",
              "@value" : 2004
            },
            "endTime" : {
              "@type" : "g:Int32",
              "@value" : 2005
            }
          }
        }
      }, {
        "@type" : "g:VertexPropertyRecord",
        "@value" : {
          "id" : {
            "@type" : "g:Int64",
            "@value" : 9
          },
          "value" : "santa fe",
          "label" : "location",
          "properties" : {
            "startTime" : {
              "@type" : "g:Int32",
              "@value" : 2005
            }
          }
        }
      } ]
    }
  }
}`

const vertexProperty30 = `{
   "@type":"g:VertexPropertyRecord",
   "@value":{
      "id":{
         "@type":"g:Int64",
         "@value":6
      },
      "value":"san diego",
      "label":"location",
      "properties":{
         "startTime":{
            "@type":"g:Int32",
            "@value":1997
         },
         "endTime":{
            "@type":"g:Int32",
            "@value":2001
         }
      }
   }
}`

const edge30 = `{
  "@type" : "g:Edge",
  "@value" : {
    "id" : {
      "@type" : "g:Int32",
      "@value" : 13
    },
    "label" : "develops",
    "inVLabel" : "software",
    "outVLabel" : "person",
    "inV" : {
      "@type" : "g:Int32",
      "@value" : 10
    },
    "outV" : {
      "@type" : "g:Int32",
      "@value" : 1
    },
    "properties" : {
      "since" : {
        "@type" : "g:Property",
        "@value" : {
          "key" : "since",
          "value" : {
            "@type" : "g:Int32",
            "@value" : 2009
          }
        }
      }
    }
  }
}`

const property30 = `{
  "@type" : "g:Property",
  "@value" : {
    "key" : "since",
    "value" : {
      "@type" : "g:Int32",
      "@value" : 2009
    }
  }
}`

const set30 = `{
  "@type" : "g:Set",
  "@value" : [ {
    "@type" : "g:Int32",
    "@value" : 1
  }, "person", true ]
}`

const class30 = `{
  "@type" : "g:Class",
  "@value" : "java.io.File"
}`

const date30 = `{
  "@type" : "g:Date",
  "@value" : 1481750076295
}`

const double30 = `{
  "@type" : "g:Double",
  "@value" : 100.0
}`

const float30 = `{
  "@type" : "g:Float",
  "@value" : 100.0
}`

const integer30 = `{
  "@type" : "g:Int32",
  "@value" : 100
}`

const long30 = `{
  "@type" : "g:Int64",
  "@value" : 100
}`

const map30 = `{
  "@type" : "g:Map",
  "@value" : [ {
    "@type" : "g:Date",
    "@value" : 1481750076295
  }, "red", {
    "@type" : "g:List",
    "@value" : [ {
      "@type" : "g:Int32",
      "@value" : 1
    }, {
      "@type" : "g:Int32",
      "@value" : 2
    }, {
      "@type" : "g:Int32",
      "@value" : 3
    } ]
  }, {
    "@type" : "g:Date",
    "@value" : 1481750076295
  }, "test", {
    "@type" : "g:Int32",
    "@value" : 123
  } ]
}`

const timestamp30 = `{
  "@type" : "g:Timestamp",
  "@value" : 1481750076295
}`

const uuid30 = `{
  "@type" : "g:UUID",
  "@value" : "41d2e28a-20a4-4ab0-b379-d810dede3786"
}`
