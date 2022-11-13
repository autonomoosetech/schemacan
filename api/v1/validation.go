package api

import (
	"errors"
	"fmt"
)

type object interface {
	Validate() error
}

func (o Object) Validate() error {
	if o.Type == "" {
		return errors.New("field 'Type' cannot be empty")
	}

	if o.Version == "" {
		return errors.New("field 'Version' cannot be empty")
	}

	err := o.Metadata.Validate()
	if err != nil {
		return fmt.Errorf("metadata validation failed: %s", err)
	}

	var obj object

	switch o.Type {
	case "device":
		obj = o.Spec.(*Device)
	case "message":
		obj = o.Spec.(*Message)
	case "slot":
		obj = o.Spec.(*Slot)
	default:
		return fmt.Errorf("object type not recognised: %v", o.Type)
	}

	obj.Validate()

	return nil
}

func (m Metadata) Validate() error {
	if m.Name == "" {
		return errors.New("field 'Metadata.Name' cannot be empty")
	}

	return nil
}

func (d Device) Validate() error {
	return nil
}

func (m Message) Validate() error {
	return nil
}

func (s Slot) Validate() error {
	return nil
}
