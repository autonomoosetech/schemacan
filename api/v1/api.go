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
