package config

type Global struct {
	ShellCmd string
}

func ParseGlobal(a any) Global {
	temp, ok := a.(map[string]any)
	if !ok {
		panic("Invalid Global Propety")
	}
	val := temp["shell_cmd"].(string)
	return Global{
		ShellCmd: val,
	}

}
