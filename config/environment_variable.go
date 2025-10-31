package config

import "os"

type EnvironmentVariable string

func (e EnvironmentVariable) Key() string {
	return string(e)
}

func (e EnvironmentVariable) Get() string {
	return os.Getenv(string(e))
}

func (e EnvironmentVariable) Lookup() (string, bool) {
	return os.LookupEnv(string(e))
}

func (e EnvironmentVariable) String() string {
	return os.Getenv(string(e))
}
