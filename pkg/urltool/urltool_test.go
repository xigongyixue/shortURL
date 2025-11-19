package urltool

import (
	"testing"
)

func TestGetBasePath(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		want     string
		hasError bool
	}{
		{
			name:     "normal URL",
			input:    "https://example.com/path/to/resource",
			want:     "resource",
			hasError: false,
		},
		{
			name:     "empty path",
			input:    "https://example.com",
			want:     ".",
			hasError: false,
		},
		{
			name:     "invalid URL",
			input:    "://invalid-url",
			want:     "",
			hasError: true,
		},
		{
			name:     "URL with query params",
			input:    "https://example.com/path?query=value",
			want:     "path",
			hasError: false,
		},
		{
			name:     "URL with fragment",
			input:    "https://example.com/path#section",
			want:     "path",
			hasError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetBasePath(tt.input)
			if (err != nil) != tt.hasError {
				t.Errorf("GetBasePath() error = %v, hasError %v", err, tt.hasError)
				return
			}
			if got != tt.want {
				t.Errorf("GetBasePath() = %v, want %v", got, tt.want)
			}
		})
	}
}