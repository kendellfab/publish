package domain

import (
	"fmt"
	"log"
	"os"
	"strconv"
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
	Email       *ConfigEmail  `toml:"email"`
}

func (config *Config) GetSessionKeys() [][]byte {
	keys := make([][]byte, 0)
	for _, element := range config.SessionKeys {
		keys = append(keys, []byte(element))
	}
	return keys
}

func EmailConfigEnvironmentOverride(prefix string, ce *ConfigEmail) *ConfigEmail {
	if ce == nil {
		ce = &ConfigEmail{}
	}
	if host := os.Getenv(prefix + EMAIL_HOST); host != "" {
		ce.Host = host
	}
	if portStr := os.Getenv(prefix + EMAIL_PORT); portStr != "" {
		if port, pErr := strconv.Atoi(portStr); pErr == nil {
			ce.Port = port
		} else {
			log.Fatal(pErr)
		}
	}
	if username := os.Getenv(prefix + EMAIL_USERNAME); username != "" {
		ce.Username = username
	}
	if password := os.Getenv(prefix + EMAIL_PASSWORD); password != "" {
		ce.Password = password
	}
	if from := os.Getenv(prefix + EMAIL_FROM); from != "" {
		ce.From = from
	}
	return ce
}

type ConfigEmail struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	Username string `toml:"username"`
	Password string `toml:"password"`
	From     string `toml:"from"`
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
	MaxIdle  int    `toml:"max_idle"`
	MaxOpen  int    `tomel:"max_open"`
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
