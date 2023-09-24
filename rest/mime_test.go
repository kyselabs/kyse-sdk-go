package rest

import (
	"testing"
)

func TestAcceptWithVersion(t *testing.T) {
	m := NewMime("1.0")
	result := m.Accept()

	expected := "application/vnd.kyse+json; version=1.0"
	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}

func TestAcceptWithoutVersion(t *testing.T) {
	m := NewMime("")
	result := m.Accept()

	expected := "application/json"
	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}
