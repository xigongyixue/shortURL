package base62

import (
	"testing"
)

func TestInt2String(t *testing.T) {
	tests := []struct {
		name     string
		input    uint64
		expected string
	}{
		// 边界值：0
		{"zero", 0, "0"},
		
		// 最小值边界：1
		{"minimum", 1, "1"},
		
		// 正常值：小于62的值（单字符）
		{"single_char_10", 10, "a"},
		{"single_char_35", 35, "z"},
		{"single_char_36", 36, "A"},
		{"single_char_61", 61, "Z"},
		
		// 正常值：大于62的值（多字符）
		{"two_chars_62", 62, "10"},
		{"two_chars_63", 63, "11"},
		{"three_chars_6347", 6347, "1En"},
		
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Int2String(tt.input)
			if result != tt.expected {
				t.Errorf("Int2String(%d) = %s, expected %s", tt.input, result, tt.expected)
			}
		})
	}
}

func TestString2Int(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    uint64
		wantErr bool
	}{
		{
			name:  "single digit",
			input: "a",
			want:  10,
		},
		{
			name:  "three digits",
			input: "1En",
			want:  6347,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := String2Int(tt.input)
			if got != tt.want {
				t.Errorf("String2Int() = %v, want %v", got, tt.want)
			}
		})
	}
}
