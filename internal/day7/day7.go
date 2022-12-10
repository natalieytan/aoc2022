package day7

const TotalDiskSpace = 70000000
const MinUnusedSpace = 30000000

func BuildRootDirectory(commands []Command) *Directory {
	rootDirectory := &Directory{
		name: "/",
	}
	lastTraversed := []*Directory{
		rootDirectory,
	}

	for _, command := range commands {
		command.Execute(&lastTraversed)
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
