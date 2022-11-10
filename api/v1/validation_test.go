package api

import (
	"testing"
)

func TestObjectValidate(t *testing.T) {
	o := Object{}

	if o.Validate() == nil {
		t.Errorf("expected validate to fail with empty data")
	}
}

func TestMetadataValidate(t *testing.T) {
	m := Metadata{}

	if m.Validate() == nil {
		t.Errorf("expected validate to fail with empty data")
	}
}
