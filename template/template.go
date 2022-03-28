package template

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Route struct {
	Package       string   `yaml:"package"`
	InputClasses  []string `yaml:"input_classes"`
	OutputClasses []string `yaml:"output_classes"`
}

type Template struct {
	Routes map[string]Route `yaml:"routes"`
}

func InitFile(path string) error {
	t := Template{
		Routes: map[string]Route{
			"/": {
				Package:       "example_package",
				InputClasses:  []string{"example_input"},
				OutputClasses: []string{"example_output"},
			},
		},
	}
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	encoder := yaml.NewEncoder(f)
	if err := encoder.Encode(&t); err != nil {
		return err
	}
	return nil
}

func LoadFile(path string) (*Template, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	decoder := yaml.NewDecoder(f)
	t := Template{}
	if err := decoder.Decode(&t); err != nil {
		panic(err)
	}
	return &t, nil
}
