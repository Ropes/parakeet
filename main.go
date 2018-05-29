package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/ropes/parakeet/git"
)

func main() {
	// Read args
	flag.Parse()
	urlRaw := flag.Arg(0)
	if urlRaw == "" {
		fmt.Println("url arg not parsed")
		os.Exit(1)
	}

	scanner := bufio.NewScanner(os.Stdin)
	var builder strings.Builder
	//  Read Lines
	for scanner.Scan() {
		raw := scanner.Text()
		log := git.NewLogParser()
		err := log.Parse(raw)
		if err != nil {
			fmt.Fprint(os.Stderr, err)
			continue
		}
		log.Parse(raw)
		u, err := url.Parse(urlRaw)
		if err != nil {
			fmt.Printf("error parsing url: %v\n", err)
			os.Exit(1)
		}
		//  Process {line} with {project}
		builder.WriteString(log.ProjectMarkdown(*u) + "\n")
	}

	// Dump to stdout
	fmt.Printf(builder.String())
}
