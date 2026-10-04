package manifest

import (
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name, input, err string
	}{
		{"valid", "version: 1\ndoctrines: 0.1\nextends: [rails/default]\n", ""},
		{"preferences", "version: 1\ndoctrines: 0.1\nextends: [go/default]\nengineering:\n  dependencies: open\n", ""},
		{"wrong version", "version: 2\ndoctrines: 0.1\nextends: [go/default]\n", "unsupported version"},
		{"no release", "version: 1\nextends: [go/default]\n", "release is missing"},
		{"no extends", "version: 1\ndoctrines: 0.1\n", "extends is empty"},
		{"unknown preference", "version: 1\ndoctrines: 0.1\nextends: [go/default]\nengineering:\n  tabs: always\n", "unknown engineering preference"},
		{"bad value", "version: 1\ndoctrines: 0.1\nextends: [go/default]\nengineering:\n  kiss: extreme\n", "must be one of strong, normal"},
		{"not yaml", "version: [", "parse manifest"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Parse([]byte(tt.input))
			if tt.err == "" && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if tt.err != "" && (err == nil || !strings.Contains(err.Error(), tt.err)) {
				t.Fatalf("error = %v, want %q", err, tt.err)
			}
		})
	}
}

func TestReleaseKeepsText(t *testing.T) {
	m, err := Parse([]byte("version: 1\ndoctrines: 0.10\nextends: [go/default]\n"))
	if err != nil {
		t.Fatal(err)
	}
	if m.Doctrines != "0.10" {
		t.Fatalf("Doctrines = %q, want 0.10", m.Doctrines)
	}
}

func TestMarshal(t *testing.T) {
	m := Manifest{Version: 1, Doctrines: "0.1", Extends: []string{"rails/default", "go/default"},
		Engineering: map[string]string{"dependencies": "open", "kiss": "strong"}}
	want := "version: 1\ndoctrines: 0.1\n\nextends:\n  - rails/default\n  - go/default\n\nengineering:\n  kiss: strong\n  dependencies: open\n"
	if got := string(m.Marshal()); got != want {
		t.Fatalf("got:\n%s\nwant:\n%s", got, want)
	}
	if _, err := Parse(m.Marshal()); err != nil {
		t.Fatalf("marshaled manifest doesn't parse: %v", err)
	}
}

func TestSetExtendsKeepsComments(t *testing.T) {
	in := "version: 1\ndoctrines: 0.1 # pinned\n\n# what we follow\nextends:\n  - rails/default\n"
	out, err := SetExtends([]byte(in), []string{"rails/default", "go/default"})
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{"# pinned", "# what we follow", "- go/default"} {
		if !strings.Contains(string(out), want) {
			t.Errorf("missing %q in:\n%s", want, out)
		}
	}
	m, err := Parse(out)
	if err != nil {
		t.Fatal(err)
	}
	if len(m.Extends) != 2 {
		t.Fatalf("Extends = %v", m.Extends)
	}
}

func TestSetRelease(t *testing.T) {
	out, err := SetRelease([]byte("version: 1\ndoctrines: 0.1\nextends: [go/default]\n"), "0.2")
	if err != nil {
		t.Fatal(err)
	}
	m, err := Parse(out)
	if err != nil {
		t.Fatal(err)
	}
	if m.Doctrines != "0.2" {
		t.Fatalf("Doctrines = %q", m.Doctrines)
	}
}
