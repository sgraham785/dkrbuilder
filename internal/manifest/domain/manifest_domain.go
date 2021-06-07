package domain

import (
	"bytes"

	"gopkg.in/yaml.v2"
)

type BaseOSImages struct {
	Versions   map[string]BaseOSVersions
	Dependents []Dependents
}

type BaseOSVersions struct {
	Build ImageBuild
	Child map[string]BaseChildImages
}

type BaseChildImages struct {
	Versions []string
}

type Images struct {
	ImageNames map[string]BaseOSImages
}

type ImageBuild struct {
	Context    string
	Dockerfile string
	Args       map[string]string
}

type Dependents struct {
	Repo   string
	Branch string
}

func (v *Images) Unmarshal(data []byte) error {
	err := yaml.NewDecoder(bytes.NewReader(data)).Decode(v)
	if err != nil {
		return err
	}
	// for k, v := range v {
	v = v
	// }
	return nil
}

func (i *Images) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var base map[string]BaseOSImages
	if err := unmarshal(&base); err != nil {
		// Here we expect an error because a boolean cannot be converted to a
		// a MajorVersion
		if _, ok := err.(*yaml.TypeError); !ok {
			return err
		}
	}
	i.ImageNames = base
	return nil
}

func (i *BaseOSImages) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var baseVersions map[string]BaseOSVersions
	var params struct {
		Dependents []Dependents
	}
	if err := unmarshal(&params); err != nil {
		return err
	}
	if err := unmarshal(&baseVersions); err != nil {
		// Here we expect an error because a boolean cannot be converted to a
		// a MajorVersion
		if _, ok := err.(*yaml.TypeError); !ok {
			return err
		}
	}
	i.Versions = baseVersions
	i.Dependents = params.Dependents
	return nil
}

func (i *BaseOSVersions) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var baseChildImages map[string]BaseChildImages
	var params struct {
		Build ImageBuild
	}
	if err := unmarshal(&params); err != nil {
		return err
	}
	if err := unmarshal(&baseChildImages); err != nil {
		// Here we expect an error because a boolean cannot be converted to a
		// a MajorVersion
		if _, ok := err.(*yaml.TypeError); !ok {
			return err
		}
	}
	i.Build = params.Build
	i.Child = baseChildImages
	return nil
}

func (i *BaseChildImages) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var baseChildVersions []string
	if err := unmarshal(&baseChildVersions); err != nil {
		// Here we expect an error because a boolean cannot be converted to a
		// a MajorVersion
		if _, ok := err.(*yaml.TypeError); !ok {
			return err
		}
	}
	i.Versions = baseChildVersions
	return nil
}
