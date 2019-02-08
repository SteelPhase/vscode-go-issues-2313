package main

import "testing"

func TestTargetForTesting(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "sanity",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			TargetForTesting()
		})
	}
}
