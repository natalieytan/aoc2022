package day7

type Command interface {
	Execute(directoryHistory *[]*Directory)
}

type LsCommand struct {
	directoryContents DirectoryContents
}

func (c LsCommand) Execute(directoryHistory *[]*Directory) {
	directiories := *directoryHistory
	currentDirectory := directiories[len(directiories)-1]
	currentDirectory.contents = c.directoryContents
}

type CdCommand struct {
	directory string
}

func (c CdCommand) Execute(directoryHistory *[]*Directory) {
	directories := *directoryHistory
	if c.directory == "/" {
		*directoryHistory = directories[0:1]
	} else if c.directory == ".." {
		*directoryHistory = directories[:len(directories)-1]
	} else {
		for _, directory := range directories[len(directories)-1].contents.subDirectories {
			if directory.name == c.directory {
				*directoryHistory = append(*directoryHistory, directory)
			}
		}
	}
}
