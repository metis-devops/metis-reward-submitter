package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Eth       *Eth       `json:"eth" yaml:"eth"`
	Metis     *Metis     `json:"metis" yaml:"metis"`
	Submitter *Submitter `json:"submitter" yaml:"submitter"`
}

func (conf *Config) mergeFromEnv() {
	if rpc := os.Getenv("ETH_RPC"); rpc != "" {
		conf.Eth.RPC = rpc
	}

	if rpc := os.Getenv("METIS_RPC"); rpc != "" {
		conf.Metis.RPC = rpc
	}

	if rpc := os.Getenv("THEMIS_REST"); rpc != "" {
		conf.Metis.Themis = rpc
	}
}

func Parse(p string) (*Config, error) {
	file, err := os.ReadFile(p)
	if err != nil {
		return nil, err
	}

	conf := new(Config)
	switch ext := path.Ext(p); ext {
	case ".json":
		err = json.Unmarshal(file, &conf)
	case ".yaml", ".yml":
		err = yaml.Unmarshal(file, &conf)
	default:
		err = fmt.Errorf("not supported file extension %s", ext)
	}
	if err != nil {
		return nil, err
	}

	conf.mergeFromEnv()

	return conf, nil
}
