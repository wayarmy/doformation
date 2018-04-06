package manifests

import (
	"io/ioutil"

	"github.com/doformation/doformation/resources"
	"github.com/pkg/errors"
	yaml "gopkg.in/yaml.v2"
)

// FormationStruct define base of yaml config file
type FormationStruct struct {
	Version    string               `yaml:"version"`
	Desc       string               `yaml:"description"`
	Parameters map[string]Parameter `yaml:"parameters"`
	Resources  map[string]Resource  `yaml:"resources"`
}

// Parameter format of parameter
type Parameter struct {
	Value   string
	Type    string `yaml:"type"`
	Default string `yaml:"default"`
}

// Resource define base configuration of any resources
type Resource struct {
	Type       string      `yaml:"type"`
	Properties interface{} `yaml:"properties"`
}

// Parser will parse yaml file into defined struct
func Parser(file string) (FormationStruct, error) {
	formFile, err := ioutil.ReadFile(file)
	if err != nil {
		return FormationStruct{}, errors.Wrap(err, "")
	}

	var f FormationStruct
	err = yaml.Unmarshal(formFile, &f)
	if err != nil {
		return FormationStruct{}, errors.Wrap(err, "")
	}

	return f, nil
}

// ClassifyResourceType will classify resource type into exactly type
func ClassifyResourceType(rType string, r interface{}) (interface{}, error) {
	switch rType {
	case "DO::Server::Droplet":
		rYAML, err := yaml.Marshal(r)
		if err != nil {
			return nil, err
		}

		var d resources.DoDroplet
		err = yaml.Unmarshal(rYAML, d)
		if err != nil {
			return nil, err
		}

		return d, nil
	default:
		return nil, errors.New("Type " + rType + " is not supported")
	}
}
