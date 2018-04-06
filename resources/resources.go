package resources

import (
	"strconv"

	yaml "gopkg.in/yaml.v2"
)

// Resource hold configure of all resources
type Resource struct {
	Type       string      `yaml:"type"`
	Properties interface{} `yaml:"properties"`
}

type element struct {
	GetParam    string `yaml:"get_param"`
	GetResource string `yaml:"get_resource"`
}

type elementWrapper struct {
	String string
	Bool   bool
	Int    int
	Type   string
	element
}

func (e *elementWrapper) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var String string
	var Bool bool
	var Int int
	var el element

	if err := unmarshal(&String); err == nil {
		e.String = String
		e.Type = "string"
		return nil
	}

	if err := unmarshal(&Bool); err == nil {
		e.Bool = Bool
		e.Type = "bool"
		return nil
	}

	if err := unmarshal(&Int); err == nil {
		e.Int = Int
		e.Type = "int"
		return nil
	}

	if err := unmarshal(&el); err != nil {
		if _, ok := err.(*yaml.TypeError); !ok {
			return err
		}
	}

	e.Type = "struct"
	e.element.GetParam = el.GetParam
	e.element.GetResource = el.GetResource
	return nil
}

func (e elementWrapper) MarshalYAML() (interface{}, error) {
	switch e.Type {
	case "string":
		return []byte(e.String), nil
	case "int":
		return []byte(strconv.Itoa(e.Int)), nil
	case "bool":
		return []byte(strconv.FormatBool(e.Bool)), nil
	default:
		return yaml.Marshal(e.element)
	}
}
