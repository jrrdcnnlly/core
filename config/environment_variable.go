package config

import "os"

// OS environemnt variable.
type EnvironmentVariable string

// Return the environment variable key.
func (e EnvironmentVariable) Key() string {
	return string(e)
}

// Return the environemnt variable value. If the variable is not set an empty string is returned.
func (e EnvironmentVariable) Get() string {
	return os.Getenv(string(e))
}

// Return the environemnt variable value. If the variable is not set false is returned as the second value.
func (e EnvironmentVariable) Lookup() (string, bool) {
	return os.LookupEnv(string(e))
}

// Implement the fmt.Stringer interface.
func (e EnvironmentVariable) String() string {
	return os.Getenv(string(e))
}
