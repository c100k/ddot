package providers

type Provider interface {
	IsVersionValid() (*string, bool, error)
	Protocol() string
	Read(session *string, name string) (updatedSession *string, content *string, err error)
	ShouldCheckVersion() bool
}
