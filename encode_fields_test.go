package jsont

import (
	"testing"
)

type user struct {
	ID        int           `json:"id"`
	Name      string        `json:"name"`
	Surname   string        `json:"surname"`
	Age       int           `json:"age"`
	Rol       rol           `json:"rol"`
	Friend    *user         `json:"friend"`
	Marshaler marshalerTest `json:"marshaler"`
}

type marshalerTest struct{}

func (m marshalerTest) MarshalJSONFields(whitelist F) ([]byte, error) {
	return []byte(`""`), nil
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

func TestWithFields(t *testing.T) {
	var err error
	p := user{1, "Paul", "McCartney", 19, rol{1, "admin", group{1, "ABC4", 12.9}}, nil, marshalerTest{}}
	j := user{2, "John", "Lennon", 20, rol{1, "admin", group{1, "ABC4", 12.9}}, &p, marshalerTest{}}

	m := user{3, "Mick", "Jagger", 18, rol{1, "admin", group{1, "ABC2", 12.9}}, nil, marshalerTest{}}
	k := user{4, "Keith", "Richards", 17, rol{1, "admin", group{1, "ABC2", 12.9}}, &m, marshalerTest{}}

	fields := F{
		"id":   nil,
		"name": nil,
		"rol": F{
			"rol_name": nil,
			"group":    F{"key": nil},
		},
		"marshaler": nil,
	}
	want := `{"id":2,"name":"John","rol":{"rol_name":"admin","group":{"key":"ABC4"}},"marshaler":""}`
	bytes, err := MarshalFields(j, fields)
	if err != nil {
		t.Fatalf("Marshal(j): %v", err)
	}
	if got := string(bytes); got != want {
		t.Errorf("Marshal(j) = %#q, want %#q", got, want)
	}

	want = `[{"id":2,"name":"John","rol":{"rol_name":"admin","group":{"key":"ABC4"}},"marshaler":""},{"id":4,"name":"Keith","rol":{"rol_name":"admin","group":{"key":"ABC2"}},"marshaler":""}]`
	bytes, err = MarshalFields([]user{j, k}, fields)
	if err != nil {
		t.Fatalf("Marshal([]user{j, k}): %v", err)
	}
	if got := string(bytes); got != want {
		t.Errorf("Marshal([]user{j, k}) = %#q, want %#q", got, want)
	}

	withFriends := F{
		"id":     nil,
		"name":   nil,
		"friend": Recursive,
	}
	want = `[{"id":2,"name":"John","friend":{"id":1,"name":"Paul","friend":null}},{"id":4,"name":"Keith","friend":{"id":3,"name":"Mick","friend":null}}]`
	bytes, err = MarshalFields([]user{j, k}, withFriends)
	if err != nil {
		t.Fatalf("Marshal([]user{j, k}, withFriends): %v", err)
	}
	if got := string(bytes); got != want {
		t.Errorf("Marshal([]user{j, k}, withFriends) = %#q, want %#q", got, want)
	}
}
