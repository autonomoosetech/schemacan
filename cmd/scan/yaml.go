package main

import (
	"bytes"

	"github.com/autonomoosetech/schemacan/api/v1"
	"gopkg.in/yaml.v2"
)

func splitStream(in []byte) (out [][]byte) {
	return bytes.Split(in, []byte("---\n"))
}

type Descriptor struct {
	Version *string `yaml:"version,omitempty"`
	Kind    *string `yaml:"kind,omitempty"`
}

func Unmarshal(in []byte) (object []interface{}, err error) {
	for _, slice := range splitStream(in) {
		var resource Descriptor

		err = yaml.Unmarshal(slice, &resource)
		if err != nil {
			return nil, err
		}

		switch *resource.Kind {
		case "slot":
			var slot *api.SLOTResource
			yaml.Unmarshal(slice, &slot)
			object = append(object, slot)
		case "message":
			var slot *api.MessageResource
			yaml.Unmarshal(slice, &slot)
			object = append(object, slot)
		case "device":
			var slot *api.DeviceResource
			yaml.Unmarshal(slice, &slot)
			object = append(object, slot)
		}
	}

	return object, err
}
