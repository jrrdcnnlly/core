package config

import "testing"

func TestEnvironmentVariable_Key(t *testing.T) {
	type Test struct {
		name   string
		key    string
		expect string
	}

	type TestBatch struct {
		tests []Test
	}

	var batches []TestBatch = []TestBatch{
		{
			tests: []Test{
				{name: "Key() should be \"HTTP_HOST\"", key: "HTTP_HOST", expect: "HTTP_HOST"},
				{name: "Key() should be \"HTTP_PORT\"", key: "HTTP_PORT", expect: "HTTP_PORT"},
				{name: "Key() should be \"SEARCH_URL\"", key: "SEARCH_URL", expect: "SEARCH_URL"},
			},
		},
	}

	for _, job := range batches {
		for _, test := range job.tests {
			t.Run(test.name, func(t *testing.T) {
				t.Setenv(test.key, "placeholder")
				var env EnvironmentVariable = EnvironmentVariable(test.key)
				result := env.Key()
				if result != test.expect {
					t.Errorf("EnvironmentVariable.Key() = %q; expect %q", result, test.expect)
				}
			})
		}
	}
}

func TestEnvironmentVariable_Get(t *testing.T) {
	type Test struct {
		name   string
		key    string
		value  string
		expect string
	}

	type TestBatch struct {
		tests []Test
	}

	var batches []TestBatch = []TestBatch{
		{
			tests: []Test{
				{name: "Get() should be \"localhost\"", key: "HTTP_HOST", value: "localhost", expect: "localhost"},
				{name: "Get() should be \"5000\"", key: "HTTP_PORT", value: "5000", expect: "5000"},
				{name: "Get() should be \"http://www.google.com\"", key: "SEARCH_URL", value: "http://www.google.com", expect: "http://www.google.com"},
			},
		},
	}

	for _, job := range batches {
		for _, test := range job.tests {
			t.Run(test.name, func(t *testing.T) {
				t.Setenv(test.key, test.value)
				var env EnvironmentVariable = EnvironmentVariable(test.key)
				result := env.Get()
				if result != test.expect {
					t.Errorf("EnvironmentVariable.Get() = %q; expect %q", result, test.expect)
				}
			})
		}
	}
}

func TestEnvironmentVariable_Lookup(t *testing.T) {
	type Test struct {
		name   string
		key    string
		value  string
		expect string
	}

	type TestBatch struct {
		tests []Test
	}

	var batches []TestBatch = []TestBatch{
		{
			tests: []Test{
				{name: "Lookup() should be \"localhost\", true", key: "HTTP_HOST", value: "localhost", expect: "localhost"},
				{name: "Lookup() should be \"5000\", true", key: "HTTP_PORT", value: "5000", expect: "5000"},
				{name: "Lookup() should be \"http://www.google.com\", true", key: "SEARCH_URL", value: "http://www.google.com", expect: "http://www.google.com"},
			},
		},
	}

	for _, job := range batches {
		for _, test := range job.tests {
			t.Run(test.name, func(t *testing.T) {
				t.Setenv(test.key, test.value)
				var env EnvironmentVariable = EnvironmentVariable(test.key)
				result, ok := env.Lookup()
				if !ok {
					t.Errorf("EnvironmentVariable.Lookup() = %q, false; expect %q, true", result, test.expect)
				}
				if result != test.expect {
					t.Errorf("EnvironmentVariable.Lookup() = %q, true; expect %q, true", result, test.expect)
				}
			})
		}
	}
}

func TestEnvironmentVariable_String(t *testing.T) {
	type Test struct {
		name   string
		key    string
		expect string
	}

	type TestBatch struct {
		tests []Test
	}

	var batches []TestBatch = []TestBatch{
		{
			tests: []Test{
				{name: "String() should be \"HTTP_HOST\"", key: "HTTP_HOST", expect: "HTTP_HOST"},
				{name: "String() should be \"HTTP_PORT\"", key: "HTTP_PORT", expect: "HTTP_PORT"},
				{name: "String() should be \"SEARCH_URL\"", key: "SEARCH_URL", expect: "SEARCH_URL"},
			},
		},
	}

	for _, job := range batches {
		for _, test := range job.tests {
			t.Run(test.name, func(t *testing.T) {
				t.Setenv(test.key, "placeholder")
				var env EnvironmentVariable = EnvironmentVariable(test.key)
				result := env.Key()
				if result != test.expect {
					t.Errorf("EnvironmentVariable.String() = %q; expect %q", result, test.expect)
				}
			})
		}
	}
}
