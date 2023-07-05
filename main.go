package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("> ")
	text, _ := reader.ReadString('\n')

	stmt, err := NewParser(strings.NewReader(text)).Parse()
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(stmt)
}
