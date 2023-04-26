package envjson

import "testing"

func TestDefaultLoadJSOND(t *testing.T) {
	result, err := loadJSON("testdata/env.json")
	if err != nil {
		t.Error(err)
	}
	resultA := result["A"]
	resultB := result["B"]
	if resultA != "got A" || resultB != "got B" {
		t.Error("properties are not loaded correctly")
	}
}

func TestLoadJSONWithProp(t *testing.T) {
	result, err := loadJSON("testdata/env3.json", "C")
	if err != nil {
		t.Error(err)
	}
	resultD := result["D"]
	resultE := result["E"]
	if resultD != "got D" || resultE != "got E" {
		t.Error("properties are not loaded correctly")
	}
}

func TestLoadJSONWrongType(t *testing.T) {
	_, err := loadJSON("testdata/env3_err.json", "C")
	if err == nil {
		t.Errorf("non-string property must cause an error")
	}
}
