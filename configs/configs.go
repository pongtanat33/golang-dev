package configs

type Configs struct {
	MSSQL    MSSQL
	App      Fiber
}

type Fiber struct {
	Host string
	Port string
}

// Database
type MSSQL struct {
	Host        string
	Port          string
	Username      string
	Password      string
	Database      string
	ConnectionTimeout int
	Encrypt       bool
}