package config

import "testing"

func TestVersion(t *testing.T) {
	tt := []struct {
		version  string
		expected string
	}{
		{"", "unknown"},
		{"1.2.3", "1.2.3"},
	}

	for _, tc := range tt {
		version = tc.version
		v := Version()
		if tc.expected != v {
			t.Errorf("Unexpcted version: wanted %s got %s", tc.expected, v)
		}
	}
}
