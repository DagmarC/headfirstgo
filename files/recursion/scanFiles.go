package recursion

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func OpenFile(path string) (*os.File, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func CloseFile(file *os.File) {
	err := file.Close()
	if err != nil {
		panic(err)
	}
}

// ScanFiles will recursively print out all the directories and files from the source directory (stated in main function).
func ScanFiles(path string) {

	//file, err := OpenFile(path)
	//if err != nil {
	//	// Panic example: Instead of returning the error value, pass it to panic (not preferred way.)
	//	panic(err)
	//}
	//defer CloseFile(file)

	fmt.Println("Current directory:", path)

	dirs, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}

	for _, file := range dirs {
		currentPath := filepath.Join(path, file.Name())
		if !file.IsDir() {
			fmt.Println(currentPath)
		} else {
			ScanFiles(currentPath)
		}
	}
}
