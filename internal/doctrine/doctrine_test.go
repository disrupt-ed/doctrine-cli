package doctrine

import (
	"slices"
	"strings"
	"testing"
	"testing/fstest"

	"github.com/disrupt-ed/doctrine-cli/doctrines"
)

func names(ds []Doctrine) []string {
	var out []string
	for _, d := range ds {
		out = append(out, d.Name)
	}
	return out
}

func TestResolve(t *testing.T) {
	tests := []struct {
		extends, want []string
	}{
		{[]string{"rails/default"}, []string{"engineering/default", "ruby/default", "rails/default"}},
		{[]string{"go/default"}, []string{"engineering/default", "go/default"}},
		{[]string{"rails/default", "go/default"}, []string{"engineering/default", "ruby/default", "rails/default", "go/default"}},
		{[]string{"go/default", "rails/default"}, []string{"engineering/default", "go/default", "ruby/default", "rails/default"}},
	}
	for _, tt := range tests {
		got, err := Resolve(doctrines.FS, tt.extends)
		if err != nil {
			t.Fatal(err)
		}
		if !slices.Equal(names(got), tt.want) {
			t.Errorf("Resolve(%v) = %v, want %v", tt.extends, names(got), tt.want)
		}
	}
}

func TestResolveUnknown(t *testing.T) {
	if _, err := Resolve(doctrines.FS, []string{"cobol/default"}); err == nil || !strings.Contains(err.Error(), "unknown doctrine") {
		t.Fatalf("err = %v", err)
	}
}

func TestResolveCycle(t *testing.T) {
	fsys := fstest.MapFS{
		"a/x/doctrine.yml": {Data: []byte("name: a/x\nparent: b/x\n")},
		"b/x/doctrine.yml": {Data: []byte("name: b/x\nparent: a/x\n")},
	}
	if _, err := Resolve(fsys, []string{"a/x"}); err == nil || !strings.Contains(err.Error(), "cycle") {
		t.Fatalf("err = %v", err)
	}
}

func TestLoadJoinsFiles(t *testing.T) {
	fsys := fstest.MapFS{
		"a/x/doctrine.yml": {Data: []byte("name: a/x\nfiles: [one.md, two.md]\n")},
		"a/x/one.md":       {Data: []byte("# One\n")},
		"a/x/two.md":       {Data: []byte("# Two\n")},
	}
	d, err := Load(fsys, "a/x")
	if err != nil {
		t.Fatal(err)
	}
	if d.Body != "# One\n\n# Two\n" {
		t.Fatalf("Body = %q", d.Body)
	}
	if d.Topic() != "a" {
		t.Fatalf("Topic = %q", d.Topic())
	}
}

func TestTopLevel(t *testing.T) {
	tests := []struct{ in, want []string }{
		{[]string{"ruby/default", "rails/default"}, []string{"rails/default"}},
		{[]string{"rails/default", "go/default", "go/default"}, []string{"rails/default", "go/default"}},
		{[]string{"engineering/default", "go/default"}, []string{"go/default"}},
		{[]string{"ruby/default"}, []string{"ruby/default"}},
	}
	for _, tt := range tests {
		got, err := TopLevel(doctrines.FS, tt.in)
		if err != nil {
			t.Fatal(err)
		}
		if !slices.Equal(got, tt.want) {
			t.Errorf("TopLevel(%v) = %v, want %v", tt.in, got, tt.want)
		}
	}
}

func TestAllOfficialDoctrinesLoad(t *testing.T) {
	all, err := All(doctrines.FS)
	if err != nil {
		t.Fatal(err)
	}
	want := []string{"engineering/default", "go/default", "rails/default", "ruby/default"}
	if !slices.Equal(names(all), want) {
		t.Fatalf("All = %v, want %v", names(all), want)
	}
	for _, d := range all {
		if d.Description == "" || len(d.Files) == 0 || d.Body == "" {
			t.Errorf("%s: missing description, files or content", d.Name)
		}
	}
}
