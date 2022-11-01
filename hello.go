package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var pl = fmt.Println

func main() {
	fmt.Println("What's your name?")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	if err == nil {
		fmt.Println("MFer called: ", name)

	} else {
		log.Fatal(err)
	}

	var vName = "Are you stupid?"
	bName := "Hans"
	bName = "Klause"
}
