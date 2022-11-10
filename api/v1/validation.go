package api

import (
	"errors"
	"fmt"
)

func (o Object) Validate() error {
	if o.Type == "" {
		return errors.New("field 'Type' cannot be empty")
	}

	if o.Version == "" {
		return errors.New("field 'Version' cannot be empty")
	}

	if err := o.Metadata.Validate(); err != nil {
		return fmt.Errorf("metadata validation failed: %s", err)
	}

	return nil
}

func (m Metadata) Validate() error {
	if m.Name == "" {
		return errors.New("field 'Metadata.Name' cannot be empty")
	}

	return nil
}
