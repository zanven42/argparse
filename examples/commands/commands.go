package main

import (
	"fmt"
	"os"

	"github.com/zanven42/argparse"
)

// Run this as `go run commands.go [start|stop]`
func main() {
	// Create new parser object
	parser := argparse.NewParser("commands", "Simple example of argparse commands")
	// help := parser.Flag("h", "help", &argparse.Options{Help: "Print help information"})

	// NoExit means when help is called we will not os.Exit(0)
	parser.NoExit = true
	// Add top level command `start`
	startCmd := parser.NewCommand("start", "Will start a process")

	// Add top level commands `stop`
	stopCmd := parser.NewCommand("stop", "Will stop a process")

	// Parse command line arguments and in case of any error print error and help information
	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		return
	}

	// if help != nil && *help {
	// 	fmt.Println("Got to help")
	// 	fmt.Print(parser.Usage(nil))
	// 	return
	// }
	// Check if `start` command was given
	if startCmd.Happened() {
		// Starting a process
		fmt.Println("Started process")
	} else if stopCmd.Happened() { // Check if `stop` command was given
		// Stopping a process
		fmt.Println("Stopped process")
	} else {
		// In fact we will never hit this one
		// because commands and sub-commands are considered as required
		err := fmt.Errorf("bad arguments, please check usage")
		fmt.Print(parser.Usage(err))
	}
}
