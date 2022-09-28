package api

import (
	"testing"
)

func TestResourceValidate(t *testing.T) {
	m := Resource{}

	if m.Validate() == nil {
		t.Errorf("expected validate to fail with empty values")
	}
}

func TestMetadataValidate(t *testing.T) {
	m := Metadata{}

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

func TestIdentifierValidate(t *testing.T) {
	i := Identifier{}

	if i.Validate() == nil {
		t.Errorf("expected failure")
	}

	i.Standard = new(uint16)
	*i.Standard = 0
	if i.Validate() != nil {
		t.Errorf("expected success")
	}
}
