package config

type Cmd struct {
	Cmd   string
	File  string
	Shell string
}

func ParseCmd(a any) Cmd {
	cmd := Cmd{
		Cmd:   "",
		File:  "",
		Shell: "",
	}
	s, ok := a.(string)
	if ok {
		cmd.Cmd = s
		return cmd
	}

	panic("Invalid Cmd shape")
}

type TestCmd struct {
	Cmd       Cmd
	DependsOn []Service
}
