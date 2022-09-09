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
