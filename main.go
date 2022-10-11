package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {
	log.SetFlags(0)

	var path string
	flag.StringVar(&path, "path", ".", "working directory, default is current directory")
	var output_file string
	flag.StringVar(&output_file, "file", "", "output file, default is stdout")
	flag.Parse()

	if output_file != "" {
		f, err := os.OpenFile(output_file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		log.SetOutput(f)
	}

	// list all files in a directory
	listDirContents(path, []string{})
}

// calculate md5 hash of a file
func md5File(filePath string) (string, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", md5.Sum(data)), nil
}

func processed(fileName string, processedDirectories []string) bool {
	for i := 0; i < len(processedDirectories); i++ {
		if processedDirectories[i] != fileName {
			continue
		}
		return true
	}
	return false
}

func listDirContents(path string, dirs []string) {
	files, _ := ioutil.ReadDir(path)

	for _, f := range files {
		var newPath string

		if path != "/" {
			newPath = fmt.Sprintf("%s/%s", path, f.Name())
		} else {
			newPath = fmt.Sprintf("%s%s", path, f.Name())
		}

		if f.IsDir() {
			if !processed(newPath, dirs) {
				dirs = append(dirs, newPath)
				listDirContents(newPath, dirs)
			}
		} else {
			hash, err := md5File(newPath)
			if err != nil {
				fmt.Println(err)
			}

			filename := filepath.Base(newPath)

			log.Printf("%s \t %s \t %s\n", filename, hash, newPath)
		}
	}
}
