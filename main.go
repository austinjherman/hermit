package main

import (
	"fmt"
	"flag"
	"os"
	"strings"
)

func main() {
	
	toUpperCmd := flag.NewFlagSet("to_upper", flag.ExitOnError)

	if len(os.Args) < 3 {
		fmt.Println("usage: hermit <command> [<args>]")
		fmt.Printf("given: %s\n", os.Args)
		os.Exit(1)
	}

	switch os.Args[1] {

		case "to_upper":
			toUpperCmd.Parse(os.Args[2:])
			args := toUpperCmd.Args()
			argsString := strings.Join(args, " ")
			fmt.Println(strings.ToUpper(argsString))
			os.Exit(0)

		default:
			fmt.Printf("%q is not valid command.\n", os.Args[1])
			os.Exit(2)
	}

}