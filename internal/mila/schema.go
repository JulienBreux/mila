package mila

// Schema represents a Mila schema
type Schema struct {
	Version   string
	Variables Variables
}

// UnmarshalYAML implements yaml.Unmarshaler interface
func (sc *Schema) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var schema struct {
		Version   string
		Variables Variables
	}
	if err := unmarshal(&schema); err != nil {
		return err
	}
	sc.Version = schema.Version
	sc.Variables = schema.Variables

	return nil
}
