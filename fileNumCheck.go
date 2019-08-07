package main

import (
	"fmt"

	"github.com/bitfield/script"
)

func main() {

	var Dir = []string{"/"}
	for a := &Dir; ; {
		if len(*a) == 0 {
			fmt.Println("All directory had been checked!!!")
			break
		} else {
			fmt.Println("hello")
			p := script.Exec("ls (*a)[0]")
			output, err := p.String()
			if err != nil {
				fmt.Println(output)
			}
			fmt.Println(output)
			*a = (*a)[1:]

		}
	}
}
