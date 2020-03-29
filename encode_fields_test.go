package jsont

import (
	"testing"
)

type user struct {
	ID      int    `json:"id,default"`
	Name    string `json:"name,default"`
	Surname string `json:"surname"`
	Age     int    `json:"age"`
	Rol     rol    `json:"rol,default"`
	Friend  *user  `json:"friend,default"`
}
type rol struct {
	ID    int    `json:"id,default"`
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
	p := user{1, "Paul", "McCartney", 19, rol{1, "admin", group{1, "ABC4", 12.9}}, nil}
	j := user{2, "John", "Lennon", 20, rol{1, "admin", group{1, "ABC4", 12.9}}, &p}

	m := user{3, "Mick", "Jagger", 18, rol{1, "admin", group{1, "ABC2", 12.9}}, nil}
	k := user{4, "Keith", "Richards", 17, rol{1, "admin", group{1, "ABC2", 12.9}}, &m}

	fields := F{
		"id":   nil,
		"name": nil,
		"rol": F{
			"rol_name": nil,
			"group":    F{"key": nil},
		},
	}
	want := `{"id":2,"name":"John","rol":{"rol_name":"admin","group":{"key":"ABC4"}}}`
	bytes, err := Marshal(j, fields)
	if err != nil {
		t.Fatalf("Marshal(j): %v", err)
	}
	if got := string(bytes); got != want {
		t.Errorf("Marshal(j) = %#q, want %#q", got, want)
	}

	want = `[{"id":2,"name":"John","rol":{"rol_name":"admin","group":{"key":"ABC4"}}},{"id":4,"name":"Keith","rol":{"rol_name":"admin","group":{"key":"ABC2"}}}]`
	bytes, err = Marshal([]user{j, k}, fields)
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
	bytes, err = Marshal([]user{j, k}, withFriends)
	if err != nil {
		t.Fatalf("Marshal([]user{j, k}, withFriends): %v", err)
	}
	if got := string(bytes); got != want {
		t.Errorf("Marshal([]user{j, k}, withFriends) = %#q, want %#q", got, want)
	}

	want = `[{"id":2,"name":"John","rol":{"id":1},"friend":{"id":1,"name":"Paul","rol":{"id":1},"friend":null}},{"id":4,"name":"Keith","rol":{"id":1},"friend":{"id":3,"name":"Mick","rol":{"id":1},"friend":null}}]`
	bytes, err = Marshal([]user{j, k}, Defaults)
	if err != nil {
		t.Fatalf("Marshal([]user{j, k}, Defaults): %v", err)
	}
	if got := string(bytes); got != want {
		t.Errorf("Marshal([]user{j, k}, Defaults) = %#q, want %#q", got, want)
	}
}
