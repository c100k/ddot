package cmds

import (
	"c100k/ddot/output"
	"c100k/ddot/rsrc"
	"c100k/ddot/ui"
	"fmt"
	"os"
	"os/signal"
)

func LoadEnv(out string, uris []string) error {
	exists, err := output.Exists(out)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("'%s' already exists. stopping here to avoid overriding it", out)
	}

	resources, err := rsrc.ReadAll(uris)
	if err != nil {
		return err
	}

	err = output.Write(out, resources)
	if err != nil {
		return err
	}

	err = waitForInterrupt(out)
	if err != nil {
		return err
	}

	return nil
}

func waitForInterrupt(out string) error {
	channel := make(chan os.Signal, 1)
	signal.Notify(channel, os.Interrupt)

	ui.Print("🚀", "Press Ctrl+C once you're done  => it will remove the file")

	sig := <-channel

	ui.Print("", "")
	ui.Print("🤙", fmt.Sprintf("Received signal '%s'", sig))
	err := output.Clean(out)

	return err
}
