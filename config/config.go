package config

type Config struct {
	Global             Global
	Include            []string
	Init               any
	E2eTests           any
	Purge              Cmd
	Services           map[string]any
	AdditionalCommands []Cmd
}

type E2eTests struct {
	Cmd string
}

func LoadConfig() Config {
	_ = LoadRawConfig()
	// parse raw config

	var conf = Config{}
	return conf
}
