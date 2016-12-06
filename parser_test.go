package andrea_test

import (
   "github.com/CarolinaRomero33/Comperas"
	"io/ioutil"
	"reflect"
	"strings"
	"testing"
)

// Ensure the parser can parse strings into Statement ASTs.

func TestParser_ParseStatement(t *testing.T) {

	//Erasmo
	f, _ := ioutil.ReadFile("C:\\Users\\Romero\\Documents\\Go\\file2.txt")
	z := string(f[:])
	
	var tests = []struct {
		s    string
		stmt *andrea.SelectStatement
		err  string
	}{
		{
			s:    z,
			stmt: &andrea.SelectStatement{
			//Fields: []string{""},
			//TableName: "andrea",
			},
		},
	}
	for i, tt := range tests {
		stmt, err := andrea.NewParser(strings.NewReader(tt.s)).Parse()
		if !reflect.DeepEqual(tt.err, errstring(err)) {
			t.Errorf("%d. %q: error mismatch:\n  exp=%s\n  zt=%s\n\n", i, tt.s, tt.err, err)
		} else if tt.err == "" && !reflect.DeepEqual(tt.stmt, stmt) {
			t.Errorf("%d. %q\n\nstmt mismatch:\n\nexp=%#v\n\nzt=%#v\n\n", i, tt.s, tt.stmt, stmt)
		}
	}

	//for i, tt := range tests {
	//	stmt, err := andrea.NewParser(strings.NewReader(tt.s)).Parsesi()
	//	if !reflect.DeepEqual(tt.err, errstring(err)) {
	//		t.Errorf("%d. %q: error mismatch:\n  exp=%s\n  zt=%s\n\n", i, tt.s, tt.err, err)
	//	} else if tt.err == "" && !reflect.DeepEqual(tt.stmt, stmt) {
	//		t.Errorf("%d. %q\n\nstmt mismatch:\n\nexp=%#v\n\nzt=%#v\n\n", i, tt.s, tt.stmt, stmt)
	//	}
	//}
	//for i, tt := range tests {
	//	stmt, err := andrea.NewParser(strings.NewReader(tt.s)).Parsesi()
	//	if !reflect.DeepEqual(tt.err, errstring(err)) {
	//		t.Errorf("%d. %q: error mismatch:\n  exp=%s\n  zt=%s\n\n", i, tt.s, tt.err, err)
	//	} else if tt.err == "" && !reflect.DeepEqual(tt.stmt, stmt) {
	//		t.Errorf("%d. %q\n\nstmt mismatch:\n\nexp=%#v\n\nzt=%#v\n\n", i, tt.s, tt.stmt, stmt)
	//	}
	//}
	//for i, tt := range tests {
	//	stmt, err := andrea.NewParser(strings.NewReader(tt.s)).Parsesw()
	//	if !reflect.DeepEqual(tt.err, errstring(err)) {
	//		t.Errorf("%d. %q:\ error mismatch:\n  exp=%s\n  zt=%s\n\n", i, tt.s, tt.err, err)
	//	} else if tt.err == "" && !reflect.DeepEqual(tt.stmt, stmt) {
	//		t.Errorf("%d. %q\n\nstmt mismatch:\n\nexp=%#v\n\nzt=%#v\n\n", i, tt.s, tt.stmt, stmt)
	//	}
	//}
	//for i, tt := range tests {
	//	stmt, err := andrea.NewParser(strings.NewReader(tt.s)).ParseSum()
	//	if !reflect.DeepEqual(tt.err, errstring(err)) {
	//		t.Errorf("%d. %q: error mismatch:\n  exp=%s\n  zt=%s\n\n", i, tt.s, tt.err, err)
	//	} else if tt.err == "" && !reflect.DeepEqual(tt.stmt, stmt) {
	//		t.Errorf("%d. %q\n\nstmt mismatch:\n\nexp=%#v\n\nzt=%#v\n\n", i, tt.s, tt.stmt, stmt)
	//	}
	//}
	//for i, tt := range tests {
	//	stmt, err := andrea.NewParser(strings.NewReader(tt.s)).ParseSumNom()
	//	if !reflect.DeepEqual(tt.err, errstring(err)) {
	//		t.Errorf("%d. %q: error mismatch:\n  exp=%s\n  zt=%s\n\n", i, tt.s, tt.err, err)
	//	} else if tt.err == "" && !reflect.DeepEqual(tt.stmt, stmt) {
	//		t.Errorf("%d. %q\n\nstmt mismatch:\n\nexp=%#v\n\nzt=%#v\n\n", i, tt.s, tt.stmt, stmt)
	//	}
	//}
	//for i, tt := range tests {
	//	stmt, err := andrea.NewParser(strings.NewReader(tt.s)).Parsepara()
	//	if !reflect.DeepEqual(tt.err, errstring(err)) {
	//		t.Errorf("%d. %q: error mismatch:\n  exp=%s\n  zt=%s\n\n", i, tt.s, tt.err, err)
	//	} else if tt.err == "" && !reflect.DeepEqual(tt.stmt, stmt) {
	//		t.Errorf("%d. %q\n\nstmt mismatch:\n\nexp=%#v\n\nzt=%#v\n\n", i, tt.s, tt.stmt, stmt)
	//	}
	//}
	//for i, tt := range tests {
	//	stmt, err := andrea.NewParser(strings.NewReader(tt.s)).Parsezcase()
	//	if !reflect.DeepEqual(tt.err, errstring(err)) {
	//		t.Errorf("%d. %q: error mismatch:\n  exp=%s\n  zt=%s\n\n", i, tt.s, tt.err, err)
	//	} else if tt.err == "" && !reflect.DeepEqual(tt.stmt, stmt) {
	//		t.Errorf("%d. %q\n\nstmt mismatch:\n\nexp=%#v\n\nzt=%#v\n\n", i, tt.s, tt.stmt, stmt)
	//	}
	//}
	//

	//for i, tt := range tests1 {
	//	stmt, err := andrea.NewParser(strings.NewReader(tt.j)).Parsesi()
	//	if !reflect.DeepEqual(tt.err, errstring(err)) {
	//		t.Errorf("%d. %q: error mismatch:\n  exp=%s\n  zt=%s\n\n", i, tt.s, tt.err, err)
	//	} else if tt.err == "" && !reflect.DeepEqual(tt.stmt, stmt) {
	//		t.Errorf("%d. %q\n\nstmt mismatch:\n\nexp=%#v\n\nzt=%#v\n\n", i, tt.s, tt.stmt, stmt)
	//	}
	//}
}

// errstring returns the string representation of an error.
func errstring(err error) string {
	if err != nil {
		return err.Error()
	}
	return ""
}
