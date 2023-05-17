package json

import (
	"testing"
	"time"
)

func TestMarshalPatch(t *testing.T) {
	type S struct{
		T int
		U string
	}
	v := struct {
		A [2]float64
		B time.Time
		C S

		X [2]float64 `json:",omitempty"`
		Y time.Time `json:",omitempty"`
		Z S `json:",omitempty"`
	}{}
	want := `{"A":[0,0],"B":"0001-01-01T00:00:00Z","C":{"T":0,"U":""}}`
	have, _ := Marshal(v)
	if h, w := string(have), want; h != w {
		t.Fatalf("have %s, want %s", h, w)
	}
}
