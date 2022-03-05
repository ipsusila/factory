package factory

// ObjectConfig stores factory configuration
type ObjectConfig struct {
	Name    string  `json:"name" toml:"name" yaml:"name" xml:"name"`
	Options Options `json:"options" toml:"options" yaml:"options" xml:"options"`
}
