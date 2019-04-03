package internal

import "fmt"

//Assembly assembly config
func Assembly(gcc *GenConfCtx) error {
	if len(gcc.Tmpl) == 0 && gcc.Cache == OFF_CACHE {
		return unsetenv()
	}
	if len(gcc.Tmpl) == 0 {
		return fmt.Errorf("tmpl invalid.")
	}
	if gcc.Cache == ON_CACHE {
		c, err := lookupenv(CONFC_CACHE)
		if err == nil {
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
	case OFF_CACHE:
		unsetenv()
	case ON_CACHE:
		err := setenv(map[string]string{
			CONFC_TMPL:  gcc.Tmpl,
			CONFC_ENV:   gcc.Env,
			CONFC_DIST:  gcc.Dist,
			CONFC_PTYPE: gcc.Ptype,
			CONFC_CACHE: gcc.Cache,
		})
		if err != nil {
			return err
		}
	}
	show()
	return nil
}

func tmplAssembly(gcc *GenConfCtx) {
	if len(gcc.Tmpl) != 0 {
		return
	}
	if gcc.Cache == OFF_CACHE {
		return
	}

	tmpl, err := lookupenv(CONFC_TMPL)
	if err == nil {
		gcc.Tmpl = tmpl

	}

}

func assembly(key, val, cache string) string {
	if len(val) != 0 && val != CONFC_DEFAULTS[key] {
		return val
	}
	if cache == OFF_CACHE {
		return val
	}
	envVal, err := lookupenv(key)
	if err == nil {
		return envVal

	}
	return val
}
