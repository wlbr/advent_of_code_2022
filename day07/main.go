package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readInput(fname string) (buffer []string) {
	f, err := os.Open(fname)
	if err != nil {
		log.Fatalf("Error opening dataset '%s':  %s", fname, err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		buffer = append(buffer, line)
	}
	return buffer
}

type file struct {
	parent *folder
	name   string
	size   int
}

func (f *file) String() string {
	return fmt.Sprintf("- %s (file, size=%d)\n", f.name, f.size)
}

func (f *file) RecursiveString(depth int) string {
	return fmt.Sprintf("%s- %s (file, size=%d)\n", strings.Repeat("  ", depth+1), f.name, f.size)
}

type folder struct {
	parent     *folder
	subfolders map[string]*folder
	files      map[string]*file
	name       string
	size       int
}

func (f *folder) String() string {
	return fmt.Sprintf("- %s (dir)\n", f.name)
}

func (f *folder) RecursiveString(depth int) string {
	indent := strings.Repeat("  ", depth)
	s := fmt.Sprintf("%s- %s (dir)\n", indent, f.name)
	for _, sf := range f.subfolders {
		s += fmt.Sprintf("%s%s", indent, sf.RecursiveString(depth+1))
	}
	for _, f := range f.files {
		s += f.RecursiveString(depth + 1)
	}
	return s
}

type fs struct {
	root *folder
	cwd  *folder
}

func (f *fs) String() string {
	return f.root.RecursiveString(0)
}

func NewFs() *fs {
	root := &folder{name: "/", size: 0, files: make(map[string]*file), subfolders: make(map[string]*folder)}
	this := &fs{root: root, cwd: root}

	return this
}

func (fs *fs) cmdCd(buffer []string, i int) int {
	command := strings.Split(buffer[i][2:], " ")
	parameters := command[1:]
	switch parameters[0] {
	case "..":
		fs.cwd = fs.cwd.parent
	case "/":
		fs.cwd = fs.root
	default:
		if newcwd, ok := fs.cwd.subfolders[parameters[0]]; !ok {
			log.Fatalf("Folder '%s' not found", parameters[0])
		} else {
			fs.cwd = newcwd
		}
	}

	return i
}

func (fs *fs) fixSize(current *folder) {
	size := 0
	for _, f := range current.files {
		size += f.size
	}
	for _, sf := range current.subfolders {
		size += sf.size
	}
	current.size = size
	if current != fs.root {
		fs.fixSize(current.parent)
	}
}

func (fs *fs) cmdLs(buffer []string, i int) int {
	for j := i + 1; j <= len(buffer)-1 && !strings.HasPrefix(buffer[j], "$ "); j++ {
		output := strings.Split(buffer[j], " ")
		switch output[0] {
		case "dir":
			fs.cwd.subfolders[output[1]] = &folder{parent: fs.cwd, name: output[1], size: 0, files: make(map[string]*file), subfolders: make(map[string]*folder)}
		default:
			size, err := strconv.Atoi(output[0])
			if err != nil {
				log.Fatalf("Error parsing size '%s':  %s", output[0], err)
			}
			fs.cwd.files[output[1]] = &file{parent: fs.cwd, name: output[1], size: size}
			fs.fixSize(fs.cwd)
		}
		i = j
	}
	return i
}

func (fs *fs) parseCommand(buffer []string, i int) int {
	command := strings.Split(buffer[i][2:], " ")
	switch command[0] {
	case "cd":
		i = fs.cmdCd(buffer, i)
	case "ls":
		i = fs.cmdLs(buffer, i)
	}
	return i
}

func createFileSystem(buffer []string) *fs {
	fs := NewFs()
	for i := 0; i < len(buffer); i++ {
		line := buffer[i]
		if strings.HasPrefix(line, "$ ") {
			i = fs.parseCommand(buffer, i)
		}
	}
	return fs
}

func (fs *fs) findSumOfSmallFolders(current *folder, limit int) int {
	sum := 0
	if current.size <= limit {
		sum = current.size
	}
	for _, sf := range current.subfolders {
		sum += fs.findSumOfSmallFolders(sf, limit)
	}

	return sum
}

func (fs *fs) findSmallestFolder(current *folder, limit int, smallestSoFar *folder) (smallestFolder *folder) {
	if current.size <= smallestSoFar.size && current.size >= limit {
		smallestFolder = current
	} else {
		smallestFolder = smallestSoFar
	}
	for _, sf := range current.subfolders {
		smallestFolder = fs.findSmallestFolder(sf, limit, smallestFolder)
	}
	return smallestFolder
}

func task1(input string) int {
	buffer := readInput(input)
	fs := createFileSystem(buffer)

	return fs.findSumOfSmallFolders(fs.root, 100000)
}

func task2(input string) int {
	buffer := readInput(input)
	fs := createFileSystem(buffer)

	unused := 70000000 - fs.root.size
	limit := 30000000 - unused
	return fs.findSmallestFolder(fs.root, limit, fs.root).size
}

func main() {
	//input := "input.txt"
	input := "example.txt"

	fmt.Println("Task 1 - sum of folder sizes below 100000             \t =  ", task1(input))
	fmt.Println("Task 2 - size of smallest folder freeing enough space \t =  ", task2(input))

}
