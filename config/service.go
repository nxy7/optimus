package config

type Service struct {
	Name string

	// in the future I'd like to be automatically get service version from package.json/cargo.toml and other files, so it can be used to tag images
	// Version            string
	// Language           string
	Root               string
	Start              Cmd
	Build              Cmd
	Dev                Cmd
	PostDev            Cmd
	AdditionalCommands map[string]Cmd
	Test               TestCmd
	DirHash            string
}
