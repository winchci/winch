/*
winch - Universal Build and Release Tool
Copyright (C) 2020 Switchbit, Inc.

This program is free software: you can redistribute it and/or modify it under the terms of the GNU
General Public License as published by the Free Software Foundation, either version 3 of the License,
or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even
the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public
License for more details.

You should have received a copy of the GNU General Public License along with this program. If not,
see <https://www.gnu.org/licenses/>.
*/

package commands

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/winchci/winch"
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
