package main

import "testing"

func TestSubtractTableDriven(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		want int
	}{
		{"positive numbers", 10, 5, 5},
		{"negative numbers", -1, -1, 0},
		{"result is negative", 5, 10, -5},
		{"minus zero", 7, 0, 7},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Subtract(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("Subtract(%d, %d) = %d; want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	t.Run("normal division", func(t *testing.T) {
		got, err := Divide(10, 2)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if got != 5 {
			t.Errorf("got %d; want 5", got)
		}
	})

	t.Run("division by zero", func(t *testing.T) {
		_, err := Divide(10, 0)
		if err == nil {
			t.Error("expected error for division by zero, but got nil")
		}
	})
}
