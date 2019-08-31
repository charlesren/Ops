package filenum

import (
	"fmt"
	"io/ioutil"
	"os"
)

// ThordHold is a leve
var ThordHold int

//Dir store directory name and total num of file and dir it contains
type Dir struct {
	Name string
	Num  int
}

// Dirs store directory name and file numbers
var Dirs []Dir

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
		info, err := ioutil.ReadDir(path)
		if err != nil {
			return err
		}
		if len(info) > ThordHold {
			iter := Dir{}
			iter.Name = path
			iter.Num = len(info)
			fmt.Printf("There are %v file or directory in %v !!!\n", len(info), path)
			Dirs = append(Dirs, iter)
			return nil
		}
	}
	return nil
}
