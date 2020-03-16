package config

import (
	"desafioPG/model"
	"errors"
	"os"

	"gopkg.in/yaml.v2"
)

// Config -> .
type Config struct {
	Server struct {
		Broker1 []string `yaml:"1"`
		Broker2 []string `yaml:"2"`
		Broker3 []string `yaml:"3"`
	} `yaml:"brokers"`
	DDD []int  `yaml:"ddd"`
	URL string `yaml:"url"`
}

// GetConf ->
func GetConf() (con Config, err error) {
	err = nil

	f, err := os.Open("/home/victor/git/estudo/golang/src/desafioPG/config/config.yml")
	if err != nil {
		err = errors.New(model.ErrLerArquivoYaml.Error() + " -> [" + err.Error() + "]")
		return
	}

	defer f.Close()

	var cfg Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		err = errors.New(model.ErrConverterArquivoYaml.Error() + " -> [" + err.Error() + "]")
		return
	}

	return cfg, err
}
