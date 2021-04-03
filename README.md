# jsont

This package was created to choose dynamically the struct fields to marshal.

Forked from `go version go1.16.3 windows/amd64`

## Usage

### Example using fields

```go
package main

import (
    "fmt"

    "github.com/jtorz/jsont/v2"
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

### Example using Encoder

```go
package main

import (
    "fmt"
    "log"
    "net/http"

    "github.com/jtorz/jsont/v2"
)

type SearchResult struct {
    Date        string      `json:"date"`
    IdCompany   int         `json:"idCompany"`
    Company     string      `json:"company"`
    IdIndustry  interface{} `json:"idIndustry"`
    Industry    string      `json:"industry"`
    IdContinent interface{} `json:"idContinent"`
    Continent   string      `json:"continent"`
    IdCountry   interface{} `json:"idCountry"`
    Country     string      `json:"country"`
    IdState     interface{} `json:"idState"`
    State       string      `json:"state"`
    IdCity      interface{} `json:"idCity"`
    City        string      `json:"city"`
} //SearchResult

type SearchResults struct {
    NumberResults int            `json:"numberResults"`
    Results       []SearchResult `json:"results"`
} //type SearchResults
func main() {
    msg := SearchResults{
        NumberResults: 2,
        Results: []SearchResult{
            {
                Date:        "12-12-12",
                IdCompany:   1,
                Company:     "alfa",
                IdIndustry:  1,
                Industry:    "IT",
                IdContinent: 1,
                Continent:   "america",
                IdCountry:   1,
                Country:     "México",
                IdState:     1,
                State:       "CDMX",
                IdCity:      1,
                City:        "Atz",
            },
            {
                Date:        "12-12-12",
                IdCompany:   2,
                Company:     "beta",
                IdIndustry:  1,
                Industry:    "IT",
                IdContinent: 1,
                Continent:   "america",
                IdCountry:   2,
                Country:     "USA",
                IdState:     2,
                State:       "TX",
                IdCity:      2,
                City:        "XYZ",
            },
        },
    }
    fmt.Println(msg)
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

        //{"numberResults":2,"results":[{"date":"12-12-12","idCompany":1,"idIndustry":1,"country":"México"},{"date":"12-12-12","idCompany":2,"idIndustry":1,"country":"USA"}]}
        err := jsont.NewEncoder(w).EncodeFields(msg, jsont.F{
            "numberResults": nil,
            "results": jsont.F{
                "date":       nil,
                "idCompany":  nil,
                "idIndustry": nil,
                "country":    nil,
            },
        })
        if err != nil {
            log.Fatal(err)
        }
    })

    http.ListenAndServe(":3009", nil)
}

```

#### Request output

```json
{
    "numberResults": 2,
    "results": [
        {
            "date": "12-12-12",
            "idCompany": 1,
            "idIndustry": 1,
            "country": "México"
        },
        {
            "date": "12-12-12",
            "idCompany": 2,
            "idIndustry": 1,
            "country": "USA"
        }
    ]
}
```
