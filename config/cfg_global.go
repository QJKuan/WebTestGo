package config

var (
	SER_PORT  string
	GBL_UPMEM int64
	DB_DNS    string
)

func SetGbl(conf Cfg) {
	SER_PORT = conf.Server.Port
	GBL_UPMEM = conf.Server.UploadMem
	DB_DNS = conf.Server.Dns
}
