package file

import (
	"errors"
	"io"
	"os"
)

type Provider struct{}

func NewProvider() *Provider {
	return &Provider{}
}

func (provider *Provider) IsVersionValid() (*string, bool, error) {
	return nil, false, errors.New("should not check the version")
}

func (provider *Provider) Protocol() string {
	return PROTOCOL
}

func (provider *Provider) Read(session *string, name string) (updatedSession *string, content *string, err error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, nil, err
	}

	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, nil, err
	}

	asString := string(bytes)
	content = &asString

	return nil, content, nil
}

func (provider *Provider) ShouldCheckVersion() bool {
	return false
}
