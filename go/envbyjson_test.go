package envbyjson

import (
	"os"
	"testing"
)

func TestDefaultLoad(t *testing.T) {
	err := Load()
	if err != nil {
		t.Error(err)
	}
	resultA := os.Getenv("A")
	resultB := os.Getenv("B")
	if resultA != "got A" || resultB != "got B" {
		t.Error("properties are not loaded correctly")
	}
}

func TestLoadWrongType(t *testing.T) {
	err := Load("testdata/env_err.json")
	if err == nil {
		t.Errorf("non-string property must cause an error")
	}
}

func TestLoadGivenPath(t *testing.T) {
	err := Load("testdata/env2.json")
	if err != nil {
		t.Error(err)
	}
	resultA := os.Getenv("A")
	resultB := os.Getenv("B")
	if resultA != "got A" || resultB != "got B" {
		t.Error("properties are not loaded correctly")
	}
}

func TestLoadProp(t *testing.T) {
	err := LoadProp("testdata/env3.json", "C")
	if err != nil {
		t.Error(err)
	}
	resultD := os.Getenv("D")
	resultE := os.Getenv("E")
	if resultD != "got D" || resultE != "got E" {
		t.Error("properties are not loaded correctly")
	}
}

func TestLoadPropNoParam(t *testing.T) {
	err := LoadProp()
	if err == nil {
		t.Error("at least one param must be given")
	}
}

func TestLoadPropWrongType(t *testing.T) {
	err := LoadProp("testdata/env3_err.json", "C")
	if err == nil {
		t.Errorf("non-string property must cause an error")
	}
}
