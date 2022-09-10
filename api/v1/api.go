package api

import (
	"fmt"
)

type TypeMeta struct {
	Version string `yaml:"version"`
	Kind    string `yaml:"kind"`
}

func (m TypeMeta) Validate() error {
	if m.Version == "" {
		return fmt.Errorf("version must be set")
	}

	if m.Kind == "" {
		return fmt.Errorf("kind must be set")
	}

	return nil
}

type ObjectMeta struct {
	Name      string            `yaml:"name,omitempty"`
	Namespace string            `yaml:"namespace,omitempty"`
	Labels    map[string]string `yaml:"labels,omitempty"`
}

func (m ObjectMeta) Validate() error {
	if m.Name == "" {
		return fmt.Errorf("name must be set")
	}

	if m.Namespace == "" {
		return fmt.Errorf("namespace must be set")
	}

	return nil
}

// scaling, limit, offset, transfer function
type SLOT struct {
	Min    float64 `yaml:"min"`
	Max    float64 `yaml:"max"`
	Offset float64 `yaml:"offset"`
	Size   uint8   `yaml:"size"`
	Unit   string  `yaml:"unit,omitempty"`
}

func (s SLOT) Validate() error {
	if s.Size < 1 || s.Size > 64 {
		return fmt.Errorf("size must be between 0 and 64, was: %d", s.Size)
	}

	return nil
}

type Identifier struct {
	Standard uint16 `yaml:"standard,omitempty"`
	Extended uint32 `yaml:"extended,omitempty"`
	J1939    struct {
		Priority         uint8 `yaml:"priority,omitempty"`
		DataPage         bool  `yaml:"data_page,omitempty"`
		ExtendedDataPage bool  `yaml:"extended_data_page,omitempty"`
		PDUFormat        uint8 `yaml:"pdu_format,omitempty"`
		PDUSpecific      uint8 `yaml:"pdu_specific,omitempty"`
		SourceAddress    uint8 `yaml:"source_address,omitempty"`
	} `yaml:"j1939,omitempty"`
}

func (i Identifier) Validate() error {
	if i.Standard > 2^11-1 {
		return fmt.Errorf("standard identifier too large")
	}

	if i.Extended > 2^29-1 {
		return fmt.Errorf("extended identifier too large")
	}

	return nil
}
