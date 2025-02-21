package rsrc

import (
	"c100k/ddot/providers"
	"c100k/ddot/providers/bwcli"
	"c100k/ddot/providers/file"
	"c100k/ddot/providers/opcli"
	"c100k/ddot/ui"
	"fmt"
)

func ReadAll(uris []string) (resources []*Resource, err error) {
	count := len(uris)
	if count == 0 {
		return nil, fmt.Errorf("pass at least one resource")
	}

	ui.Print("📜", fmt.Sprintf("Processing %d resource(s) : %v", count, uris))

	// NOTE : The most expensive thing is to search for the resource in the provider.
	// Therefore, in the some cases, we can search N-1 resources and have the Nth one fail (e.g. not found).
	// In this case, all the operations made before are useless.
	// That's why if an URI is not valid, or some versions are not valid, it's better to fail early.

	resources, err = parseURIs(uris)
	if err != nil {
		return nil, err
	}

	err = assertProvidersVersions(resources)
	if err != nil {
		return nil, err
	}

	sessions := make(map[string]*string)

	for _, resource := range resources {
		provider, err := initProvider(resource)
		if err != nil {
			return nil, err
		}

		updatedSession, content, err := provider.Read(sessions[resource.Protocol], resource.Path)
		if err != nil {
			return nil, err
		}

		sessions[resource.Protocol] = updatedSession
		resource.Content = content
	}

	return resources, nil
}

func assertProvidersVersions(resources []*Resource) error {
	providers := make(map[string]providers.Provider)
	for _, resource := range resources {
		if _, exists := providers[resource.Protocol]; exists {
			continue
		}
		provider, err := initProvider(resource)
		if err != nil {
			return err
		}
		providers[resource.Protocol] = provider
	}

	for _, provider := range providers {
		if !provider.ShouldCheckVersion() {
			continue
		}

		ui.Print("#️⃣", fmt.Sprintf("Checking %s version", provider.Protocol()))
		version, isValid, err := provider.IsVersionValid()
		if err != nil {
			return err
		}
		if !isValid {
			return fmt.Errorf("nope ! %s version is invalid : %s", provider.Protocol(), *version)
		}

		ui.Print("|_ 👍", fmt.Sprintf("Yes ! %s version %s is valid", provider.Protocol(), *version))
	}

	return nil
}

func initProvider(resource *Resource) (providers.Provider, error) {
	switch resource.Protocol {
	case bwcli.PROTOCOL:
		return bwcli.NewProvider(), nil
	case file.PROTOCOL:
		return file.NewProvider(), nil
	case opcli.PROTOCOL:
		return opcli.NewProvider(), nil
	default:
		return nil, fmt.Errorf("unknown protocol : %s", resource.Protocol)
	}
}

func parseURIs(uris []string) (resources []*Resource, err error) {
	for _, uri := range uris {
		resource, err := parseURI(uri)
		if err != nil {
			return nil, err
		}
		resources = append(resources, resource)
	}

	return resources, nil
}
