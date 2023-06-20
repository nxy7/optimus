package config

type Service struct {
	Name    string
	Root    string
	Start   *Cmd
	Dev     *Cmd
	PostDev *Cmd
	Test    *TestCmd
	DirHash string
}
