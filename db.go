package winch

import (
	"fmt"
	"github.com/switch-bit/winch/config"
	"os"
	"os/exec"
	"strconv"
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
