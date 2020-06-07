package actions

import (
	"context"
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
)

func LoadActionDefinition(ctx context.Context, filename string) (*ActionDefinition, error) {
	if len(filename) == 0 {
		filename = "winch-action.yml"
	}

	if _, err := os.Stat(filename); err != nil {
		return nil, fmt.Errorf("cannot open action file %s", filename)
	}

	b , err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	cfg := new(ActionDefinition)
	err = yaml.Unmarshal(b, cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
