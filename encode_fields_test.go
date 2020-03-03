package jsont

import "testing"

type user struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Age     int    `json:"age"`
	Rol     rol    `json:"rol"`
	Friend  *user  `json:"friend"`
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
	f := user{1, "Paul", "McCartney", 19, rol{1, "admin", group{1, "ABC4", 12.9}}, nil}
	u := user{2, "John", "Lennon", 20, rol{1, "admin", group{1, "ABC4", 12.9}}, &f}
	fields := F{
		"id":   nil,
		"name": nil,
		"rol": F{
			"rol_name": nil,
			"group":    F{"key": nil},
		},
	}
	_, err = Marshal(u, fields)
	if err != nil {
		t.Error(err)
	}

	fields["rol"]["group"]["number_anything"] = nil
	_, err = Marshal([]user{u, u, u}, fields)
	if err != nil {
		t.Error(err)
	}

	_, err = Marshal([]user{u, u, u}, nil)
	if err != nil {
		t.Error(err)
	}

	fields = F{
		"id":     nil,
		"name":   nil,
		"friend": Recursive,
	}
	_, err = Marshal([]user{u, u, u}, fields)
	if err != nil {
		t.Error(err)
	}
}
