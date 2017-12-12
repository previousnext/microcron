package main

import (
	"bufio"
	"fmt"
	"os/exec"
	"time"

	"github.com/ausrasul/hashgen"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	cliFrequency = kingpin.Arg("frequency", "How frequently to run the command").Required().Duration()
	cliCommand   = kingpin.Arg("command", "The command to run as part of this task").Required().String()
)

func main() {
	kingpin.Parse()

	limiter := time.Tick(*cliFrequency)

	for  {
		<- limiter
		go taskWithParams(*cliCommand)
	}
}

// Helper function for running our task.
func taskWithParams(command string) {
	hash := hashgen.Get(10)

	fmt.Println(hash, "|", "Starting new execution")

	cmd := exec.Command("/bin/bash", "-c", command)

	cmdReader, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(hash, "|", "Error creating StdoutPipe for Cmd", err)
		return
	}

	scanner := bufio.NewScanner(cmdReader)
	go func() {
		for scanner.Scan() {
			fmt.Println(hash, "|", scanner.Text())
		}
	}()

	err = cmd.Start()
	if err != nil {
		fmt.Println(hash, "|", "Error starting command:", err)
		return
	}

	err = cmd.Wait()
	if err != nil {
		fmt.Println(hash, "|", "Error waiting for command:", err)
	}
}