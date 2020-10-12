package main

import (
	"fmt"
	"flag"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	
	toCapsCmd := flag.NewFlagSet("caps", flag.ExitOnError)
	randCmd := flag.NewFlagSet("rand", flag.ExitOnError)
	randMax := randCmd.Int("mx", 100, "the maximum")
	randMin := randCmd.Int("mn", 0, "the minimun")

	if len(os.Args) < 2 {
		fmt.Println("usage: hermit <command> [<args>]")
		fmt.Printf("given: %s\n", os.Args)
		os.Exit(1)
	}

	switch os.Args[1] {

		case "caps":
			toCapsCmd.Parse(os.Args[2:])
			args := toCapsCmd.Args()
			argsString := strings.Join(args, " ")
			fmt.Println(strings.ToUpper(argsString))
			os.Exit(0)

		case "rand":
			randCmd.Parse(os.Args[2:])
			rand.Seed(time.Now().UnixNano())
			fmt.Println(rand.Intn(*randMax + 1 - *randMin) + *randMin)
			os.Exit(0)

		default:
			fmt.Printf("%q is not valid command.\n", os.Args[1])
			os.Exit(2)
	}

}