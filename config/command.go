package config

type Cmd struct {
	Cmd   string
	File  string
	Shell string
}
type TestCmd struct {
	Cmd       Cmd
	DependsOn []Service
}
