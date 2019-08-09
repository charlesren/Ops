package main 

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/bitfield/script"
)

	var ThordHold int
	ThordHold = 20000
	var Dir = []string{"/"}
func temp() {
	func checkNum (path string, info os.FileInfo, err error ) error{
         if info.IsDir {
		dir, err := ioutil.ReadDir(pth)
		if err != nil {
		 return nil, err
		}
		if len(dir) < ThordHold {
			fmt.Printf("There are %v file or directory in %v !!!\n", len(dir), path)
		}
	}
