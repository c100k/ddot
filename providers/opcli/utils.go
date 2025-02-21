package opcli

import (
	"bytes"
	"errors"
	"os/exec"
	"strings"
)

func fmtErr(stderr *bytes.Buffer) error {
	// [ERROR] 2025/01/04 14:30:19 "Netflix" isn't an item. Specify the item with its UUID, name, or domain.
	// [ERROR] 2025/01/04 14:41:14 (401) Unauthorized: You aren't authorized to perform this action.
	return errors.New(strings.TrimSpace(stderr.String()))
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
