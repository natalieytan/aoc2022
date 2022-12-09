package day7

const TotalDiskSpace = 70000000
const MinUnusedSpace = 30000000

func BuildRootDirectory(CommandDetails []CommandDetails) *Directory {
	rootDirectory := &Directory{
		name: "/",
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
				for _, directory := range lastTraversed[len(lastTraversed)-1].contents.subDirectories {
					if directory.name == commandDetail.directory {
						lastTraversed = append(lastTraversed, directory)
					}
				}

			}
		case CommandLs:
			currentDirectory := lastTraversed[len(lastTraversed)-1]
			currentDirectory.contents = commandDetail.directoryContents
		}
	}
	return rootDirectory
}

func Part1(d *Directory) int {
	return d.totalSizeIncludingSubdirectoriesUpTo100000()
}

func Part2(directory *Directory) int {
	totalSize := directory.totalSize()
	spaceToFree := totalSize + MinUnusedSpace - TotalDiskSpace

	return directory.directorySizeToDelete(spaceToFree, directory.totalSize())
}
