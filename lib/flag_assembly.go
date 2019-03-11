package lib

import "fmt"

func Assembly(gcc *GenConfCtx) error {
	if len(gcc.Tmpl) == 0 {
		return fmt.Errorf("tmpl invalid.")
	}
	if gcc.Cache == ON_CACHE {
		c, ok := lookupenv(CONFC_CACHE)
		if ok {
			gcc.Cache = c
		}
	}
	for k, v := range map[string]string{
		CONFC_ENV:   gcc.Env,
		CONFC_DIST:  gcc.Dist,
		CONFC_PTYPE: gcc.Ptype,
	} {
		v = assembly(k, v, gcc.Cache)
		if len(v) == 0 {
			return fmt.Errorf("%v invalid.\n", k)
		}
		switch k {
		case CONFC_ENV:
			gcc.Env = v
		case CONFC_DIST:
			gcc.Dist = v
		case CONFC_PTYPE:
			gcc.Ptype = v
		}
	}
	switch gcc.Cache {
	case FLUSH_CACHE:
		unsetenv(CONFC_TMPL)
		unsetenv(CONFC_ENV)
		unsetenv(CONFC_DIST)
		unsetenv(CONFC_PTYPE)
		unsetenv(CONFC_CACHE)
	case ON_CACHE:
		err := setenv(CONFC_TMPL, gcc.Tmpl)
		fmt.Println("set TMPL env - ", err)
		err = setenv(CONFC_ENV, gcc.Env)
		fmt.Println("set ENV env - ", err)
		err = setenv(CONFC_DIST, gcc.Dist)
		fmt.Println("set DIST env - ", err)
		err = setenv(CONFC_PTYPE, gcc.Ptype)
		fmt.Println("set PTYPE env - ", err)
		err = setenv(CONFC_CACHE, gcc.Cache)
		fmt.Println("set CACHE env - ", err)
	}
	show()
	return nil
}

func tmplAssembly(gcc *GenConfCtx) {
	if len(gcc.Tmpl) != 0 {
		return
	}
	fmt.Println("1")
	if gcc.Cache == FLUSH_CACHE {
		fmt.Println("2")
		return
	}
	fmt.Println("3")

	tmpl, ok := lookupenv(CONFC_TMPL)
	if ok {
		gcc.Tmpl = tmpl

	}

}

func assembly(key, val, cache string) string {
	if len(val) != 0 && val != CONFC_DEFAULTS[key] {
		return val
	}
	if cache == FLUSH_CACHE {
		return val
	}
	envVal, ok := lookupenv(key)
	if ok {
		return envVal

	}
	return val
}
