package commands

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/switch-bit/winch"
	"io/ioutil"
	"strings"
)

func hookCommitMsg(_ context.Context, args []string) error {
	b, err := ioutil.ReadFile(args[0])
	if err != nil {
		return err
	}

	m := strings.TrimSpace(string(b))
	if strings.HasPrefix(m, "Merge pull request") {
		m = "chore(merge): merge " + strings.TrimPrefix(m, "Merge ")
	}

	msg := winch.ParseMessage(m)

	if msg.Type.String() == "change" {
		return fmt.Errorf("invalid commit message format")
	}

	if len(msg.Subject) == 0 {
		return fmt.Errorf("invalid commit message format")
	}

	return nil
}

func init() {
	var cmd = &cobra.Command{
		Use:   "commit-msg FILE TYPE [SHA1]",
		Short: "Hook for commit-msg",
		Run:   RunnerWithArgs(hookCommitMsg),
		Args:  cobra.RangeArgs(1, 3),
	}

	hookCmd.AddCommand(cmd)
}
