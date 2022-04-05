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
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/radu-stefan-dt/fleet-simulator/pkg/rest"
	"github.com/radu-stefan-dt/fleet-simulator/pkg/simulator"
	"github.com/radu-stefan-dt/fleet-simulator/pkg/util"
)

const (
	startCommandHelp = `
	Command starts up the fleet simulation. Format is "start [arguments] [flags]"
	
	Arguments:
		--environment, -e	- Dynatrace SaaS or Managed tenant and domain. You don't need https:// or the ending slash
		--token, -t		- Dynatrace API Token with Metrics (V2) permission or the name of an environment variable (if -ev is used)
		--fleets, -f		- Number of fleets to simulate (max. 10) (default: 2)
		--taxisPerFleet, -tpf	- Number of taxis per fleet to simulate. Ranges supported using the 'min-max' format for more variety (default: 5)
	Flags:
	    --env-vars, -ev		- Token is taken from an environment variable specified with -t
		--verbose, -v		- Prints out every single metric line sent
	
	Example:
		start -e abc123.live.dynatrace.com -t abcdefg1234567 -f 3 -tpf 2-5`
)

func startCommand(args []string) {
	var (
		environment string
		token       string
		taxisFlag   string
		taxis       int
		fleets      int
	)
	verbose := false
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "ERROR: No arguments provided. Type 'help' to see usage.")
		return
	}
	if args[0] == "help" {
		printStartCommandHelp()
		return
	}
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--environment", "-e":
			environment = parseFlagEnvironment(args[i+1])
		case "--token", "-t":
			env := strings.Contains(strings.Join(args, ""), "-ev") || strings.Contains(strings.Join(args, ""), "--env-var")
			token = parseFlagToken(args[i+1], env)
		case "--fleet", "-f":
			nfleets, err := strconv.Atoi(args[i+1])
			if err != nil {
				util.PrintError(err)
			}
			fleets = parseFlagNumFleets(nfleets)
		case "--taxisPerFleet", "-tpf":
			taxisFlag = args[i+1]
		case "--verbose", "-v":
			verbose = true
		}
	}
	if fleets == 0 {
		fleets = 2
	}
	if taxisFlag == "" {
		taxis = 5
	} else {
		taxis = parseFlagNumTaxis(taxisFlag)
	}
	client := rest.NewDTClient(environment, token)
	simulator.StartSimulation(client, int(fleets), taxis, verbose)
}

func printStartCommandHelp() {
	fmt.Fprintln(os.Stdout, startCommandHelp)
}
