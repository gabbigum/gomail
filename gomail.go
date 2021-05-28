package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const (
	usageMessage = `
Usage: gomail <command> [<args>]

Commands:  
  config    Configures gmail account
  send      Sends email

Run 'gomail <command> -help' for more information on a command.`
	credentialsFile = "credentials.txt"
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
		if *emailFlag == "" && *passFlag == "" {
			fmt.Println("Example usage of config:\n'gomail config -u your_email@gmail.com -p pass123'")
		} else {
			// make them csv
			load := *emailFlag + "," + *passFlag
			f, err := os.OpenFile(credentialsFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)

			if err != nil {
				log.Fatalf("Error creating the %s file", credentialsFile)
				os.Exit(1)
			}
			saveCredentials(load, f)
			fmt.Println("Configuration was successful.")
		}
	}

	// send mail happens here
	if sendCommand.Parsed() {
		if *textFile == "" && *receiver == "" {
			fmt.Println("Example usage of send:\n'gomail send -f <file-name> -r <receiver-name>'")
		} else {
			email, pass := readCredentials(credentialsFile)
			// perform mail send
			fmt.Println(email, pass)
			// mail, password, file
			sendMail(email, pass)
		}
	}
}

// Saves credentials to specified Writer
func saveCredentials(load string, writer io.Writer) {
	_, err := writer.Write([]byte(load))
	if err != nil {
		log.Fatalf("Could not write the data.")
	}
}

func readCredentials(fileName string) (email, pass string) {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalf("Couldn't read from %s file", fileName)
	}
	tokens := strings.Split(string(file), ",")
	return tokens[0], tokens[1]
}

func sendMail(email, pass string) {

}