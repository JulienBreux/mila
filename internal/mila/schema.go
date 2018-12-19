package mila

// Schema represents a Mila schema
type Schema struct {
	Version string
	Vars    Vars
}

// UnmarshalYAML implements yaml.Unmarshaler interface
func (sc *Schema) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var schema struct {
		Version string
		Vars    Vars
	}
	if err := unmarshal(&schema); err != nil {
		return err
	}
	sc.Version = schema.Version
	sc.Vars = schema.Vars

	return nil
}
