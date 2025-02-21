package opcli

import (
	"c100k/ddot/ui"
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

func (provider *Provider) ShouldCheckVersion() bool {
	return true
}

func (provider *Provider) Read(session *string, name string) (updatedSession *string, content *string, err error) {
	err = assertLoggedIn()
	if err != nil {
		return nil, nil, err
	}

	err = assertUnlocked()
	if err != nil {
		return nil, nil, err
	}

	// The session management is automatically done via the app and Touch ID
	// So it works on macOS for now and not tested for other platforms.
	// Unlike bw, op does not allow to pass the password as a CLI argument to unlock.
	// So we cannot prompt the user for it
	// And since we don't want to enable os.Stdin in `run` to stay generic and prompt ourselves,
	// in order to stay "generic", we rely on the easiest solution with op (for now).
	// TODO : Handle op unlock without Touch ID
	stubSession := ""
	updatedSession = &stubSession

	ui.Print("🔍", fmt.Sprintf("Searching for note '%s' on %s", name, PROTOCOL))
	content, err = get(*updatedSession, name)
	if err != nil {
		return nil, nil, err
	}

	return updatedSession, content, nil
}
