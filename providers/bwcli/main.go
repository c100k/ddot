package bwcli

import (
	"c100k/ddot/providers"
	"c100k/ddot/ui"
	"errors"
	"fmt"
)

type Provider struct{}

func NewProvider() *Provider {
	return &Provider{}
}

func (provider *Provider) IsVersionValid() (*string, bool, error) {
	version, err := version()
	if err != nil {
		return version, false, err
	}

	// Ready to welcome the code checking the version if/when need be
	isValid := true

	return version, isValid, err
}

func (provider *Provider) Protocol() string {
	return PROTOCOL
}

func (provider *Provider) Read(session *string, name string) (updatedSession *string, content *string, err error) {
	err = assertLoggedIn()
	if err != nil {
		if !errors.Is(err, providers.ErrNotLoggedIn) {
			return nil, nil, err
		}

		email, err := promptEmail()
		if err != nil {
			return nil, nil, err
		}

		password, err := promptPassword()
		if err != nil {
			return nil, nil, err
		}

		err = login(email, password)
		if err != nil {
			return nil, nil, err
		}
	}

	verb := "Got"
	if session == nil {
		err = assertUnlocked()
		if err != nil {
			if !errors.Is(err, providers.ErrLocked) {
				return nil, nil, err
			}

			password, err := promptPassword()
			if err != nil {
				return nil, nil, err
			}

			ui.Print("", "")
			updatedSession, err = unlock(password)
			if err != nil {
				return nil, nil, err
			}
		}
	} else {
		updatedSession = session
		verb = "Reusing"
	}

	obfuscated := (*updatedSession)[:10] + "******"
	ui.Print("|_ 🔑", fmt.Sprintf("%s session : %s", verb, obfuscated))

	ui.Print("🔍", fmt.Sprintf("Searching for note '%s' on %s", name, PROTOCOL))
	content, err = get(*updatedSession, name)
	if err != nil {
		return nil, nil, err
	}

	return updatedSession, content, nil
}

func (provider *Provider) ShouldCheckVersion() bool {
	return true
}
