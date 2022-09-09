package api

import (
	"testing"
)

func TestTypeMetaValidate(t *testing.T) {
	m := TypeMeta{}

	if m.Validate() == nil {
		t.Errorf("expected validate to fail with empty values")
	}
}

func TestObjectMetaValidate(t *testing.T) {
	m := ObjectMeta{}

	if m.Validate() == nil {
		t.Errorf("expected validate to fail with empty values")
	}
}
