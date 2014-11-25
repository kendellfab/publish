package domain

type Config struct {
	Port        int      `toml:"port"`
	SiteName    string   `toml:"site"`
	AdminDir    string   `toml:"admin"`
	ThemeDir    string   `toml:"theme"`
	Sqlite      string   `toml:"sqlite"`
	SessionKeys []string `toml:"session"`
	UploadDir   string   `toml:"upload"`
}

func (config *Config) GetSessionKeys() [][]byte {
	keys := make([][]byte, 0)
	for _, element := range config.SessionKeys {
		keys = append(keys, []byte(element))
	}
	return keys
}
