package config

type Config struct {
	Back back `toml:"back"`
	Mail mail `toml:"mail"`
}

type back struct {
	Folders []string `toml:"folders"`
	Files   []string `toml:"files"`
}

type mail struct {
	To       []string `toml:"to"`
	Host     string   `toml:"host"`
	Port     int      `toml:"port"`
	User     string   `toml:"user"`
	Password string   `toml:"password"`
}
