// Package doctrines embeds the official doctrines shipped with the CLI.
package doctrines

import "embed"

// Release is the official doctrine release embedded in this binary.
const Release = "0.1"

// FS holds the official doctrines, one folder per doctrine (rails/default, ...).
//
//go:embed */default
var FS embed.FS
