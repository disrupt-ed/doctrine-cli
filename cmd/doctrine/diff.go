package main

import (
	"fmt"
	"strings"
)

// diff returns a simple line diff of one file, or "" when nothing changed.
// Unchanged lines are left out except for one line of context around changes.
func diff(path, before, after string) string {
	if before == after {
		return ""
	}
	a, b := lines(before), lines(after)

	// Longest common subsequence table.
	lcs := make([][]int, len(a)+1)
	for i := range lcs {
		lcs[i] = make([]int, len(b)+1)
	}
	for i := len(a) - 1; i >= 0; i-- {
		for j := len(b) - 1; j >= 0; j-- {
			if a[i] == b[j] {
				lcs[i][j] = lcs[i+1][j+1] + 1
			} else {
				lcs[i][j] = max(lcs[i+1][j], lcs[i][j+1])
			}
		}
	}

	type line struct {
		op   byte
		text string
	}
	var ops []line
	i, j := 0, 0
	for i < len(a) || j < len(b) {
		switch {
		case i < len(a) && j < len(b) && a[i] == b[j]:
			ops = append(ops, line{' ', a[i]})
			i++
			j++
		case i < len(a) && (j == len(b) || lcs[i+1][j] >= lcs[i][j+1]):
			ops = append(ops, line{'-', a[i]})
			i++
		default:
			ops = append(ops, line{'+', b[j]})
			j++
		}
	}

	var out strings.Builder
	fmt.Fprintf(&out, "--- %s\n+++ %s\n", path, path)
	skipped := false
	for k, l := range ops {
		near := l.op != ' ' ||
			(k > 0 && ops[k-1].op != ' ') ||
			(k+1 < len(ops) && ops[k+1].op != ' ')
		if !near {
			skipped = true
			continue
		}
		if skipped {
			out.WriteString("@@\n")
			skipped = false
		}
		fmt.Fprintf(&out, "%c%s\n", l.op, l.text)
	}
	out.WriteString("\n")
	return out.String()
}

func lines(s string) []string {
	if s == "" {
		return nil
	}
	return strings.Split(strings.TrimSuffix(s, "\n"), "\n")
}
