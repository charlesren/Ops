package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// ThordHold is a level
var ThordHold = 5

func main() {
	filepath.Walk("/", checkNum)

}
func checkNum(path string, f os.FileInfo, err error) error {
	if f == nil {
		return err
	}
	if f.IsDir() {
		dir, err := ioutil.ReadDir(path)
		if err != nil {
			return err
		}
		if len(dir) > ThordHold {
			fmt.Printf("There are %v file or directory in %v !!!\n", len(dir), path)
			return nil
		}
	}
	//fmt.Println("hello")
	return nil
}
