package utils

import (
	"testing"
)

func TestTrimSpace(t *testing.T) {
	result := TrimSpace(" hello world ")
	if result != "hello world" {
		t.Errorf("ожидалось 'hello world', получено  '%s'", result)
	}
}
