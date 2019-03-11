package lib

type GenConfCtx struct {
	Tmpl  string //template path
	Env   string //config env, multi-env split by ',' default: dev
	Dist  string //config generate path; default: ./config
	Ptype string //template file parser type, support json,yaml,yml,toml, default:json
	Cache string //cache cmdline parameters into the system environment variable; on: cache, off: no cache, flush: flush cache
}
