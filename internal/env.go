package internal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var (
	CONFC_TMPL  = "CONFC_TMPL"  //(相对/非相对路径)  --tmpl (当前路径 *.confc文件) [数组]
	CONFC_ENV   = "CONFC_ENV"   //(环境以逗号隔离) --env (dev)
	CONFC_DIST  = "CONFC_DIST"  //默认生成路径 --dist (config)
	CONFC_PTYPE = "CONFC_PTYPE" //解析类型;  --type (json)
	CONFC_CACHE = "CONFC_CACHE" //命令行参数缓存策略 --cache(on)
)
var CONFC_ENVS = map[string]string{
	CONFC_TMPL:  CONFC_TMPL,  //(相对/非相对路径)  --tmpl (当前路径 *.confc文件) [数组]
	CONFC_ENV:   CONFC_ENV,   //(环境以逗号隔离) --env (dev)
	CONFC_DIST:  CONFC_DIST,  //默认生成路径 --dist (config)
	CONFC_PTYPE: CONFC_PTYPE, //解析类型;  --type (json)
	CONFC_CACHE: CONFC_CACHE, //命令行参数缓存策略 --cache(on)
}

const (
	DEFAULT_CONFC_ENV   = "dev"
	DEFAULT_CONFC_DIST  = "config/"
	DEFAULT_CONFC_PTYPE = JSON_PTYPE
	DEFAULT_CONFC_CACHE = ON_CACHE

	JSON_PTYPE = "json"
	TOML_PTYPE = "toml"
	YAML_PTYPE = "yaml"
	YML_PTYPE  = "yml"

	ON_CACHE  = "on"
	OFF_CACHE = "off"
)

var CONFC_DEFAULTS = map[string]string{
	CONFC_ENV:   DEFAULT_CONFC_ENV,
	CONFC_DIST:  DEFAULT_CONFC_DIST,
	CONFC_PTYPE: DEFAULT_CONFC_PTYPE,
	CONFC_CACHE: DEFAULT_CONFC_CACHE,
}

var ENVS map[string]string

func confPath() (string, error) {
	home, err := Home()
	if err != nil {
		return "", err
	}
	confPath := fmt.Sprintf("%s/.confc", home)
	err, _ = CreatePath(confPath)
	confPath = fmt.Sprintf("%s/conf.json", confPath)
	return confPath, err
}

func setenv(envs map[string]string) error {
	filePath, err := confPath()
	if err != nil {
		return err
	}
	f, err := os.OpenFile(filePath, os.O_CREATE|os.O_TRUNC|os.O_RDWR, os.ModePerm)
	if err != nil {
		return err
	}
	envs[CONFC_CACHE] = ON_CACHE
	envsBytes, _ := json.Marshal(envs)
	_, err = f.Write(envsBytes)
	if err == nil {
		ENVS = envs
	}
	return err
}

func unsetenv() error {
	filePath, err := confPath()
	if err != nil {
		return err
	}
	envsBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	envs := make(map[string]string)
	err = json.Unmarshal(envsBytes, &envs)
	if err != nil {
		return err
	}
	if envs[CONFC_CACHE] == OFF_CACHE {
		return nil
	}
	envs[CONFC_CACHE] = OFF_CACHE
	f, err := os.OpenFile(filePath, os.O_CREATE|os.O_TRUNC|os.O_RDWR, os.ModePerm)
	if err != nil {
		return err
	}
	envsBytes, err = json.Marshal(envs)
	if err != nil {
		return err
	}
	_, err = f.Write(envsBytes)
	if err == nil {
		ENVS = envs
	}
	return err
}

func lookupenv(key string) (string, error) {
	if ENVS != nil {
		if val, ok := ENVS[key]; ok && len(val) != 0 {
			return val, nil
		}
	}
	filePath, err := confPath()
	if err != nil {
		return "", err
	}
	envsBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	envs := make(map[string]string)
	err = json.Unmarshal(envsBytes, &envs)
	if err != nil {
		return "", err
	}
	if envs[CONFC_CACHE] == OFF_CACHE {
		return "", fmt.Errorf("config is disabled.\n")
	}
	ENVS = envs
	return envs[key], nil
}

func show() {
	if v, err := lookupenv(CONFC_TMPL); err == nil {
		log.Printf("%v=%v\n", CONFC_TMPL, v)
	}
	if v, err := lookupenv(CONFC_ENV); err == nil {
		log.Printf("%v=%v\n", CONFC_ENV, v)
	}
	if v, err := lookupenv(CONFC_DIST); err == nil {
		log.Printf("%v=%v\n", CONFC_DIST, v)
	}
	if v, err := lookupenv(CONFC_PTYPE); err == nil {
		log.Printf("%v=%v\n", CONFC_PTYPE, v)
	}
	if v, err := lookupenv(CONFC_CACHE); err == nil {
		log.Printf("%v=%v\n", CONFC_CACHE, v)
	}
}
