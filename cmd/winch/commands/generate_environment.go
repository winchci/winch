package commands

import (
	"bytes"
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/switch-bit/winch/config"
	"io/ioutil"
)

func generateEnvironment(ctx context.Context, args []string) error {
	ctx, err := config.LoadConfig(ctx)
	if err != nil {
		return err
	}

	cfg := config.ConfigFromContext(ctx)

	if env, ok := cfg.Environments[args[0]]; ok {
		buf := bytes.NewBuffer(nil)
		for k, v := range env {
			buf.WriteString(fmt.Sprintf("%s=%s\n", k, v))
		}
		return ioutil.WriteFile(".env", buf.Bytes(), 0644)
	} else {
		return fmt.Errorf("environment %s not defined", args[0])
	}
}

func init() {
	var cmd = &cobra.Command{
		Use:   "environment",
		Short: "Generate .env file",
		Run:   RunnerWithArgs(generateEnvironment),
		Args:  cobra.ExactArgs(1),
	}

	generateCmd.AddCommand(cmd)
}
