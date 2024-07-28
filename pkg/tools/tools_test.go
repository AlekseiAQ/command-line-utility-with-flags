package tools_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/AlekseiAQ/command-line-utility-with-flags/pkg/tools"
)

func TestFormatSize(t *testing.T) {
	tests := []struct {
		name     string
		expected string
		size     int64
		isHuman  bool
	}{
		{
			name:     "bytes",
			size:     100,
			isHuman:  true,
			expected: "100 B",
		},
		{
			name:     "kilobytes",
			size:     1024,
			isHuman:  true,
			expected: "1.0 KB",
		},
		{
			name:     "megabytes",
			size:     1024 * 1024,
			isHuman:  true,
			expected: "1.0 MB",
		},
		{
			name:     "gigabytes",
			size:     1024 * 1024 * 1024,
			isHuman:  true,
			expected: "1.0 GB",
		},
		{
			name:     "terabytes",
			size:     1024 * 1024 * 1024 * 1024,
			isHuman:  true,
			expected: "1.0 TB",
		},
		{
			name:     "petabytes",
			size:     1024 * 1024 * 1024 * 1024 * 1024,
			isHuman:  true,
			expected: "1.0 PB",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tools.FormatSize(tt.size, tt.isHuman))
		})
	}
}
