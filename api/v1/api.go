package api

import (
	"fmt"
)

//
// Object metadata
//

type Metadata struct {
	Name      string            `yaml:"name,omitempty"`
	Namespace string            `yaml:"namespace,omitempty"`
	Labels    map[string]string `yaml:"labels,omitempty"`
}

func (m Metadata) Validate() error {
	if m.Name == "" {
		return fmt.Errorf("name must be set")
	}

	if m.Namespace == "" {
		return fmt.Errorf("namespace must be set")
	}

	return nil
}

//
// Resource
//

type Resource struct {
	Version  string      `yaml:"version"`
	Kind     string      `yaml:"kind"`
	Metadata Metadata    `yaml:"metadata"`
	Spec     interface{} `yaml:"spec"`
}

func (r Resource) Validate() error {
	if r.Version == "" {
		return fmt.Errorf("version must be set")
	}

	if r.Kind == "" {
		return fmt.Errorf("kind must be set")
	}

	err := r.Metadata.Validate()
	if err != nil {
		return err
	}

	switch v := r.Spec.(type) {
	case *SLOT:
		err = v.Validate()
	case *Message:
		err = v.Validate()
	default:
		err = fmt.Errorf("type used in resource spec is not valid: %T", v)
		return err
	}

	return nil
}

//
// Scaling, limit, offset, transfer function
//

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

//
// Identifier
//

type Identifier struct {
	Standard *uint16 `yaml:"standard,omitempty"`
	Extended *uint32 `yaml:"extended,omitempty"`
}

func (i Identifier) Validate() error {
	standardDef := i.Standard != nil
	extendedDef := i.Extended != nil

	if standardDef == extendedDef {
		return fmt.Errorf("at least one and not both identifier types must be defined")
	}

	if i.Standard != nil {

		if *i.Standard > 2^11-1 {
			return fmt.Errorf("standard identifier too large")
		}
	}

	if i.Extended != nil {
		if *i.Extended > 2^29-1 {
			return fmt.Errorf("extended identifier too large")
		}
	}

	return nil
}

//
// Message
//

type Message struct {
	Identifier Identifier `yaml:"id"`
	Length     *uint8     `yaml:"length"`
	Data       []struct {
		Name    *string `yaml:"name,omitempty"`
		Size    *string `yaml:"size,omitempty"`
		SLOT    *string `yaml:"slot,omitempty"`
		Padding *uint8  `yaml:"padding,omitempty"`
	} `yaml:"data" json:"data"`
}

func (m Message) Validate() error {
	err := m.Identifier.Validate()
	if err != nil {
		return err
	}

	return nil
}
