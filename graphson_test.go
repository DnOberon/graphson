package graphson

const valuePairSimple30 = `{"@type": "g:Int64","@value": 1}`

const vertex30 = `{
  "@type" : "g:Vertex",
  "@value" : {
    "id" : {
      "@type" : "g:Int32",
      "@value" : 1
    },
    "label" : "person",
    "properties" : {
      "name" : [ {
        "@type" : "g:VertexProperty",
        "@value" : {
          "id" : {
            "@type" : "g:Int64",
            "@value" : 0
          },
          "value" : "marko",
          "label" : "name"
        }
      } ],
      "location" : [ {
        "@type" : "g:VertexProperty",
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
        "@type" : "g:VertexProperty",
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
        "@type" : "g:VertexProperty",
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
        "@type" : "g:VertexProperty",
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
   "@type":"g:VertexProperty",
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

const property = `{
  "@type" : "g:Property",
  "@value" : {
    "key" : "since",
    "value" : {
      "@type" : "g:Int32",
      "@value" : 2009
    }
  }
}`
