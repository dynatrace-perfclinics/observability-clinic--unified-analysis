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

package util

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

var locations = [10]string{
	"Bristol",
	"London",
	"Manchester",
	"Birmingham",
	"Liverpool",
	"Leeds",
	"Newcastle",
	"Nottingham",
	"Basingstoke",
	"Reading",
}

func PrintError(err error) {
	fmt.Fprintln(os.Stderr, err)
}

func Locations() [10]string {
	return locations
}

func RandomLetter() rune {
	time.Sleep(time.Nanosecond)
	return 'A' + rune(rand.New(rand.NewSource(time.Now().UnixNano())).Intn(26))
}

func GenerateRegNumber() string {
	var reg string
	reg += string(RandomLetter())
	reg += string(RandomLetter())
	reg += fmt.Sprintf("%d", rand.Intn(99))
	reg += " "
	reg += string(RandomLetter())
	reg += string(RandomLetter())
	reg += string(RandomLetter())
	return reg
}

func RandomCoord() float64 {
	rand.Seed(time.Now().UnixNano())
	time.Sleep(time.Nanosecond)
	return 20 + rand.Float64()*80
}
