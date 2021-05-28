package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	configCommand := flag.NewFlagSet("config", flag.ExitOnError)
	emailFlag := configCommand.String("uname", "", "Gmail username e.g. 'my_account@gmail.com'")
	passFlag := configCommand.String("pass", "", "Gmail account password")

	if len(os.Args) == 1 {
		fmt.Println("usage: gomail <command> [<args>]")
		fmt.Println("The most commonly used gomail commands are: ")
		fmt.Println("config Config gmail account")
		fmt.Println("send Send an email")
		return
	}

	switch os.Args[1] {
	case "config":
		configCommand.Parse(os.Args[2:])
	default:
		fmt.Printf("%q is not a valid command.\n", os.Args[1])
		os.Exit(2)
	}

	if configCommand.Parsed() {
		if *emailFlag == "" {
			fmt.Println("Please provide your gmail account name.")
		}

		if *passFlag == "" {
			fmt.Println("Please provide your gmail password.")
		}
	}

	fmt.Println(*emailFlag)
	fmt.Println(*passFlag)
}
