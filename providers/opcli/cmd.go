package opcli

import (
	"encoding/json"
	"errors"
	"fmt"
)

func assertLoggedIn() error {
	stdout, stderr, err := run("account", "list", "--format", "json")
	if err != nil {
		return fmtErr(stderr)
	}

	var items []Account
	json.Unmarshal(stdout.Bytes(), &items)
	count := len(items)
	if count == 0 {
		return errors.New("no authenticated account found : authenticate at least one")
	}

	return nil
}

func assertUnlocked() error {
	_, stderr, err := run("signin", "--raw")
	if err != nil {
		return fmtErr(stderr)
	}

	return nil
}

func get(_ string, name string) (content *string, err error) {
	stdout, stderr, err := run("item", "get", name, "--format", "json")
	if err != nil {
		return nil, fmtErr(stderr)
	}

	var item Note
	err = json.Unmarshal(stdout.Bytes(), &item)
	if err != nil {
		return nil, err
	}

	var contentField *NoteField
	for _, field := range item.Fields {
		if field.Id == "notesPlain" {
			contentField = &field
			break
		}
	}
	if contentField == nil {
		return nil, fmt.Errorf("note %s does not contain a field with id 'notesPlain''", item.Id)
	}

	content = &contentField.Value

	return content, nil
}

func version() (*string, error) {
	stdout, stderr, err := run("--version")
	if err != nil {
		return nil, fmtErr(stderr)
	}

	version := fmtOut(stdout)

	return &version, nil
}
