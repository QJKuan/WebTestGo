package config

// Server 配置文件中的 server 块
type Server struct {
	Port      string `yaml:"port"`
	UploadMem int64  `yaml:"uploadMem"`
}
