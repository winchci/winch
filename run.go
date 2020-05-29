package winch

import (
	"bytes"
	"context"
	"fmt"
	"github.com/switch-bit/winch/config"
	"os"
	"os/exec"
	"strings"
)

func Run(ctx context.Context, cmd *config.RunConfig, cfg *config.Config) error {
	if !cmd.IsEnabled() {
		return nil
	}

	if !CheckFilters(ctx, cmd.Branches, cmd.Tags) {
		return nil
	}

	c := exec.CommandContext(ctx, "sh", "-c", cmd.Command)
	env := make(map[string]string)

	for _, e := range os.Environ() {
		parts := strings.SplitN(e, "=", 2)
		env[parts[0]] = parts[1]
	}

	for k, v := range cfg.Environment {
		env[k] = v
	}

	for k, v := range cmd.Environment {
		env[k] = v
	}

	for replacementsRequired := true; replacementsRequired; {
		replacementsRequired = false

		for k, v := range env {
			env[k] = os.Expand(v, func(s string) string {
				return env[s]
			})

			if strings.Contains(env[k], "$") {
				replacementsRequired = true
			}
		}
	}

	for k, v := range env {
		c.Env = append(c.Env, fmt.Sprintf("%s=%s", k, v))
	}

	if len(cmd.Input) > 0 {
		c.Stdin = bytes.NewBufferString(cmd.Input)
	}

	if cfg.Verbose {
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
	}

	if len(cmd.Name) > 0 {
		fmt.Println(cmd.Name)
	}

	if !cfg.Quiet {
		fmt.Println(cmd.Command)
	}

	return c.Run()
}
