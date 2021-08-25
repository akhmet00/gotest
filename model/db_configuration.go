package model

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Config struct {
	Name    string
	Pass    string
	Address string
	DbName  string
	SSLMode string
}

func ReadConfig() Config {

	confFile, _ := os.Open("config.json")

	jsonConf, _ := ioutil.ReadAll(confFile)

	var conf Config

	json.Unmarshal(jsonConf, &conf)

	return conf
}
