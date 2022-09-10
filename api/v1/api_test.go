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

func TestSLOTValidate(t *testing.T) {
	s := SLOT{}

	if s.Validate() == nil {
		t.Errorf("expected validate to fail with empty values")
	}

	s.Size = 64
	if s.Validate() != nil {
		t.Errorf("expected validate to pass with size within bounds")
	}

	s.Size = 65
	if s.Validate() == nil {
		t.Errorf("expected validate to fail with size outside of bounds")
	}
}
