# jsonreader(in Go)

jsonreader is a simple go library that is used to read valid json files in a folder or a json file and merge them to a map. The value(s) for a particular key is returned either as a map or a string. jsonreader also returns a value if a dot seperated keys are passed in. Some of the basic usage below -


To read all valid json files in a folder into a map

```
import (
 github.com/gocoolgeek/jsonreader
)

jsonreader.Load("/go/samples/")

```

To read just one valid json file into a map

```
import (
 github.com/gocoolgeek/jsonreader
)

jsonreader.Load("/go/samples/blah.json")

```

Assume a following `blah.json` file in the folder `/go/samples/blah.json`:


```
{
    "glossary": {
        "title": "example glossary",
		"GlossDiv": {
            "title": "S",
			"GlossList": {
                "GlossEntry": {
                    "ID": "SGML",
					"SortAs": "SGML",
					"GlossTerm": "Standard Generalized Markup Language",
					"Acronym": "SGML",
					"Abbrev": "ISO 8879:1986",
					"GlossDef": {
                        "para": "A meta-markup language, used to create markup languages such as DocBook.",
						"GlossSeeAlso": ["GML", "XML"]
                    },
					"GlossSee": "markup"
                }
            }
        }
    }
}

```

To read the `GlossDiv` key as a map, the usage something like below,


```
import (
 github.com/gocoolgeek/jsonreader
)

jsonreader.Load("/go/samples/blah.json")
glossary := jsonreader.GetMap("glossary")
glossdiv := jsonreader.TransformInterfaceToMap(glossary["GlossDiv"])


```

To read the `GlossSee` key's value the usage is below,

```
import (
 github.com/gocoolgeek/jsonreader
)

jsonreader.Load("/go/samples/blah.json")
glossary := jsonreader.GetValue("glossary.GlossDiv.GlossList.GlossEntry.GlossSee")

```

For further references please refer the tests
