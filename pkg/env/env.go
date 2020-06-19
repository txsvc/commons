package env

import "os"

// Getenv returns the environment variable env or def if env is not set.
// Note: def is only returned if the env is not set, i.e. an EMPTY env is still returned !
func Getenv(env, def string) string {
	e, ok := os.LookupEnv(env)
	if ok {
		return e
	}
	return def
}
