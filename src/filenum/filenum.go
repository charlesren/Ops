package filenum

import (
	"fmt"
	"io/ioutil"
	"os"
)

// ThordHold is a leve
var ThordHold int

type file struct {
	Name string
	Num  int
}

// Files store directory name and file numbers
var Files []file

//ErrCode define errcode number of filenum func
const ErrCode = "29999"

//WarnDesc describle waring message
const WarnDesc = "File or dir number out of range"

// WalkDir define the directory you want to walk
//var WalkDir string

// CheckNum is a walkfunc  used to check which directory had files over the thordhold
func CheckNum(path string, f os.FileInfo, err error) error {
	if f == nil {
		return err
	}
	if f.IsDir() {
		dir, err := ioutil.ReadDir(path)
		if err != nil {
			return err
		}
		if len(dir) > ThordHold {
			iter := file{}
			iter.Name = path
			iter.Num = len(dir)
			fmt.Printf("There are %v file or directory in %v !!!\n", len(dir), path)
			Files = append(Files, iter)
			return nil
		}
	}
	return nil
}
