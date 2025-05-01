package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"

	"github.com/Kam1k4dze/wirez/command"
)

// these still get overridden by -ldflags if you pass them
var (
	version = "dev"
	commit  = ""
)

func main() {
	command.Main(buildVersion(version, commit))
}

func buildVersion(v, c string) string {
	// if no override, fill in from Git
	if v == "dev" {
		if desc := gitOut("describe", "--tags", "--always"); desc != "" {
			v = desc
		}

		// get hash
		hash := gitOut("rev-parse", "--short", "HEAD")
		// get message
		msg := gitOut("log", "-1", "--pretty=%s")
		// get date in YYYY-MM-DD
		date := gitOut("log", "-1", "--date=format:%Y-%m-%d", "--pretty=%cd")

		if hash != "" {
			// build the commit string with hash, message, date
			parts := []string{hash}
			if msg != "" {
				parts = append(parts, msg)
			}
			if date != "" {
				parts = append(parts, date)
			}
			// e.g. "5316463 (bug fixes, 2025-04-30)"
			c = fmt.Sprintf("%s (%s)", parts[0], strings.Join(parts[1:], ", "))
		}
	}

	result := v
	if c != "" {
		result = fmt.Sprintf("%s\ncommit: %s", v, c)
	}
	return result
}

// gitOut runs git <args> and returns its trimmed stdout, or "" on error.
func gitOut(args ...string) string {
	cmd := exec.Command("git", args...)
	var out bytes.Buffer
	cmd.Stdout = &out
	_ = cmd.Run() // ignore errors
	return strings.TrimSpace(out.String())
}
