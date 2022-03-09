package factory

// Config stores factory configuration
type Config struct {
	Name    string  `json:"name" toml:"name" yaml:"name" xml:"name"`
	Options Options `json:"options" toml:"options" yaml:"options" xml:"options"`
}
