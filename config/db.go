package config

import (
	"context"
	"fmt"
	"github.com/spf13/pflag"
	"strconv"
)

// DatabaseConfig provides the database configuration.
type DatabaseConfig struct {
	Dialect      string `json:"dialect" yaml:"dialect"`
	Host         string `json:"host" yaml:"host"`
	Port         int    `json:"port" yaml:"port"`
	Database     string `json:"database" yaml:"database"`
	Username     string `json:"username" yaml:"username"`
	Password     string `json:"password" yaml:"password"`
	Dir          string `json:"dir" yaml:"dir"`
	Secure       bool   `json:"secure" yaml:"secure"`
	Timestamp    bool   `json:"timestamp" yaml:"timestamp"`
	UpdateConfig bool   `json:"updateconfig" yaml:"updateconfig"`
}

// String returns a string version of the config, suitable for passing as a connection string.
func (c DatabaseConfig) String() string {
	sslMode := "enable"
	if !c.Secure {
		sslMode = "disable"
	}

	return fmt.Sprintf("sslmode=%s host=%s dbname=%s user=%s password=%s port=%d",
		sslMode, c.Host, c.Database, c.Username, c.Password, c.Port)
}

func AddDB(flags *pflag.FlagSet) {
	flags.String("dialect", "postgres", "database dialect")
	flags.String("host", "localhost", "database host")
	flags.Int("port", 5432, "database port")
	flags.String("database", "", "database name")
	flags.String("username", "postgres", "database username")
	flags.String("password", "", "database password")
	flags.Bool("secure", false, "database secure connection")
}

func LoadDBConfig(ctx context.Context, cfg *Config) error {
	var err error
	cmd := CommandFromContext(ctx)
	flags := cmd.Flags()

	if cfg.Database == nil {
		cfg.Database = &DatabaseConfig{}
	}

	if flags.Lookup("dialect").Changed {
		cfg.Database.Dialect, err = flags.GetString("dialect")
		if err != nil {
			return err
		}
	}

	if len(cfg.Database.Dialect) == 0 {
		cfg.Database.Dialect = flags.Lookup("dialect").DefValue
	}

	if cfg.Database.Dialect != "postgres" {
		return fmt.Errorf("dialect %s not supported", cfg.Database.Dialect)
	}

	if flags.Lookup("host").Changed {
		cfg.Database.Host, err = flags.GetString("host")
		if err != nil {
			return err
		}
	}

	if len(cfg.Database.Host) == 0 {
		cfg.Database.Host = flags.Lookup("host").DefValue
	}

	if flags.Lookup("port").Changed {
		cfg.Database.Port, err = flags.GetInt("port")
		if err != nil {
			return err
		}
	}

	if cfg.Database.Port == 0 {
		cfg.Database.Port, err = strconv.Atoi(flags.Lookup("port").DefValue)
		if err != nil {
			return err
		}
	}

	if flags.Lookup("database").Changed {
		cfg.Database.Database, err = flags.GetString("database")
		if err != nil {
			return err
		}
	}

	if len(cfg.Database.Database) == 0 {
		cfg.Database.Database = cfg.Name
	}

	if flags.Lookup("username").Changed {
		cfg.Database.Username, err = flags.GetString("username")
		if err != nil {
			return err
		}
	}

	if len(cfg.Database.Username) == 0 {
		cfg.Database.Username = flags.Lookup("username").DefValue
	}

	if flags.Lookup("password").Changed {
		cfg.Database.Password, err = flags.GetString("password")
		if err != nil {
			return err
		}
	}

	if len(cfg.Database.Password) == 0 {
		cfg.Database.Password = flags.Lookup("password").DefValue
	}

	if flags.Lookup("secure").Changed {
		cfg.Database.Secure, err = flags.GetBool("secure")
		if err != nil {
			return err
		}
	}

	if flags.Lookup("timestamp") != nil && flags.Lookup("timestamp").Changed {
		cfg.Database.Timestamp, err = flags.GetBool("timestamp")
		if err != nil {
			return err
		}
	}

	if flags.Lookup("dir") != nil && flags.Lookup("dir").Changed {
		cfg.Database.Dir, err = flags.GetString("dir")
		if err != nil {
			return err
		}
	}

	if len(cfg.Database.Dir) == 0 && flags.Lookup("dir") != nil {
		cfg.Database.Dir = flags.Lookup("dir").DefValue
	}

	if len(cfg.Database.Dir) == 0 {
		cfg.Database.Dir = "./data"
	}

	cfg.Database.Dir = cfg.Database.Dir + "/" + cfg.Database.Dialect

	if flags.Lookup("update") != nil {
		cfg.Database.UpdateConfig, err = flags.GetBool("update")
		if err != nil {
			return err
		}
	}

	return nil
}
