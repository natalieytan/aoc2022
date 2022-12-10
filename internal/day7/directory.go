package day7

type Directory struct {
	name     string
	contents DirectoryContents
}

type DirectoryContents struct {
	files          []File
	subDirectories []*Directory
}

type File struct {
	name string
	size int
}

func (d *Directory) totalSize() int {
	total := 0
	for _, file := range d.contents.files {
		total += file.size
	}
	for _, directory := range d.contents.subDirectories {
		total += directory.totalSize()
	}
	return total
}

func (d *Directory) totalSizeUpTo100000() int {
	total := d.totalSize()
	if total <= 100000 {
		return total
	} else {
		return 0
	}
}

func (d *Directory) totalSizeIncludingSubdirectoriesUpTo100000() int {
	total := 0
	total += d.totalSizeUpTo100000()
	for _, subdirectory := range d.contents.subDirectories {
		total += subdirectory.totalSizeIncludingSubdirectoriesUpTo100000()
	}
	return total
}

func (d *Directory) directorySizeToDelete(spaceRequired int, closestSizeFound int) int {
	closestInSize := closestSizeFound

	if (d.totalSize() > spaceRequired) && (d.totalSize() <= closestInSize) {
		closestInSize = d.totalSize()
	}

	for _, subdirectory := range d.contents.subDirectories {
		if subdirectory.directorySizeToDelete(spaceRequired, closestInSize) < closestInSize {
			closestInSize = subdirectory.directorySizeToDelete(spaceRequired, closestInSize)
		}
	}

	return closestInSize
}
