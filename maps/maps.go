package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	args := os.Args[1:]
	if len(args) == 0 {
		log.Fatal("Filename not provided")
	}

	log.Print(args)

	file, err := os.Open(args[0])

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	m := make(map[string]int)

	for scanner.Scan() {
		s := scanner.Text()
		v, existed := m[s]
		if existed {
			m[s] = v + 1
		} else {
			m[s] = 1
		}
	}

	for k, v := range m {
		fmt.Println(k, ":", v)
	}
}
