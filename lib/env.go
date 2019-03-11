package lib

import (
	"log"
	"os"
)

//custom-defined-config
//env-defined-config
//default-config

var (
	CONFC_TMPL  = "CONFC_TMPL"  //(相对/非相对路径)  --tmpl (当前路径 *.confc文件) [数组]
	CONFC_ENV   = "CONFC_ENV"   //(环境以逗号隔离) --env (dev)
	CONFC_DIST  = "CONFC_DIST"  //默认生成路径 --dist (config)
	CONFC_PTYPE = "CONFC_PTYPE" //解析类型;  --type (json)
	CONFC_CACHE = "CONFC_CACHE" //命令行参数缓存策略 --cache(on)
)
var CONFC_ENVS = map[string]string{
	CONFC_TMPL:  "CONFC_TMPL",  //(相对/非相对路径)  --tmpl (当前路径 *.confc文件) [数组]
	CONFC_ENV:   "CONFC_ENV",   //(环境以逗号隔离) --env (dev)
	CONFC_DIST:  "CONFC_DIST",  //默认生成路径 --dist (config)
	CONFC_PTYPE: "CONFC_PTYPE", //解析类型;  --type (json)
	CONFC_CACHE: "CONFC_CACHE", //命令行参数缓存策略 --cache(on)
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

	ON_CACHE    = "on"
	OFF_CACHE   = "off"
	FLUSH_CACHE = "flush"
)

var CONFC_DEFAULTS = map[string]string{
	CONFC_ENV:   DEFAULT_CONFC_ENV,
	CONFC_DIST:  DEFAULT_CONFC_DIST,
	CONFC_PTYPE: DEFAULT_CONFC_PTYPE,
	CONFC_CACHE: DEFAULT_CONFC_CACHE,
}

func setenv(key, value string) error {
	return os.Setenv(key, value)
}

func unsetenv(key string) error {
	return os.Unsetenv(key)
}

func lookupenv(key string) (string, bool) {
	return os.LookupEnv(key)
}

func show() {
	if v, ok := lookupenv(CONFC_TMPL); ok {
		log.Printf("%v=%v\n", CONFC_TMPL, v)
	}
	if v, ok := lookupenv(CONFC_ENV); ok {
		log.Printf("%v=%v\n", CONFC_ENV, v)
	}
	if v, ok := lookupenv(CONFC_DIST); ok {
		log.Printf("%v=%v\n", CONFC_DIST, v)
	}
	if v, ok := lookupenv(CONFC_PTYPE); ok {
		log.Printf("%v=%v\n", CONFC_PTYPE, v)
	}
	if v, ok := lookupenv(CONFC_CACHE); ok {
		log.Printf("%v=%v\n", CONFC_CACHE, v)
	}
}
