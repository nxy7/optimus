package config

type Config struct {
	Global Global
	// Include            []string
	Init               Cmd
	E2eTests           Cmd
	Purge              Cmd
	Services           map[string]any
	AdditionalCommands []Cmd
}

type E2eTests struct {
	Cmd string
}

func LoadConfig() Config {
	var conf = DefaultConfig()
	// raw := LoadRawConfig()
	// if raw.Global.(map[string]any) {
	// 	conf.Global = ParseGlobal(raw.Global)
	// }
	// init := ParseCmd(raw.Init)
	// global := raw.
	// parse raw config

	return conf
}

func DefaultConfig() Config {
	return Config{
		Global: Global{
			ShellCmd: "bash -c",
		},
		Init:               Cmd{},
		E2eTests:           Cmd{},
		Purge:              Cmd{},
		Services:           map[string]any{},
		AdditionalCommands: []Cmd{},
	}
}
