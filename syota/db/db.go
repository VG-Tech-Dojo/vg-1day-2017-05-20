package db

import (
	"database/sql"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

// Configs はdbconfig.ymlから複数のenvでConfigを読むためのmapです
type Configs map[string]*Config

// Open はenvで指定された設定で
func (cs Configs) Open(env string) (*sql.DB, error) {
	config, ok := cs[env]
	if !ok {
		return nil, fmt.Errorf("no such env in config file: %s", env)
	}
	return config.Open()
}

// Config はdbconfig.ymlを読むための構造体です
type Config struct {
	Datasource string `yaml:"datasource"`
}

// Open は新しくデータベースとのコネクションを返します
func (c *Config) Open() (*sql.DB, error) {
	return sql.Open("sqlite3", c.Datasource)
}

// NewConfigsFromFile はファイルパスから新しいConfigsを返します
func NewConfigsFromFile(path string) (Configs, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return NewConfigs(f)
}

// NewConfigs はyamlを読み込んで新しいConfigsを返します
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
