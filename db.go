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

package winch

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"

	"github.com/winchci/winch/config"
)

func Psql(cfg *config.Config, path string, database string) error {
	if cfg.Verbose {
		fmt.Println(path)
	}

	password := "-w"
	if len(cfg.Database.Password) > 0 {
		password = "-W"
	}

	c := exec.Command("psql",
		"-h", cfg.Database.Host,
		"-p", strconv.Itoa(cfg.Database.Port),
		"-U", cfg.Database.Username,
		password,
		"-d", database,
		"-f", path,
		"-1",
		"-q",
		"-b",
	)

	if !cfg.Quiet {
		c.Stdin = os.Stdin
		c.Stderr = os.Stderr
	}

	err := c.Run()
	if err != nil {
		return err
	}

	return nil
}
