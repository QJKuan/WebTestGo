package config

type Cfg struct {
	Server Server `yaml:"server"`
	Admin  Admin  `yaml:"admin"`
}
