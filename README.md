# jsont

This package was created to choose dynamically the struct fields to marshal.

Forked from `go version go1.16.3 windows/amd64`

## Usage

### Example using fields

```go
package main

import (
    "fmt"

    "github.com/jtorz/jsont"
)

type user struct {
    ID      int    `json:"id"`
    Name    string `json:"name"`
    Surname string `json:"surname"`
    Age     int    `json:"age"`
    Rol     rol    `json:"rol"`
    Friend  *user  `json:"friend,omitempty"`
}
type rol struct {
    ID    int    `json:"id"`
    Rol   string `json:"rol_name"`
    Group group  `json:"group"`
}
type group struct {
    ID      int     `json:"id"`
    Key     string  `json:"key"`
    Anthing float64 `json:"number_anything,string"`
}

func main() {
    var j []byte
    var err error
    f := user{1, "Paul", "McCartney", 19, rol{1, "admin", group{1, "ABC4", 12.9}}, nil}
    u := user{2, "John", "Lennon", 20, rol{1, "admin", group{1, "ABC4", 12.9}}, &f}

    //marshaling single struct
    j, err = jsont.MarshalFields(u, jsont.F{
        "id":   nil,
        "name": nil,
        "rol": jsont.F{
            "rol_name": nil,
            "group":    jsont.F{"key": nil},
        },
    })
    if err != nil {
        panic(err)
    }
    fmt.Print(string(j), "\n\n")
    // OUTPUT 1*

    //marshaling slice
    j, err = jsont.MarshalFields([]user{u, u}, jsont.F{
        "id":     nil,
        "name":   nil,
        "friend": jsont.Recursive,
    })
    if err != nil {
        panic(err)
    }
    fmt.Print(string(j), "\n\n")
    // Output 2*
}
```

#### Output 1*

```json
    {
        "id": 2,
        "name": "John",
        "rol": {
            "rol_name": "admin",
            "group": {
                "key": "ABC4"
            }
        }
    }

```

#### Output 2*

```json
[
    {
        "id": 2,
        "name": "John",
        "friend": {
            "id": 1,
            "name": "Paul"
        }
    },
    {
        "id": 2,
        "name": "John",
        "friend": {
            "id": 1,
            "name": "Paul"
        }
    }
]

```
