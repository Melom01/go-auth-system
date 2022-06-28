package config_test

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"path/filepath"
	"runtime"
	"sentinel/config"
	"testing"
)

func TestConfig(t *testing.T) {
	var raw []byte
	var err error

	_, b, _, _ := runtime.Caller(0)
	basePath := filepath.Dir(b)

	if raw, err = ioutil.ReadFile(basePath + "/config_test.json"); err != nil {
		log.Fatal("Unable to read configuration file: ", err)
	}
	if err = json.Unmarshal(raw, &config.Config); err != nil {
		log.Fatal("Unable to parse configuration file: ", err)
	}
}
