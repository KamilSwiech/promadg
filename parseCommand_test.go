package main

import (
	"testing"
)

func TestInitCmds(t *testing.T) {
	tests := []struct {
		osArgs []string
	}{
		{[]string{"get", "-context"}},
		{[]string{"set", "-context"}},
		{[]string{"get", "-data alerts"}},
	}
	_ = tests

	t.Errorf("TestFailed %s", "BOIIIII")
}
