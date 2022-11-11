// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.11.0 DO NOT EDIT.
package api

import (
	"encoding/json"
	"fmt"
)

// Device defines model for device.
type Device struct {
	Messages *[]Message `json:"messages,omitempty"`
}

// Identifier defines model for identifier.
type Identifier struct {
	// 29-bit identifier
	Extended *float32 `json:"extended,omitempty"`

	// 11-bit identifier
	Standard *float32 `json:"standard,omitempty"`
}

// Messages define the details of a CAN Bus frame an the format of its data
type Message struct {
	// Data format
	Data       []Signal    `json:"data"`
	Identifier *Identifier `json:"identifier,omitempty"`
	Length     uint8       `json:"length"`
}

// Metadata defines model for metadata.
type Metadata struct {
	Labels    *Metadata_Labels `json:"labels,omitempty"`
	Name      string           `json:"name"`
	Namespace *string          `json:"namespace,omitempty"`
}

// Metadata_Labels defines model for Metadata.Labels.
type Metadata_Labels struct {
	AdditionalProperties map[string]string `json:"-"`
}

// Object defines model for object.
type Object struct {
	Metadata Metadata     `json:"metadata"`
	Spec     *interface{} `json:"spec,omitempty"`
	Type     string       `json:"type"`
	Version  string       `json:"version"`
}

// Signal defines model for signal.
type Signal struct {
	// Name of the signal
	Name string `json:"name"`

	// Number of padding bits between signals
	Padding *uint8 `json:"padding,omitempty"`

	// Reference to a J1939 SLOT
	Slot *string `json:"slot,omitempty"`

	// Number of bits the signal occupies
	Type *string `json:"type,omitempty"`
}

// Slots are a J1939 construct that define how real values are encoded and decoded to and from fixed-point integers.
type Slot struct {
	// Maximum possible value
	Max *float32 `json:"max,omitempty"`

	// Minimum possible value
	Min *float32 `json:"min,omitempty"`

	// Offset of value
	Offset *float32 `json:"offset,omitempty"`

	// Number of bits the signal occupiess
	Size *uint8 `json:"size,omitempty"`

	// Unit of measurement
	Unit *string `json:"unit,omitempty"`
}

// Getter for additional properties for Metadata_Labels. Returns the specified
// element and whether it was found
func (a Metadata_Labels) Get(fieldName string) (value string, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for Metadata_Labels
func (a *Metadata_Labels) Set(fieldName string, value string) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]string)
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for Metadata_Labels to handle AdditionalProperties
func (a *Metadata_Labels) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]string)
		for fieldName, fieldBuf := range object {
			var fieldVal string
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshaling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for Metadata_Labels to handle AdditionalProperties
func (a Metadata_Labels) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}
