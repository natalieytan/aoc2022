package day7

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type File struct {
	Name string
	Size int
}

type Directory struct {
	Name     string
	Contents DirectoryContents
}

type DirectoryContents struct {
	Files          []*File
	SubDirectories []*Directory
}

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

func PrepareData(input []byte) ([]CommandDetails, error) {
	data := string(input)
	rawCommands := strings.Split(data, "$")

	commandDetails := make([]CommandDetails, len(rawCommands))

	for i, rawCommand := range rawCommands[1:] {
		commandDetail, err := parseCommandDetails(rawCommand)
		if err != nil {
			return commandDetails, err
		}
		commandDetails[i] = commandDetail
	}

	return commandDetails, nil
}

func parseCommandDetails(rawCommand string) (CommandDetails, error) {
	trimmedCommand := strings.TrimSpace(rawCommand)
	if strings.HasPrefix(trimmedCommand, "cd") {
		return parseCdCommand(trimmedCommand)
	} else if strings.HasPrefix(trimmedCommand, "ls") {
		return parseLsCommand(trimmedCommand)
	} else {
		return CommandDetails{}, fmt.Errorf("invalid command - %s", trimmedCommand)
	}
}

func parseCdCommand(rawCommand string) (CommandDetails, error) {
	strs := strings.Split(rawCommand, " ")
	if len(strs) != 2 {
		return CommandDetails{}, fmt.Errorf("invalid cd command %s --", rawCommand)
	}
	return CommandDetails{
		command:   CommandCd,
		directory: strs[1],
	}, nil
}

func parseLsCommand(rawCommand string) (CommandDetails, error) {
	strs := strings.Split(rawCommand, "\n")
	files := []*File{}
	directories := []*Directory{}

	for _, rawLs := range strs[1:] {
		contentInfo := strings.Split(rawLs, " ")
		if len(contentInfo) > 2 {
			return CommandDetails{}, errors.New("invalid file content info")
		} else if contentInfo[0] == "dir" {
			directories = append(directories, &Directory{
				Name:     contentInfo[1],
				Contents: DirectoryContents{},
			})
		} else {
			fileSize, err := strconv.Atoi(contentInfo[0])
			if err != nil {
				return CommandDetails{}, errors.New("invalid file content info")
			}
			files = append(files, &File{
				Name: contentInfo[1],
				Size: fileSize,
			})
		}
	}

	return CommandDetails{
		command: CommandLs,
		directoryContents: DirectoryContents{
			Files:          files,
			SubDirectories: directories,
		},
	}, nil
}

func NewDirectory(name string) *Directory {
	return &Directory{
		Name: name,
	}
}

func (d *Directory) BuildContents(contents DirectoryContents) {
	d.Contents = contents
}

func (d *Directory) TotalSize() int {
	total := 0
	for _, file := range d.Contents.Files {
		total += file.Size
	}
	for _, directory := range d.Contents.SubDirectories {
		total += directory.TotalSize()
	}
	return total
}

func (d *Directory) SizeMax100_000() int {
	total := 0
	for _, file := range d.Contents.Files {
		total += file.Size
	}
	for _, directory := range d.Contents.SubDirectories {
		total += directory.TotalSize()
	}

	if total <= 100000 {
		return total
	} else {
		return 0
	}
}

func (d *Directory) TotalSizeMax100_000() int {
	total := 0
	total += d.SizeMax100_000()

	for _, subdirectory := range d.Contents.SubDirectories {
		total += subdirectory.TotalSizeMax100_000()
	}
	return total
}

func DirectoryTraverser(CommandDetails []CommandDetails) *Directory {
	rootDirectory := &Directory{
		Name: "/",
	}
	lastTraversed := []*Directory{}

	for _, commandDetail := range CommandDetails {
		switch commandDetail.command {
		case CommandCd:
			if commandDetail.directory == "/" {
				lastTraversed = []*Directory{rootDirectory}
			} else if commandDetail.directory == ".." {
				lastTraversed = lastTraversed[:len(lastTraversed)-1]
			} else {
				for _, directory := range lastTraversed[len(lastTraversed)-1].Contents.SubDirectories {
					if directory.Name == commandDetail.directory {
						lastTraversed = append(lastTraversed, directory)
					}
				}

			}
		case CommandLs:
			lastTraversed[len(lastTraversed)-1].BuildContents(commandDetail.directoryContents)
		}
	}
	return rootDirectory
}

func FindTotalSize(commandDetails []CommandDetails) int {
	rootDirectory := DirectoryTraverser(commandDetails)
	rootDirectory.PrintContents("")
	return rootDirectory.TotalSize()
}

func FindTotalSizeMax100_000(commandDetails []CommandDetails) int {
	rootDirectory := DirectoryTraverser(commandDetails)
	return rootDirectory.TotalSizeMax100_000()
}

func (directory *Directory) FindDirectorySizeToDelete() int {
	totalSize := directory.TotalSize()
	spaceToFree := totalSize + 30000000 - 70000000

	return directory.directorySizeToDelete(spaceToFree, directory.TotalSize())
}

func (d *Directory) directorySizeToDelete(spaceRequired int, closestSizeFound int) int {
	closestInSize := closestSizeFound

	if (d.TotalSize() > spaceRequired) && (d.TotalSize() <= closestInSize) {
		closestInSize = d.TotalSize()
	}

	for _, subdirectory := range d.Contents.SubDirectories {
		if subdirectory.directorySizeToDelete(spaceRequired, closestInSize) < closestInSize {
			closestInSize = subdirectory.directorySizeToDelete(spaceRequired, closestInSize)
		}
	}

	return closestInSize
}

func (d *Directory) PrintContents(padding string) {
	fmt.Printf(padding + "🐱 Files: \n")
	for _, file := range d.Contents.Files {
		fmt.Printf(padding+"%d - %s\n", file.Size, file.Name)
	}
	fmt.Printf(padding + "🐶 Directories: \n")
	for _, subdirectory := range d.Contents.SubDirectories {
		fmt.Printf(padding+"%s\n", subdirectory.Name)
		subdirectory.PrintContents(fmt.Sprintf("%s  ", padding))
	}
}
