package db

import (
	"database/sql"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type Configs map[string]*Config

func (cs Configs) Open(env string) (*sql.DB, error) {
	config, ok := cs[env]
	if !ok {
		return nil, fmt.Errorf("no such env in config file: %s", env)
	}
	return config.Open()
}

type Config struct {
	Datasource string `yaml:"datasource"`
}

func (c *Config) Open() (*sql.DB, error) {
	return sql.Open("sqlite3", c.Datasource)
}

func NewConfigsFromFile(path string) (Configs, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return NewConfigs(f)
}

func NewConfigs(r io.Reader) (Configs, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	var configs Configs
	if err = yaml.Unmarshal(b, &configs); err != nil {
		return nil, err
	}
	return configs, nil
}
