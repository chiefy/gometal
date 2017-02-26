package gometal

import (
	"testing"
)

func TestNewGometal(t *testing.T) {
	if _, err := NewGometal("isnotthere"); err == nil {
		t.Errorf("Directory does not exist but no error was returned!")
	}
	if _, err := NewGometal("../example/src"); err != nil {
		t.Errorf("Directory exists but error was returned!")
	}
}

func TestUse(t *testing.T) {

}

func TestBuild(t *testing.T) {

}
