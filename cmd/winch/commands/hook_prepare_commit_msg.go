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
	"bufio"
	"bytes"
	"context"
	"fmt"
	"github.com/mitchellh/colorstring"
	"github.com/spf13/cobra"
	"github.com/winchci/winch"
	"github.com/winchci/winch/config"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func q(r *bufio.Reader, prompt, suggestion string) (string, error) {
	fmt.Println(colorstring.Color(fmt.Sprintf("[white]%s[reset]", prompt)))
	fmt.Print("> ")
	text, err := r.ReadString('\n')
	if err != nil {
		return "", err
	}

	fmt.Println("q: ", strings.TrimSpace(text))
	return strings.TrimSpace(text), nil

	//if text, err := reader.PromptWithSuggestion("> ", suggestion, -1); err == nil {
	//	fmt.Println()
	//	return strings.TrimSpace(text), nil
	//} else if err == liner.ErrPromptAborted {
	//	return "", io.EOF
	//} else {
	//	return "", err
	//}
}

func getCommitType(reader *bufio.Reader) (string, error) {
	//reader.SetCompleter(func(l string) (c []string) {
	//	for _, n := range winch.GetCommitTypes() {
	//		if strings.HasPrefix(n.String(), strings.ToLower(l)) {
	//			c = append(c, n.String())
	//		}
	//	}
	//	return
	//})
	//
	//defer reader.SetCompleter(nil)

	//for n, ct := range winch.GetCommitTypes() {
	//	p("%2d) %s - [light_gray]%s[reset]", n, ct.String(), ct.Description())
	//}

	var s string
	var err error
	for err == nil {
		s, err = q(reader, "What type of change is this?", s)
		if err == nil {
			for _, n := range winch.GetCommitTypes() {
				if n.String() == s {
					return s, nil
				}
			}
		}
	}

	return s, err
}

func getScope(reader *bufio.Reader) (string, error) {
	return q(reader, "What is the scope of this change?", "")
}

func getMessage(reader *bufio.Reader, msg string) (string, error) {
	var s string
	var err error
	for err == nil && len(s) == 0 {
		if len(s) == 0 {
			s = msg
		}

		s, err = q(reader, "Describe this change", s)
	}

	return s, err
}

func getIsBreakingChange(reader *bufio.Reader) (bool, error) {
	//reader.SetCompleter(func(l string) (c []string) {
	//	if strings.HasPrefix("yes", strings.ToLower(l)) {
	//		return []string{"yes"}
	//	}
	//
	//	if strings.HasPrefix("no", strings.ToLower(l)) {
	//		return []string{"no"}
	//	}
	//	return
	//})
	//defer reader.SetCompleter(nil)

	var s string
	var err error

	for err == nil {
		s, err = q(reader, "Is this a breaking change? (yes/[bold]no[reset_bold])", s)
		if err == nil {
			if s == "yes" || s == "y" {
				return true, nil
			} else if s == "no" || s == "n" {
				return false, nil
			} else if len(s) == 0 {
				return false, nil
			}
		}
	}

	return false, nil
}

func getBreakingChangeMessage(reader *bufio.Reader) (string, error) {
	var s string
	var err error
	for err == nil && len(s) == 0 {
		s, err = q(reader, "Provide details on the breaking change", s)
	}

	return s, nil
}

func hookPrepareCommitMsg(ctx context.Context, args []string) error {
	var err error
	var msg string

	_, err = config.LoadConfig(ctx)
	if err != nil {
		return err
	}

	//line := liner.NewLiner()
	//defer line.Close()
	//line.SetCtrlCAborts(true)
	line := bufio.NewReader(os.Stdin)

	b, err := ioutil.ReadFile(args[0])
	if err != nil {
		return err
	}

	msg = strings.TrimSpace(string(b))

	commitType, err := getCommitType(line)
	if err != nil {
		if err == io.EOF {
			return nil
		}
		return err
	}

	scope, err := getScope(line)
	if err != nil {
		if err == io.EOF {
			return nil
		}
		return err
	}

	msg, err = getMessage(line, msg)
	if err != nil {
		if err == io.EOF {
			return nil
		}
		return err
	}

	var bcMsg string
	bc, err := getIsBreakingChange(line)
	if err != nil {
		if err == io.EOF {
			return nil
		}
		return err
	}

	if bc {
		bcMsg, err = getBreakingChangeMessage(line)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
	}

	buf := new(bytes.Buffer)

	fmt.Fprint(buf, commitType)
	if len(scope) > 0 {
		fmt.Fprintf(buf, "(%s)", scope)
	}
	fmt.Fprintln(buf, ":", msg)
	if bc {
		fmt.Fprintln(buf)
		fmt.Fprintln(buf, "BREAKING CHANGE:", bcMsg)
	}

	err = ioutil.WriteFile(args[0], buf.Bytes(), 0644)
	if err != nil {
		return err
	}

	return nil
}

func init() {
	var cmd = &cobra.Command{
		Use:   "prepare-commit-msg FILE TYPE [SHA1]",
		Short: "Hook for prepare-commit-msg",
		Run:   RunnerWithArgs(hookPrepareCommitMsg),
		Args:  cobra.RangeArgs(1, 3),
	}

	hookCmd.AddCommand(cmd)
}
