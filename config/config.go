package config

type Config struct {
	Global             Global
	BuildCmd           Cmd
	PushCmd            Cmd
	E2eTests           Cmd
	Services           map[string]Service
	AdditionalCommands map[string]Cmd
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
		Global:   Global{ShellCmd: "bash -c"},
		BuildCmd: Cmd{},
		PushCmd:  Cmd{},
		E2eTests: Cmd{},
		Services: map[string]Service{
			"test": Service{},
		},
		AdditionalCommands: map[string]Cmd{
			"testcmd": Cmd{
				Description: "Gówno",
				Run:         "echo 2",
			},
		},
	}
}
