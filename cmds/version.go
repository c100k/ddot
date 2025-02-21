package cmds

import (
	"c100k/ddot/ui"
	"fmt"
)

func Version(version string) error {
	ui.Print("ℹ️", fmt.Sprintf("Version : %s", version))

	return nil
}
