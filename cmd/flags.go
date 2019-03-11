package cmd

var (
	tmpl  string //template path
	env   string //config env, multi-env split by ',' default: dev
	dist  string //config generate path; default: ./config
	ptype string //template file parser type, support json,yaml,yml,toml, default:json
	cache string //cache cmdline parameters into the system environment variable; on: cache, off: no cache, flush: flush cache
)
