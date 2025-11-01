package config

import "testing"

func TestNewSetting(t *testing.T) {
	type Test struct {
		value any
	}

	tests := []Test{
		{value: "https"},
		{value: 5000},
		{value: true},
	}

	for _, test := range tests {
		t.Run("NewSetting()", func(t *testing.T) {
			setting := NewSetting(test.value)
			expect := test.value
			result := setting.Value
			if result != expect {
				t.Errorf("Setting.Value = %v; expect %v", result, expect)
			}
		})
	}
}

func TestSetting_Resolve(t *testing.T) {
	type Test struct {
		name   string
		key    string
		value  string
		expect any
	}

	tests := []Test{
		{name: "Setting[string].Resolve()", key: "TESTING", value: "https", expect: "https"},
		{name: "Setting[number].Resolve()", key: "TESTING", value: "5000", expect: 5000},
		{name: "Setting[boolean].Resolve()", key: "TESTING", value: "true", expect: true},
	}

	index := 0
	t.Run(tests[index].name, func(t *testing.T) {
		t.Setenv(tests[index].key, tests[index].value)
		env := EnvironmentVariable(tests[index].key)
		result := Setting[string]{}.Resolve(ConvString(env, false))
		if result != tests[index].expect {
			t.Errorf("Setting.Resolve() = %v; expect %v", result, tests[index].expect)
		}
	})

	index = 1
	t.Run(tests[index].name, func(t *testing.T) {
		t.Setenv(tests[index].key, tests[index].value)
		env := EnvironmentVariable(tests[index].key)
		result := Setting[int]{}.Resolve(ConvInt(env))
		if result != tests[index].expect {
			t.Errorf("Setting.Resolve() = %v; expect %v", result, tests[index].expect)
		}
	})

	index = 2
	t.Run(tests[index].name, func(t *testing.T) {
		t.Setenv(tests[index].key, tests[index].value)
		env := EnvironmentVariable(tests[index].key)
		result := Setting[bool]{}.Resolve(ConvBool(env))
		if result != tests[index].expect {
			t.Errorf("Setting.Resolve() = %v; expect %v", result, tests[index].expect)
		}
	})
}
