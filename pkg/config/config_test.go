package config

import (
	"testing"
)

type test struct {
	IntVal int      `json:"intVal"`
	StrVal string   `json:"strVal"`
	ArrVal []string `json:"arrVal"`
}

func TestGetLocalConfig(t *testing.T) {
	v, err := GetLocalConfig("../../test/testdata/config/test.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	test := &test{}
	err = v.UnmarshalKey("testField", &test)
	if err != nil {
		t.Error(err)
		return
	}
	if test.IntVal != 11 || test.StrVal != "hello" ||
		len(test.ArrVal) != 3 ||
		test.ArrVal[0] != "item1" || test.ArrVal[1] != "item2" || test.ArrVal[2] != "item3" {
		t.Error("value not match")
		return
	}
}
