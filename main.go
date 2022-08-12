package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func path_trawl(dir string) []string {
	var blacklist_file_types = []string{
		".json",
		".yaml",
		".gitattributes",
		".gitmodules",
		".gitignore",
		".mod",
		".sum",
	}

	all_paths := make([]string, 1000) // slice of max length 1000

	err := filepath.Walk(dir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			} else {

				for _, elem := range strings.Split(path, "\n") {
					//TODO: skip dot files
					if contains(blacklist_file_types, filepath.Ext(elem)) {
						continue
					}

					all_paths = append(all_paths, elem)
				}
				return nil
			}
		})
	if err != nil {
		log.Println(err)
	}

	return all_paths
}

func file_length(filepath string) int {
	var length int = 0

	if len(filepath) == 0 {
		return 0
	}

	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//TODO: handle commented out lines
		if scanner.Text() != "" { // Ignore empty lines
			length++
		}
	}

	return length
}

func main() {
	PLACEHOLDER := "/users/ashisbitt/workspace/gh-stats"
	dir := PLACEHOLDER
	var count int = 0
	var file_count int = 0

	all_files := path_trawl(dir)

	for _, file := range all_files {
		count += file_length(file)
		if count != 0 {
			file_count++
		}
	}

	fmt.Println(fmt.Sprintf("Found %d lines of code across %d files", count, file_count))

}
