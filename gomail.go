package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// config command and its subcommands
	configCommand := flag.NewFlagSet("config", flag.ExitOnError)
	emailFlag := configCommand.String("u", "", "Gmail username e.g. 'my_account@gmail.com'")
	passFlag := configCommand.String("p", "", "Gmail account password")

	// sendCommand mail command and its subcommands
	sendCommand := flag.NewFlagSet("send", flag.ExitOnError)
	textFile := sendCommand.String("f", "", "File containing the email text.")
	receiver := sendCommand.String("r", "", "The email of the receiver")

	if len(os.Args) == 1 {
		fmt.Println("usage: gomail <command> [<args>]")
		fmt.Println("The most commonly used gomail commands are: ")
		fmt.Println("config Config gmail account")
		fmt.Println("sendCommand Send an email")
		return
	}

	// parse the subcommands
	switch os.Args[1] {
	case "config":
		configCommand.Parse(os.Args[2:])
	case "send":
		sendCommand.Parse(os.Args[2:])
	default:
		fmt.Printf("%q is not a valid command.\n", os.Args[1])
		os.Exit(2)

	}

	// config happens here
	if configCommand.Parsed() {
		if *emailFlag == "" {
			fmt.Println("Please provide your gmail account name.")
		}

		if *passFlag == "" {
			fmt.Println("Please provide your gmail password.")
		}
	}

	// send mail happens here
	if sendCommand.Parsed() {
		if *textFile == "" {
			fmt.Println("Please provide the file which is to be sent.")
		}

		// do validation if the receiver email is correct
		if *receiver == "" {
			fmt.Println("Please provide the receiver email.")
		}
	}
}
