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
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/radu-stefan-dt/fleet-simulator/pkg/util"
)

func parseFlagEnvironment(env string) string {
	env = strings.TrimPrefix(env, "https://")
	env = strings.TrimSuffix(env, "/")
	return env
}

func parseFlagToken(val string, env bool) string {
	if env {
		return os.Getenv(val)
	}
	return val
}

func parseFlagNumFleets(nf int) int {
	switch {
	case nf < 1:
		return 1
	case nf > 10:
		return 10
	default:
		return nf
	}
}

func parseFlagNumTaxis(nt string) int {
	if strings.Contains(nt, "-") {
		splits := strings.Split(nt, "-")
		min, err := strconv.ParseInt(splits[0], 0, 0)
		if err != nil {
			util.PrintError(err)
		}
		max, err := strconv.ParseInt(splits[1], 0, 0)
		if err != nil {
			util.PrintError(err)
		}
		return rand.New(rand.NewSource(time.Now().UnixNano())).Intn(int(max-min)) + int(min)
	}
	num, err := strconv.ParseInt(nt, 0, 0)
	if err != nil {
		util.PrintError(err)
	}
	return int(num)
}
