package bwcli

import (
	"c100k/ddot/ui"
	"encoding/json"
	"fmt"
	"strings"

	"golang.org/x/term"
)

func assertLoggedIn() error {
	_, stderr, err := run("login", "--check")
	if err != nil {
		return fmtErr(stderr)
	}

	return nil
}

func assertUnlocked() error {
	_, stderr, err := run("unlock", "--check")
	if err != nil {
		return fmtErr(stderr)
	}

	return nil
}

func get(session string, name string) (content *string, err error) {
	stdout, stderr, err := run("get", "item", "--session", session, name)
	if err != nil {
		return nil, fmtErr(stderr)
	}

	var item Note
	err = json.Unmarshal(stdout.Bytes(), &item)
	if err != nil {
		return nil, err
	}

	content = &item.Notes

	return content, nil
}

func login(email *string, password *string) error {
	_, stderr, err := run("login", *email, *password)
	if err != nil {
		return fmtErr(stderr)
	}

	return nil
}

func promptEmail() (*string, error) {
	var email string
	ui.Prompt("💌", "bw email")
	_, err := fmt.Scanln(&email)
	if err != nil {
		return nil, err
	}

	return &email, nil
}

func promptPassword() (*string, error) {
	ui.Prompt("🔐", "bw password")
	bytes, err := term.ReadPassword(0)
	if err != nil {
		return nil, err
	}

	password := string(bytes)
	return &password, nil
}

func unlock(password *string) (session *string, err error) {
	stdout, stderr, err := run("unlock", *password)
	if err != nil {
		return nil, fmtErr(stderr)
	}

	out := fmtOut(stdout)
	lines := strings.Split(out, "\n")
	if len(lines) == 0 {
		return nil, fmt.Errorf("unexpected empty output : %s", out)
	}
	last := lines[len(lines)-1]
	if last[0] != '$' {
		return nil, fmt.Errorf("expected command with --session as last line (got : %s)", last)
	}
	args := strings.Split(last, " ")
	if len(args) == 0 {
		return nil, fmt.Errorf("unexpected empty last line : %s", last)
	}
	session = &args[len(args)-1]

	return session, nil
}

func version() (*string, error) {
	stdout, stderr, err := run("--version")
	if err != nil {
		return nil, fmtErr(stderr)
	}

	version := fmtOut(stdout)

	return &version, nil
}
