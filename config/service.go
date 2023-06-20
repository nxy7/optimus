package config

type Service struct {
	Name    string `json:"name"`
	Root    string `json:"root"`
	Start   Cmd
	Dev     Cmd
	PostDev Cmd
	Test    TestCmd
	DirHash string
}
