package main

import (
	"fmt"
	"os/exec"
)

func main() {

	var Dir = []string{"/"}
	for a := &Dir; ; {
		if len(*a) == 0 {
			fmt.Println("All directory had been checked!!!")
			break
		} else {
			fmt.Println("hello")
			cmd := exec.Command("ls", (*a)[0])
			*a = (*a)[1:]

		}
	}
}
