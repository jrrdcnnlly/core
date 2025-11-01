package config

import (
	"log/slog"
	"net/url"
	"slices"
	"testing"
)

func TestFallback(t *testing.T) {
	type Test struct {
		name   string
		value  any
		expect any
	}

	tests := []Test{
		{name: "TestFallback[string]()", value: "https", expect: "https"},
		{name: "TestFallback[int]()", value: 5000, expect: 5000},
		{name: "TestFallback[bool]()", value: true, expect: true},
	}

	index := 0
	t.Run(tests[index].name, func(t *testing.T) {
		var setting Setting[string]
		result := Fallback(tests[index].value.(string))(setting).Value
		if result != tests[index].expect {
			t.Errorf("Fallback() = %v; expected %v", result, tests[index].expect)
		}
	})

	index = 1
	t.Run(tests[index].name, func(t *testing.T) {
		var setting Setting[int]
		result := Fallback(tests[index].value.(int))(setting).Value
		if result != tests[index].expect {
			t.Errorf("Fallback() = %v; expected %v", result, tests[index].expect)
		}
	})

	index = 2
	t.Run(tests[index].name, func(t *testing.T) {
		var setting Setting[bool]
		result := Fallback(tests[index].value.(bool))(setting).Value
		if result != tests[index].expect {
			t.Errorf("Fallback() = %v; expected %v", result, tests[index].expect)
		}
	})
}

func TestPanicIfUnset(t *testing.T) {
	defer func() { _ = recover() }()

	var setting Setting[string]
	setting.Resolve(PanicIfUnset[string]())

	t.Errorf("TestPanicIfUnset() did not panic")
}

func TestConvBool(t *testing.T) {
	type Test struct {
		name   string
		value  string
		expect bool
	}

	tests := []Test{
		{name: "ConvBool() shoud be true", value: "true", expect: true},
		{name: "ConvBool() shoud be true", value: "t", expect: true},
		{name: "ConvBool() shoud be true", value: "1", expect: true},
		{name: "ConvBool() shoud be false", value: "false", expect: false},
		{name: "ConvBool() shoud be false", value: "f", expect: false},
		{name: "ConvBool() shoud be false", value: "f", expect: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Setenv("TESTING", test.value)
			env := EnvironmentVariable("TESTING")
			var setting Setting[bool]
			result := ConvBool(env)(setting).Value
			if result != test.expect {
				t.Errorf("got %v; expect %v", result, test.expect)
			}
		})
	}
}

func TestConvFloat32(t *testing.T) {
	type Test struct {
		name   string
		value  string
		expect float32
	}

	tests := []Test{
		{name: "ConvFloat32() shoud be 0.1", value: "0.1", expect: float32(0.1)},
		{name: "ConvFloat32() shoud be 0.2", value: "0.2", expect: float32(0.2)},
		{name: "ConvFloat32() shoud be 0.3", value: "0.3", expect: float32(0.3)},
		{name: "ConvFloat32() shoud be 0.4", value: "0.4", expect: float32(0.4)},
		{name: "ConvFloat32() shoud be 1.0", value: "1.0", expect: float32(1.0)},
		{name: "ConvFloat32() shoud be 2.0", value: "2.0", expect: float32(2.0)},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Setenv("TESTING", test.value)
			env := EnvironmentVariable("TESTING")
			var setting Setting[float32]
			result := ConvFloat32(env)(setting).Value
			if result != test.expect {
				t.Errorf("got %v; expect %v", result, test.expect)
			}
		})
	}
}

func TestConvFloat64(t *testing.T) {
	type Test struct {
		name   string
		value  string
		expect float64
	}

	tests := []Test{
		{name: "ConvFloat64() shoud be 0.1", value: "0.1", expect: float64(0.1)},
		{name: "ConvFloat64() shoud be 0.2", value: "0.2", expect: float64(0.2)},
		{name: "ConvFloat64() shoud be 0.3", value: "0.3", expect: float64(0.3)},
		{name: "ConvFloat64() shoud be 0.4", value: "0.4", expect: float64(0.4)},
		{name: "ConvFloat64() shoud be 1.0", value: "1.0", expect: float64(1.0)},
		{name: "ConvFloat64() shoud be 2.0", value: "2.0", expect: float64(2.0)},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Setenv("TESTING", test.value)
			env := EnvironmentVariable("TESTING")
			var setting Setting[float64]
			result := ConvFloat64(env)(setting).Value
			if result != test.expect {
				t.Errorf("got %v; expect %v", result, test.expect)
			}
		})
	}
}

func TestConvInt(t *testing.T) {
	type Test struct {
		name   string
		value  string
		expect int
	}

	tests := []Test{
		{name: "ConvInt() shoud be 1", value: "1", expect: 1},
		{name: "ConvInt() shoud be 200", value: "200", expect: 200},
		{name: "ConvInt() shoud be 3000", value: "3000", expect: 3000},
		{name: "ConvInt() shoud be -1", value: "-1", expect: -1},
		{name: "ConvInt() shoud be -200", value: "-200", expect: -200},
		{name: "ConvInt() shoud be -3000", value: "-3000", expect: -3000},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Setenv("TESTING", test.value)
			env := EnvironmentVariable("TESTING")
			var setting Setting[int]
			result := ConvInt(env)(setting).Value
			if result != test.expect {
				t.Errorf("got %v; expect %v", result, test.expect)
			}
		})
	}
}

func TestConvInt8(t *testing.T) {
	type Test struct {
		name   string
		value  string
		expect int8
	}

	tests := []Test{
		{name: "ConvInt8() shoud be 1", value: "1", expect: 1},
		{name: "ConvInt8() shoud be 100", value: "100", expect: 100},
		{name: "ConvInt8() shoud be 127", value: "127", expect: 127},
		{name: "ConvInt8() shoud be -1", value: "-1", expect: -1},
		{name: "ConvInt8() shoud be -100", value: "-100", expect: -100},
		{name: "ConvInt8() shoud be -128", value: "-128", expect: -128},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Setenv("TESTING", test.value)
			env := EnvironmentVariable("TESTING")
			var setting Setting[int8]
			result := ConvInt8(env)(setting).Value
			if result != test.expect {
				t.Errorf("got %v; expect %v", result, test.expect)
			}
		})
	}
}

func TestConvInt16(t *testing.T) {
	type Test struct {
		name   string
		value  string
		expect int16
	}

	tests := []Test{
		{name: "ConvInt16() shoud be 1", value: "1", expect: 1},
		{name: "ConvInt16() shoud be 100", value: "100", expect: 100},
		{name: "ConvInt16() shoud be 127", value: "127", expect: 127},
		{name: "ConvInt16() shoud be -1", value: "-1", expect: -1},
		{name: "ConvInt16() shoud be -100", value: "-100", expect: -100},
		{name: "ConvInt16() shoud be -128", value: "-128", expect: -128},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Setenv("TESTING", test.value)
			env := EnvironmentVariable("TESTING")
			var setting Setting[int16]
			result := ConvInt16(env)(setting).Value
			if result != test.expect {
				t.Errorf("got %v; expect %v", result, test.expect)
			}
		})
	}
}

func TestConvInt32(t *testing.T) {
	type Test struct {
		name   string
		value  string
		expect int32
	}

	tests := []Test{
		{name: "ConvInt32() shoud be 1", value: "1", expect: 1},
		{name: "ConvInt32() shoud be 100", value: "100", expect: 100},
		{name: "ConvInt32() shoud be 127", value: "127", expect: 127},
		{name: "ConvInt32() shoud be -1", value: "-1", expect: -1},
		{name: "ConvInt32() shoud be -100", value: "-100", expect: -100},
		{name: "ConvInt32() shoud be -128", value: "-128", expect: -128},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Setenv("TESTING", test.value)
			env := EnvironmentVariable("TESTING")
			var setting Setting[int32]
			result := ConvInt32(env)(setting).Value
			if result != test.expect {
				t.Errorf("got %v; expect %v", result, test.expect)
			}
		})
	}
}

func TestConvInt64(t *testing.T) {
	type Test struct {
		name   string
		value  string
		expect int64
	}

	tests := []Test{
		{name: "ConvInt64() shoud be 1", value: "1", expect: 1},
		{name: "ConvInt64() shoud be 100", value: "100", expect: 100},
		{name: "ConvInt64() shoud be 127", value: "127", expect: 127},
		{name: "ConvInt64() shoud be -1", value: "-1", expect: -1},
		{name: "ConvInt64() shoud be -100", value: "-100", expect: -100},
		{name: "ConvInt64() shoud be -128", value: "-128", expect: -128},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Setenv("TESTING", test.value)
			env := EnvironmentVariable("TESTING")
			var setting Setting[int64]
			result := ConvInt64(env)(setting).Value
			if result != test.expect {
				t.Errorf("got %v; expect %v", result, test.expect)
			}
		})
	}
}

func TestConvLevel(t *testing.T) {
	type Test struct {
		name   string
		value  string
		expect slog.Level
	}

	tests := []Test{
		{name: "ConvLevel() shoud be debug", value: "debug", expect: slog.LevelDebug},
		{name: "ConvLevel() shoud be info", value: "INFO", expect: slog.LevelInfo},
		{name: "ConvLevel() shoud be warn", value: "WArn", expect: slog.LevelWarn},
		{name: "ConvLevel() shoud be error", value: "error", expect: slog.LevelError},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Setenv("TESTING", test.value)
			env := EnvironmentVariable("TESTING")
			var setting Setting[slog.Level]
			result := ConvLevel(env)(setting).Value
			if result != test.expect {
				t.Errorf("got %v; expect %v", result, test.expect)
			}
		})
	}
}

func TestConvString(t *testing.T) {
	type Test struct {
		name   string
		value  string
		expect string
	}

	tests := []Test{
		{name: "ConvString() shoud be \"how\"", value: "how", expect: "how"},
		{name: "ConvString() shoud be \"much\"", value: "much", expect: "much"},
		{name: "ConvString() shoud be \"wood\"", value: "wood", expect: "wood"},
		{name: "ConvString() shoud be \"would\"", value: "would", expect: "would"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Setenv("TESTING", test.value)
			env := EnvironmentVariable("TESTING")
			var setting Setting[string]
			result := ConvString(env, false)(setting).Value
			if result != test.expect {
				t.Errorf("got %v; expect %v", result, test.expect)
			}
		})
	}
}

func TestConvStringSlice(t *testing.T) {
	type Test struct {
		name   string
		value  string
		expect []string
	}

	tests := []Test{
		{name: "ConvString() shoud be [a, b, c]", value: "a,b,c", expect: []string{"a", "b", "c"}},
		{name: "ConvString() shoud be ", value: "one,two,three", expect: []string{"one", "two", "three"}},
		{name: "ConvString() shoud be ", value: "dog,cat,fish", expect: []string{"dog", "cat", "fish"}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Setenv("TESTING", test.value)
			env := EnvironmentVariable("TESTING")
			var setting Setting[[]string]
			result := ConvStringSlice(env, ",", false)(setting).Value
			if slices.Compare(result, test.expect) != 0 {
				t.Errorf("got %v; expect %v", result, test.expect)
			}
		})
	}
}

func TestConvUint(t *testing.T) {
	type Test struct {
		name   string
		value  string
		expect uint
	}

	tests := []Test{
		{name: "ConvUint() shoud be 1", value: "1", expect: 1},
		{name: "ConvUint() shoud be 200", value: "200", expect: 200},
		{name: "ConvUint() shoud be 3000", value: "3000", expect: 3000},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Setenv("TESTING", test.value)
			env := EnvironmentVariable("TESTING")
			var setting Setting[uint]
			result := ConvUint(env)(setting).Value
			if result != test.expect {
				t.Errorf("got %v; expect %v", result, test.expect)
			}
		})
	}
}

func TestConvUint8(t *testing.T) {
	type Test struct {
		name   string
		value  string
		expect uint8
	}

	tests := []Test{
		{name: "ConvUint8() shoud be 1", value: "1", expect: 1},
		{name: "ConvUint8() shoud be 200", value: "200", expect: 200},
		{name: "ConvUint8() shoud be 3000", value: "255", expect: 255},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Setenv("TESTING", test.value)
			env := EnvironmentVariable("TESTING")
			var setting Setting[uint8]
			result := ConvUint8(env)(setting).Value
			if result != test.expect {
				t.Errorf("got %v; expect %v", result, test.expect)
			}
		})
	}
}

func TestConvUint16(t *testing.T) {
	type Test struct {
		name   string
		value  string
		expect uint16
	}

	tests := []Test{
		{name: "ConvUint16() shoud be 1", value: "1", expect: 1},
		{name: "ConvUint16() shoud be 200", value: "200", expect: 200},
		{name: "ConvUint16() shoud be 3000", value: "3000", expect: 3000},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Setenv("TESTING", test.value)
			env := EnvironmentVariable("TESTING")
			var setting Setting[uint16]
			result := ConvUint16(env)(setting).Value
			if result != test.expect {
				t.Errorf("got %v; expect %v", result, test.expect)
			}
		})
	}
}

func TestConvUint32(t *testing.T) {
	type Test struct {
		name   string
		value  string
		expect uint32
	}

	tests := []Test{
		{name: "ConvUint32() shoud be 1", value: "1", expect: 1},
		{name: "ConvUint32() shoud be 200", value: "200", expect: 200},
		{name: "ConvUint32() shoud be 3000", value: "3000", expect: 3000},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Setenv("TESTING", test.value)
			env := EnvironmentVariable("TESTING")
			var setting Setting[uint32]
			result := ConvUint32(env)(setting).Value
			if result != test.expect {
				t.Errorf("got %v; expect %v", result, test.expect)
			}
		})
	}
}

func TestConvUint64(t *testing.T) {
	type Test struct {
		name   string
		value  string
		expect uint64
	}

	tests := []Test{
		{name: "ConvUint64() shoud be 1", value: "1", expect: 1},
		{name: "ConvUint64() shoud be 200", value: "200", expect: 200},
		{name: "ConvUint64() shoud be 3000", value: "3000", expect: 3000},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Setenv("TESTING", test.value)
			env := EnvironmentVariable("TESTING")
			var setting Setting[uint64]
			result := ConvUint64(env)(setting).Value
			if result != test.expect {
				t.Errorf("got %v; expect %v", result, test.expect)
			}
		})
	}
}

func TestConvURL(t *testing.T) {
	type Test struct {
		name   string
		value  string
		expect *url.URL
	}

	google, _ := url.Parse("https://www.google.com")
	apple, _ := url.Parse("https://www.apple.com")
	microsoft, _ := url.Parse("https://www.microsoft.com")
	tests := []Test{
		{name: "ConvURL() shoud be https://www.google.com", value: "https://www.google.com", expect: google},
		{name: "ConvURL() shoud be https://www.apple.com", value: "https://www.apple.com", expect: apple},
		{name: "ConvURL() shoud be https://www.microsoft.com", value: "https://www.microsoft.com", expect: microsoft},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Setenv("TESTING", test.value)
			env := EnvironmentVariable("TESTING")
			var setting Setting[*url.URL]
			result := ConvURL(env)(setting).Value
			if result.String() != test.expect.String() {
				t.Errorf("got %v; expect %v", result, test.expect)
			}
		})
	}
}
