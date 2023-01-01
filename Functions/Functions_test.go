package main_test

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	args1 := []string{"asitha", "divisekara"}
	args2 := []string{"asitha", "23"}

	tests := []struct {
		name    string
		args    []string
		wantErr bool
	}{
		{
			name:    "test1",
			args:    args1,
			wantErr: false,
		},
		{
			name:    "test2",
			args:    args2,
			wantErr: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			fmt.Println(test)
		})
	}
}
