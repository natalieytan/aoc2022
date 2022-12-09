package day7

type Command int

const (
	CommandCd Command = iota
	CommandLs
)

type CommandDetails struct {
	command           Command
	directory         string
	directoryContents DirectoryContents
}
