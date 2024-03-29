/*
Copyright (c) 2013 K Kutani

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package main

import (
	"strconv"
)

type GoCmd struct {
	cmd      string
	help     string
	longhelp string
	function func()
}

var FuncMap map[string]Command = make(map[string]Command)

func (self *GoCmd) Help() string {
	return self.help
}

func (self *GoCmd) LongHelp() string {
	return self.longhelp
}

func (self *GoCmd) Do() {
	self.function()
}

func (self *GoCmd) SetFunc(f func()) {
	self.function = f
}


func WHO(out chan<- string) {
	l := len(cStack)

	out <- "There are "+strconv.Itoa(l)+" users connected:"
	for i := 0; i < l; i++ {
		if cStack[i].user == "" {
			continue
		}
		out <- cStack[i].user+"\tConnTime: "+strconv.Itoa(cStack[i].conntime)
	}
}
