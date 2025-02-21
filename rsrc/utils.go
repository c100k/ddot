package rsrc

import (
	"fmt"
	"strings"
)

func parseURI(uri string) (*Resource, error) {
	parts := strings.Split(uri, URI_PROTOCOL_PATH_SEP)
	length := len(parts)
	if length != URI_SPLIT_EXPECTED_LENGTH {
		return nil, fmt.Errorf("invalid resource uri : %v (expected length of %d but got %d)", parts, URI_SPLIT_EXPECTED_LENGTH, length)
	}

	return &Resource{
		Path:     parts[1],
		Protocol: parts[0],
		Uri:      uri,
	}, nil
}
