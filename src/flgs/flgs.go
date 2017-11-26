package flgs

import (
	"fmt"
	"log"
	"os/exec"
)

func ShowHelp() {
	fmt.Println(`

Usage: CLI Template [OPTIONS]
Options:
	-h, --help      Print the help log.
	
`)
}

func Status() {
	sCommand := exec.Command("git", "status")
	out, err := sCommand.Output()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s", out)
}
