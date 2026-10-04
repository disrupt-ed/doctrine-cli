// Command doctrine sets up Agent Doctrine in a repository.
package main

import (
	"context"
	"fmt"
	"io"
	"os"
)

// version is set at build time with -ldflags "-X main.version=...".
var version = "dev"

const usage = `doctrine teaches your coding agent how you build software.

Usage:
  doctrine detect                    show detected technologies and suggested doctrines
  doctrine init [stacks...]          set up Doctrine in this repository
  doctrine install HANDLE            install a configuration made on doctrine.codedynamic.com
  doctrine generate [claude]         regenerate agent configuration
  doctrine inspect [doctrine...]     print the effective Doctrine, or the named doctrines
  doctrine list                      list official doctrines
  doctrine add STACK...              add doctrines and regenerate
  doctrine remove STACK...           remove doctrines and regenerate
  doctrine update [-yes]             review and accept a newer doctrine release
  doctrine dependency inspect NAME   report facts about a gem or Go module
  doctrine version                   print the version

Stacks: ruby, rails, go, or a doctrine name like rails/default.
`

func main() {
	if err := run(context.Background(), os.Args[1:], os.Stdin, os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, "doctrine:", err)
		os.Exit(1)
	}
}

func run(ctx context.Context, args []string, stdin io.Reader, stdout io.Writer) error {
	if len(args) == 0 {
		fmt.Fprint(stdout, usage)
		return nil
	}
	root, err := os.Getwd()
	if err != nil {
		return err
	}
	cmd, args := args[0], args[1:]
	switch cmd {
	case "detect":
		return detectCmd(root, stdout)
	case "init":
		return initCmd(ctx, root, args, stdin, stdout)
	case "install":
		return installCmd(ctx, root, args, stdin, stdout)
	case "generate":
		return generateCmd(ctx, root, args, stdout)
	case "inspect":
		return inspectCmd(ctx, root, args, stdout)
	case "list":
		return listCmd(stdout)
	case "add":
		return addCmd(ctx, root, args, stdout)
	case "remove":
		return removeCmd(ctx, root, args, stdout)
	case "update":
		return updateCmd(ctx, root, args, stdin, stdout)
	case "dependency":
		return dependencyCmd(ctx, args, stdout)
	case "version", "-v", "--version":
		fmt.Fprintln(stdout, "doctrine", version)
		return nil
	case "help", "-h", "--help":
		fmt.Fprint(stdout, usage)
		return nil
	default:
		return fmt.Errorf("unknown command %q (run doctrine help)", cmd)
	}
}
