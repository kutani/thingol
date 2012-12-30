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
	"bufio"
	"fmt"
	"net"
)

type UserConnection struct {
	C      net.Conn
	recbuf *bufio.Reader
	senbuf *bufio.Writer
}

func (self *UserConnection) handleConnection() {
	defer self.C.Close()

	self.recbuf = bufio.NewReader(self.C)
	self.senbuf = bufio.NewWriter(self.C)

	var input chan string = make(chan string)
	var output chan string = make(chan string)

	go self.Recieve(input)
	go self.Send(output)

	for {
		m := <-input
		go self.handleCommand(m, output)
	}

	return
}

func (self *UserConnection) handleCommand(m string, out chan<- string) {
	fmt.Println(m)
	return
}

func (self *UserConnection) Send(in <-chan string) {
	for {
		m := <-in

		ret, err := self.senbuf.WriteString(m)
		err = self.senbuf.WriteByte('\n')
		self.senbuf.Flush()

		if ret != len(m) && err != nil {
			// TODO Handle the error
			continue
		}
	}
}

func (self *UserConnection) Recieve(out chan<- string) {
	for {
		ret, err := self.recbuf.ReadString('\n')

		if err != nil {
			// TODO Handle the error
			continue
		}

		ret = ret[:len(ret)-1]

		out <- ret
	}
}
