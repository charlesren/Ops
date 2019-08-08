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
			cmd := exec.Command("ls", "l", (*a)[0])
			output, err := cmd.CombinedOutput()
			if err != nil {
				q.SetError(err)
			}
			q.WithReader(bytes.NewReader(output))
			var totalnum int
			totalnum, err = q.CountLines()
			if totalnum > ThordHold {
				fmt.Println(totalnum)
			}

			*a = (*a)[1:]

		}
	}
}
