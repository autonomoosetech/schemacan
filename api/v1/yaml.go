package api

import (
	"gopkg.in/yaml.v3"
)

func (o *Object) UnmarshalYAML(n *yaml.Node) error {
	type O Object
	type T struct {
		*O   `yaml:",inline"`
		Spec yaml.Node `yaml:"spec"`
	}

	obj := &T{O: (*O)(o)}
	err := n.Decode(obj)
	if err != nil {
		return err
	}

	switch o.Type {
	case "device":
		o.Spec = new(Device)
	case "message":
		o.Spec = new(Message)
	case "slot":
		o.Spec = new(Slot)
	default:
		panic("type unknown")
	}

	return obj.Spec.Decode(o.Spec)
}
