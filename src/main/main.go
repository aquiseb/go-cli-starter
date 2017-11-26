package main

import (
	"flag"
	"flgs"
	"header"
	log "logger"
	"userinput"
)

// If user makes wrong usage of flags, we trigger ShowHelp()
func validateFlag(flag *flag.FlagSet) {
	flag.Usage = func() {
		flgs.ShowHelp()
	}
}

func main() {

	// Simple name prompt
	question := "What's your name? \n"
	response := "Your name is: "
	required := false
	userinput.QAR(question, response, required)

	// Define status
	var status bool
	flag.BoolVar(&status, "s", false, "")
	flag.BoolVar(&status, "status", false, "")

	// Check if the current flag is valid
	validateFlag(flag.CommandLine)
	// Show the fancy ASCII header
	header.Show()
	// Initialize the logger (to get date)
	log.InitLogger()

	// Parse the command-line flags
	flag.Parse()

	// If bool status is true, execute git status
	if status {
		flgs.Speak()
	}

}
