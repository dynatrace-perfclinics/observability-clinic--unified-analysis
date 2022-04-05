/**
 * Copyright (c) 2021 Radu Stefan
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 **/

package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/radu-stefan-dt/fleet-simulator/pkg/constants"
	"github.com/radu-stefan-dt/fleet-simulator/pkg/util"
)

func NewCli() {
	fmt.Println("Welcome to " + constants.Title + " " + constants.Version)
	fmt.Println(constants.ShortHelp)
	for {
		fmt.Print(constants.InputStr)
		reader := bufio.NewReader(os.Stdin)
		cmdString, err := reader.ReadString('\n')
		if err != nil {
			util.PrintError(err)
		}
		cmdFields := strings.Fields(strings.TrimSuffix(cmdString, "\n"))
		cmd := cmdFields[0]
		args := cmdFields[1:]

		switch cmd {
		case "help":
			help()
		case "start":
			startCommand(args)
		case "stop":
			stopCommand()
		case "exit":
			exit()
		}
	}
}

func help() {
	fmt.Fprintln(os.Stdout, constants.LongHelp)
}

func stopCommand() {
	fmt.Fprintln(os.Stdout, "You ran the STOP command")
}

func exit() {
	os.Exit(0)
}
