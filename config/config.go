package config

type Cfg struct {
	Server Server `yaml:"server"`
	Work   Work   `yaml:"work"`
}
