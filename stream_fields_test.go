// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package jsont

import (
	"bytes"
	"testing"
)

func TestEncoderFields(t *testing.T) {
	type streamTest struct {
		ID    int
		Name  string
		Other string
	}
	type streamTest1 struct {
		ID        int
		Name      string
		Other     string
		Recursive []streamTest1
		Stream    *streamTest
	}

	st := streamTest1{
		ID:        1,
		Name:      "test",
		Other:     "other value",
		Recursive: []streamTest1{{ID: 2, Name: "test2", Recursive: []streamTest1{}}},
		Stream:    &streamTest{ID: 123, Name: "st", Other: "other"},
	}
	streamTestFields := []interface{}{st, st, st}

	whitelists := []F{
		{"ID": nil, "Recursive": Recursive},
		{"ID": nil, "Recursive": Recursive, "Stream": nil},
		{"ID": nil, "Name": nil},
	}

	streamEncodedFields := `{"ID":1,"Recursive":[{"ID":2,"Recursive":[]}]}
{"ID":1,"Recursive":[{"ID":2,"Recursive":[],"Stream":null}],"Stream":{"ID":123,"Name":"st","Other":"other"}}
{"ID":1,"Name":"test"}
`

	for i := 0; i <= len(streamTestFields); i++ {
		var buf bytes.Buffer
		enc := NewEncoder(&buf)
		// Check that enc.SetIndent("", "") turns off indentation.
		enc.SetIndent(">", ".")
		enc.SetIndent("", "")
		for j, v := range streamTestFields[0:i] {
			whitelist := whitelists[j]
			if err := enc.EncodeFields(v, whitelist); err != nil {
				t.Fatalf("encode #%d: %v", j, err)
			}
		}
		if have, want := buf.String(), nlines(streamEncodedFields, i); have != want {
			t.Errorf("encoding %d items: mismatch", i)
			diff(t, []byte(have), []byte(want))
			break
		}
	}
}
