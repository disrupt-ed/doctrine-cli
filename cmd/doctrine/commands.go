package main

import (
	"bufio"
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/disrupt-ed/doctrine-cli/doctrines"
	"github.com/disrupt-ed/doctrine-cli/internal/adapters/claude"
	"github.com/disrupt-ed/doctrine-cli/internal/compose"
	"github.com/disrupt-ed/doctrine-cli/internal/dependencies"
	"github.com/disrupt-ed/doctrine-cli/internal/detect"
	"github.com/disrupt-ed/doctrine-cli/internal/doctrine"
	"github.com/disrupt-ed/doctrine-cli/internal/generate"
	"github.com/disrupt-ed/doctrine-cli/internal/manifest"
)

func detectCmd(root string, out io.Writer) error {
	r, err := detect.Detect(root)
	if err != nil {
		return err
	}
	fmt.Fprintf(out, "Technologies: %s\n", orNone(r.Technologies))
	fmt.Fprintf(out, "Doctrines:    %s\n", orNone(r.Doctrines))
	fmt.Fprintf(out, "Agents:       %s\n", orNone(r.Agents))
	return nil
}

func initCmd(ctx context.Context, root string, args []string, in io.Reader, out io.Writer) error {
	fs := flag.NewFlagSet("init", flag.ContinueOnError)
	fs.SetOutput(out)
	if err := fs.Parse(args); err != nil {
		return err
	}
	if _, err := os.Stat(filepath.Join(root, manifest.Path)); err == nil {
		return fmt.Errorf("%s already exists. Use doctrine add, remove or generate", manifest.Path)
	}

	r, err := detect.Detect(root)
	if err != nil {
		return err
	}
	extends := r.Doctrines
	if fs.NArg() > 0 {
		extends = nil
		for _, s := range fs.Args() {
			extends = append(extends, detect.DoctrineFor(s))
		}
	}
	if len(extends) == 0 {
		return fmt.Errorf("no Ruby, Rails or Go found. Name the stacks: doctrine init rails go")
	}
	extends, err = doctrine.TopLevel(doctrines.FS, extends)
	if err != nil {
		return err
	}

	m := manifest.Manifest{Version: 1, Doctrines: manifest.Release(doctrines.Release), Extends: extends}
	if err := writeManifest(root, m.Marshal()); err != nil {
		return err
	}
	fmt.Fprintf(out, "Wrote %s (%s)\n", manifest.Path, strings.Join(extends, ", "))
	if err := regenerate(ctx, root, out); err != nil {
		return err
	}
	askStats(root, r, in, out)
	return nil
}

func generateCmd(ctx context.Context, root string, args []string, out io.Writer) error {
	for _, a := range args {
		if a != "claude" {
			return fmt.Errorf("unknown agent %q (only claude is supported)", a)
		}
	}
	return regenerate(ctx, root, out)
}

func inspectCmd(ctx context.Context, root string, args []string, out io.Writer) error {
	if len(args) > 0 {
		var names []string
		for _, a := range args {
			names = append(names, detect.DoctrineFor(a))
		}
		for i, name := range names {
			d, err := doctrine.Load(doctrines.FS, name)
			if err != nil {
				return err
			}
			if i > 0 {
				fmt.Fprintln(out)
			}
			fmt.Fprint(out, d.Body)
		}
		return nil
	}
	e, _, err := project(ctx, root)
	if err != nil {
		return err
	}
	fmt.Fprint(out, e.Markdown())
	return nil
}

func listCmd(out io.Writer) error {
	all, err := doctrine.All(doctrines.FS)
	if err != nil {
		return err
	}
	fmt.Fprintf(out, "Official doctrines, release %s:\n\n", doctrines.Release)
	for _, d := range all {
		fmt.Fprintf(out, "  %-22s %s\n", d.Name, d.Description)
	}
	return nil
}

func addCmd(ctx context.Context, root string, args []string, out io.Writer) error {
	return changeExtends(ctx, root, args, out, func(extends []string, name string) []string {
		return append(extends, name)
	})
}

func removeCmd(ctx context.Context, root string, args []string, out io.Writer) error {
	return changeExtends(ctx, root, args, out, func(extends []string, name string) []string {
		return slices.DeleteFunc(extends, func(e string) bool { return e == name })
	})
}

func changeExtends(ctx context.Context, root string, args []string, out io.Writer, change func([]string, string) []string) error {
	if len(args) == 0 {
		return fmt.Errorf("name at least one stack, like rails or go")
	}
	m, data, err := manifest.Read(root)
	if err != nil {
		return notInitialized(err)
	}
	release, err := doctrine.Release(ctx, string(m.Doctrines))
	if err != nil {
		return err
	}
	extends := slices.Clone(m.Extends)
	for _, a := range args {
		name := detect.DoctrineFor(a)
		if _, err := doctrine.Load(release, name); err != nil {
			return err
		}
		extends = change(extends, name)
	}
	if extends, err = doctrine.TopLevel(release, extends); err != nil {
		return err
	}
	if len(extends) == 0 {
		return fmt.Errorf("a project needs at least one doctrine")
	}
	data, err = manifest.SetExtends(data, extends)
	if err != nil {
		return err
	}
	if err := writeManifest(root, data); err != nil {
		return err
	}
	fmt.Fprintf(out, "Doctrines: %s\n", strings.Join(extends, ", "))
	return regenerate(ctx, root, out)
}

func updateCmd(ctx context.Context, root string, args []string, in io.Reader, out io.Writer) error {
	fs := flag.NewFlagSet("update", flag.ContinueOnError)
	fs.SetOutput(out)
	yes := fs.Bool("yes", false, "accept without asking")
	if err := fs.Parse(args); err != nil {
		return err
	}
	m, data, err := manifest.Read(root)
	if err != nil {
		return notInitialized(err)
	}
	if string(m.Doctrines) == doctrines.Release {
		fmt.Fprintf(out, "Already on doctrine release %s.\n", doctrines.Release)
		return nil
	}
	oldRelease, err := doctrine.Release(ctx, string(m.Doctrines))
	if err != nil {
		return err
	}
	before, err := compose.Compose(oldRelease, m, root)
	if err != nil {
		return err
	}
	after, err := compose.Compose(doctrines.FS, m, root)
	if err != nil {
		return err
	}

	fmt.Fprintf(out, "Doctrine release %s → %s. What changes for your agent:\n\n", m.Doctrines, doctrines.Release)
	oldFiles, newFiles := claude.Files(before), claude.Files(after)
	var paths []string
	for p := range oldFiles {
		paths = append(paths, p)
	}
	for p := range newFiles {
		if _, ok := oldFiles[p]; !ok {
			paths = append(paths, p)
		}
	}
	slices.Sort(paths)
	changed := false
	for _, p := range paths {
		if d := diff(p, string(oldFiles[p]), string(newFiles[p])); d != "" {
			fmt.Fprint(out, d)
			changed = true
		}
	}
	if !changed {
		fmt.Fprintln(out, "Nothing changes for your agent.")
	}

	if !*yes && !confirm(in, out, "\nAccept and update?") {
		fmt.Fprintln(out, "Nothing changed.")
		return nil
	}
	data, err = manifest.SetRelease(data, doctrines.Release)
	if err != nil {
		return err
	}
	if err := writeManifest(root, data); err != nil {
		return err
	}
	return regenerate(ctx, root, out)
}

func dependencyCmd(ctx context.Context, args []string, out io.Writer) error {
	if len(args) < 1 || args[0] != "inspect" {
		return fmt.Errorf("usage: doctrine dependency inspect NAME")
	}
	fs := flag.NewFlagSet("dependency inspect", flag.ContinueOnError)
	fs.SetOutput(out)
	eco := fs.String("ecosystem", "", "rubygems or go (guessed from the name if empty)")
	if err := fs.Parse(args[1:]); err != nil {
		return err
	}
	if fs.NArg() != 1 {
		return fmt.Errorf("usage: doctrine dependency inspect [-ecosystem rubygems|go] NAME")
	}
	name := fs.Arg(0)
	e := dependencies.Ecosystem(*eco)
	if e == "" {
		e = dependencies.Guess(name)
	}
	if e != dependencies.RubyGems && e != dependencies.Go {
		return fmt.Errorf("unknown ecosystem %q", *eco)
	}
	c := &dependencies.Client{GitHubToken: dependencies.TokenFromEnv()}
	r, err := c.Inspect(ctx, e, name)
	if err != nil {
		return err
	}
	fmt.Fprint(out, r.String())
	return nil
}

// project reads the manifest and composes the effective doctrine.
func project(ctx context.Context, root string) (compose.Effective, manifest.Manifest, error) {
	m, _, err := manifest.Read(root)
	if err != nil {
		return compose.Effective{}, m, notInitialized(err)
	}
	release, err := doctrine.Release(ctx, string(m.Doctrines))
	if err != nil {
		return compose.Effective{}, m, err
	}
	e, err := compose.Compose(release, m, root)
	return e, m, err
}

func regenerate(ctx context.Context, root string, out io.Writer) error {
	e, _, err := project(ctx, root)
	if err != nil {
		return err
	}
	written, err := generate.Write(root, claude.Files(e), claude.Owned)
	if err != nil {
		return err
	}
	fmt.Fprintln(out, "Generated Claude Code configuration:")
	for _, w := range written {
		fmt.Fprintf(out, "  %s\n", w)
	}
	return nil
}

func writeManifest(root string, data []byte) error {
	path := filepath.Join(root, manifest.Path)
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	return os.WriteFile(path, data, 0o644)
}

func notInitialized(err error) error {
	if errors.Is(err, fs.ErrNotExist) {
		return fmt.Errorf("no %s here. Run doctrine init first", manifest.Path)
	}
	return err
}

func confirm(in io.Reader, out io.Writer, question string) bool {
	fmt.Fprintf(out, "%s [y/N] ", question)
	line, _ := bufio.NewReader(in).ReadString('\n')
	answer := strings.ToLower(strings.TrimSpace(line))
	return answer == "y" || answer == "yes"
}

func orNone(s []string) string {
	if len(s) == 0 {
		return "none"
	}
	return strings.Join(s, ", ")
}
