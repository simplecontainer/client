package configuration

type Configuration struct {
	Target      string `default:"development" json:"target"`
	Root        string `json:"root"`
	Environment *Environment
	Flags       Flags
}

type Flags struct {
	Context string
	Y       bool
}

type Environment struct {
	HOMEDIR    string
	OPTDIR     string
	PROJECTDIR string
	PROJECT    string
	CLIENTIP   string
}
