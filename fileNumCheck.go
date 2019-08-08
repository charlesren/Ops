package main

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/bitfield/script"
)

func main() {
	var ThordHold int
	ThordHold = 20000
	var Dir = []string{"/"}
	for a := &Dir; ; {
		if len(*a) == 0 {
			fmt.Println("All directory had been checked!!!")
			break
		} else {
			fmt.Println("hello")
			q := script.NewPipe()
			cmd := exec.Command("ls", "-al", (*a)[0])
			output, err := cmd.CombinedOutput()
			if err != nil {
				q.SetError(err)
			}
			q.WithReader(bytes.NewReader(output))
			var totalnum int
			totalnum, err = q.CountLines()
			if totalnum < ThordHold {
				fmt.Printf("There are %v file or directory in %v !!!\n", totalnum, (*a)[0])
			}
			data := string(output)
			fmt.Println(data)
			*a = (*a)[1:]
		}
	}
}
