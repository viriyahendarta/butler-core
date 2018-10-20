package config

//Config holds all config
type Config struct {
	AuthSecretKey string     `json:"auth_secret_key"`
	Debug         bool       `json:"debug"`
	HTTPServer    HTTPServer `json:"http_server"`
	Databases     Databases  `json:"databases"`
}

//HTTPServer holds data needed for serving http
type HTTPServer struct {
	Port int `json:"port"`
}

//Databases holds data needed for connect to database
type Databases struct {
	CoreDatabase Database `json:"core"`
}

//Database holds database fields
type Database struct {
	Name                  string   `json:"name"`
	Driver                string   `json:"driver"`
	MasterURL             string   `json:"master_url"`
	SlaveURLs             []string `json:"slave_urls"`
	MaxOpenConnections    int      `json:"max_open_connections"`
	MaxIdleConnections    int      `json:"max_idle_connections"`
	ConnectionMaxLifetime int      `json:"connection_max_lifetime"`
}
