package ops

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// ThordHold is a level
var ThordHold = 1000

type file struct {
	name string
	num  int
}

// Files store directory name and file numbers
var Files []file

// WalkDir define the directory you want to walk
var WalkDir = "/usr"

func main() {
	filepath.Walk(WalkDir, checkNum)
	fmt.Println(Files)
	SaveData()
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
			iter := file{}
			iter.name = path
			iter.num = len(dir)
			fmt.Printf("There are %v file or directory in %v !!!\n", len(dir), path)
			Files = append(Files, iter)
			return nil
		}
	}
	return nil
}
