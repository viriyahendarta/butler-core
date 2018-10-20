package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sync"

	"github.com/viriyahendarta/butler-core/infra/constant"
)

const (
	configPath string = "./config/"
	appEnv     string = "APPENV"
)

var conf *Config
var once sync.Once
var err error

//Get initializes config (if not yet initialized) then returns Config
func Get() Config {
	Init()
	return *conf
}

//Init initialized config by reading config json
func Init() error {
	once.Do(func() {
		conf, err = readConfig()
	})
	return err
}

//GetEnv returns application env
func GetEnv() string {
	env := os.Getenv(appEnv)

	switch env {
	case constant.EnvProduction, constant.EnvStaging, constant.EnvDevelopment, constant.EnvTest:
	default:
		env = constant.EnvDevelopment
	}

	return env
}

func readConfig() (*Config, error) {
	env := GetEnv()

	path := fmt.Sprintf("%s%s.json", configPath, env)
	c, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	result := new(Config)
	err = json.Unmarshal(c, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
