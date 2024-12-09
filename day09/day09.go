package main

import (
	_ "embed"
	"fmt"
	"unsafe"
)

//go:embed input
var input string

func main() {
	index := parse()
	part1(index)
}

type File struct {
	ID     int
	Length int
}

func parse() []*File {
	index := make([]*File, 0, len(input)*2)

	nextFile := true
	var pos, fileID int
	for _, char := range input {
		num := int(char - '0')
		if nextFile {
			f := &File{ID: fileID, Length: num}
			fileID++
			// if len(index) <= pos+num {
			// 	index = slices.Grow(index, cap(index)*2)
			// 	index = index[:cap(index)]
			// }
			for range num {
				index = append(index, f)
			}
		} else {
			for range num {
				index = append(index, nil)
			}
		}
		pos += num
		nextFile = !nextFile
	}
	fmt.Printf("Index of capacity %d takes around %d bytes in memory\n", cap(index), int(unsafe.Sizeof(index[0]))*cap(index))
	return index
}

func part1(index []*File) {
	// printIndex(index)
	defragmentate(index)
	checksum := calChecksum(index)
	fmt.Printf("(Part 1) Filesystem checksum: %d \n", checksum)
}

func defragmentate(index []*File) {
	last := len(index) - 1
	for i, file := range index {
		if file != nil {
			continue
		}
		for index[last] == nil {
			if last == i {
				return
			}
			last--
		}
		index[i] = index[last]
		index[last] = nil
		// printIndex(index)
	}
}

func calChecksum(index []*File) (checksum int) {
	for i, file := range index {
		if file == nil {
			break
		}
		checksum += i * file.ID
	}
	return checksum
}

func printIndex(index []*File) {
	for _, file := range index {
		if file == nil {
			fmt.Print(".")
		} else {
			fmt.Print(file.ID)
		}
	}
	fmt.Println()
}
