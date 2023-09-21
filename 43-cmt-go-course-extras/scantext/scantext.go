package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

var (
	r_colon  = regexp.MustCompile(`:`)
	r_indent = regexp.MustCompile(`^\s+`)
)

func read_file() {
	file, err := os.Open("teams-go-3.vcs")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if r_indent.MatchString(line) {
			continue
		}

		a := r_colon.Split(line, 2)
		field := strings.ToLower(a[0])
		contents := a[1]
		fmt.Printf("%-20s  --  %s\n", field, contents)

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	read_file()
}
