package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/smtp"
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
	host = "smtp.gmail.com"
	port = "587"
)

type Email struct {
	Subject string
	Content []byte
}

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
			username, pass := readCredentials(credentialsFile)
			readMail, _ := ioutil.ReadFile(*textFile)

			title := strings.Split(*textFile, ".")

			email := Email{
				Subject: title[0],
				Content: readMail,
			}

			sendMail(username, pass, *receiver, email)
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

// reads credentials from config file
func readCredentials(fileName string) (email, pass string) {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalf("Couldn't read from %s file", fileName)
	}
	tokens := strings.Split(string(file), ",")
	return tokens[0], tokens[1]
}

func sendMail(username, pass, receiver string, email Email) {
	auth := smtp.PlainAuth("", username, pass, host)
	to := []string{
		receiver,
	}

	err := smtp.SendMail(host + ":" + port, auth, username, to, email.Content)

	if err != nil {
		log.Fatalf("Could not send email %v, err %v", email, err)
		return
	}
	fmt.Println("Email sent successfully.")
}
