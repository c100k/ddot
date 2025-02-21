package bwcli

import (
	"bytes"
	"c100k/ddot/providers"
	"errors"
	"os/exec"
	"strings"
)

func fmtErr(stderr *bytes.Buffer) error {
	full := strings.TrimSpace(stderr.String())

	// Sanitize the error as it contains things we don't want. Typically, an error message looks like this :
	// 		(node:15682) [DEP0040] DeprecationWarning: The `punycode` module is deprecated. Please use a userland alternative instead.
	// 		(Use `node --trace-deprecation ...` to show where the warning was created)
	// 		Invalid master password.
	// Only the last part is interesting. This is flacky and can change at anyone of there release so keep an eye on it.
	lines := strings.Split(full, "\n")
	var linesToKeep []string
	for _, line := range lines {
		if strings.Contains(line, "punycode") || strings.Contains(line, "deprecation") {
			continue
		}
		linesToKeep = append(linesToKeep, line)
	}
	msg := strings.Join(linesToKeep, "\n")

	// Voluntarily using a "broader" string avoid breaking because of this.
	lowercase := strings.ToLower(msg)
	if strings.Contains(lowercase, "not logged in") {
		// With the current version, the actual message is "You are not logged in."
		return providers.ErrNotLoggedIn
	} else if strings.Contains(lowercase, "is locked") {
		// With the current version, the actual message is "Vault is locked."
		return providers.ErrLocked
	}

	return errors.New(msg)
}

func fmtOut(stdout *bytes.Buffer) string {
	return strings.TrimSpace(stdout.String())
}

func run(args ...string) (*bytes.Buffer, *bytes.Buffer, error) {
	cmd := exec.Command(BIN_CMD, args...)

	// In case a command does not return a result as expected, uncomment this to make sure the CLI does not prompt
	// cmd.Stdin = os.Stdin

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()

	return &stdout, &stderr, err
}
