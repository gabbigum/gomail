package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	usageMessage = `
Usage: gomail <command> [<args>]

Commands:  
  config    Configures gmail account
  send      Sends email

Run 'gomail <command> -help' for more information on a command.`
	)

func main() {
	// config command and its subcommands
	configCommand := flag.NewFlagSet("config", flag.ExitOnError)
	emailFlag := configCommand.String("u", "", "The gmail username\n'gomail config -u example@mail.com ...'")
	passFlag := configCommand.String("p", "", "The gmail password\n'gomail config -p pass123'")

	// sendCommand mail command and its subcommands
	sendCommand := flag.NewFlagSet("send", flag.ExitOnError)
	textFile := sendCommand.String("f", "", "File containing the text needed for the email.\n'gomail send -f <file_name> ...'")
	receiver := sendCommand.String("r", "", "The receivers email\n'gomail send -r example@mail.com ...'")

	if len(os.Args) == 1 {
		fmt.Println(usageMessage)
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
		if *emailFlag == "" && *passFlag == ""{
			fmt.Println("Example usage of config:\n'gomail config -u your_email@gmail.com -p pass123'")
		}
	}

	// send mail happens here
	if sendCommand.Parsed() {
		if *textFile == "" && *receiver == "" {
			fmt.Println("Example usage of send:\n'gomail send -f <file-name> -r <receiver-name>'")
			os.Exit(0)
		}
	}
}
