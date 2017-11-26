package userinput

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// QAR means Question Answer Response. Used to prompt
// a question to the user, and give a response to it
func QAR(q string, r string, required bool) string {
	reader := bufio.NewReader(os.Stdin)

	// Print the question
	fmt.Printf(q)
	answer, _ := reader.ReadString('\n')

	if required && answer == "\n" {
		log.Fatal("You must enter a value.")
	}

	// Print the response with the answer
	if r != "" {
		fmt.Println(r, answer)
	}

	return answer
}
