package domain

import (
	"fmt"
)

type Config struct {
	Port        int           `toml:"port"`
	SiteName    string        `toml:"site"`
	PerPage     int           `toml:"per_page"`
	AdminDir    string        `toml:"admin"`
	ThemeDir    string        `toml:"theme"`
	SessionKeys []string      `toml:"session"`
	UploadDir   string        `toml:"upload"`
	Sqlite      *ConfigSqlite `toml:"sqlite"`
	Mysql       *ConfigMysql  `toml:"mysql"`
}

func (config *Config) GetSessionKeys() [][]byte {
	keys := make([][]byte, 0)
	for _, element := range config.SessionKeys {
		keys = append(keys, []byte(element))
	}
	return keys
}

type ConfigSqlite struct {
	Path string `toml:"path"`
}

type ConfigMysql struct {
	Username string `toml:"username"`
	Password string `toml:"password"`
	Db       string `toml:"db"`
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
}

func (c *ConfigMysql) GetConnectionString() string {

	var host string
	if c.Host != "" && c.Port != 0 {
		host = fmt.Sprintf("%s:%d", c.Host, c.Port)
	} else if c.Host != "" && c.Port == 0 {
		host = c.Host
	}

	return fmt.Sprintf("%s:%s@%s/%s", c.Username, c.Password, host, c.Db)
}
