package day7

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func PrepareData(input []byte) ([]Command, error) {
	data := string(input)
	rawCommands := strings.Split(data, "$")
	commandDetails := make([]Command, len(rawCommands)-1)

	for i, rawCommand := range rawCommands[1:] {
		commandDetail, err := parseCommandDetails(rawCommand)
		if err != nil {
			return commandDetails, err
		}
		commandDetails[i] = commandDetail
	}

	return commandDetails, nil
}

func parseCommandDetails(rawCommand string) (Command, error) {
	trimmedCommand := strings.TrimSpace(rawCommand)
	if strings.HasPrefix(trimmedCommand, "cd") {
		return parseCdCommand(trimmedCommand)
	} else if strings.HasPrefix(trimmedCommand, "ls") {
		return parseLsCommand(trimmedCommand)
	} else {
		return nil, fmt.Errorf("invalid command - %s", trimmedCommand)
	}
}

func parseCdCommand(rawCommand string) (CdCommand, error) {
	strs := strings.Split(rawCommand, " ")
	if len(strs) != 2 {
		return CdCommand{}, fmt.Errorf("invalid cd command %s --", rawCommand)
	}
	return CdCommand{
		directory: strs[1],
	}, nil
}

func parseLsCommand(rawCommand string) (LsCommand, error) {
	strs := strings.Split(rawCommand, "\n")
	files := []File{}
	directories := []*Directory{}

	for _, rawLs := range strs[1:] {
		contentInfo := strings.Split(rawLs, " ")
		if len(contentInfo) > 2 {
			return LsCommand{}, errors.New("invalid file content info")
		} else if contentInfo[0] == "dir" {
			directories = append(directories, &Directory{
				name:     contentInfo[1],
				contents: DirectoryContents{},
			})
		} else {
			fileSize, err := strconv.Atoi(contentInfo[0])
			if err != nil {
				return LsCommand{}, errors.New("invalid file content info")
			}
			files = append(files, File{
				name: contentInfo[1],
				size: fileSize,
			})
		}
	}

	return LsCommand{
		directoryContents: DirectoryContents{
			files:          files,
			subDirectories: directories,
		},
	}, nil
}
