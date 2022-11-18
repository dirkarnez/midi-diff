package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
)

func main() {
	leftFiles, err := GetLeftDirRecursive()
	if err != nil {
		panic(err)
	}
	rightFiles, err := GetRightDirRecursive()
	if err != nil {
		panic(err)
	}

	fmt.Println("left")
	for _, file := range leftFiles {
		fmt.Println(file.Name(), file.IsDir())
	}

	fmt.Println("right")
	for _, file := range rightFiles {
		fmt.Println(file.Name(), file.IsDir())
	}
}

func GetLeftDirRecursive() ([]fs.FileInfo, error) {
	return ioutil.ReadDir(os.Args[1])
}

func GetRightDirRecursive() ([]fs.FileInfo, error) {
	return ioutil.ReadDir(os.Args[2])
}
