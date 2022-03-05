package factory

// ObjectConfig stores factory configuration
type ObjectConfig struct {
	Name         string `json:"name" toml:"name" yaml:"name"`
	ConfigSource string `json:"configSource" toml:"configSource" yaml:"configSource"`
}
