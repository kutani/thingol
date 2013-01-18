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
	"fmt"
	"io"
	"crypto/sha256"
)

type Muck struct {
	db *dbHandler
}

var theMuck Muck


func (self *Muck) auth(u string, p string) bool {
	passhash := shapass(p)

	// TODO Sanitize string u for sql safety

	q := "select passhash from users where name = '"+u+"';\n"

	self.db.dbsend <- q

	r := <-self.db.dbrecv

	var ret string

	if ! r.Next() {
		return false
	}

	err := r.Scan(&ret)

	if err != nil {
		fmt.Println(err)
		return false
	}


	fmt.Printf("Comparing %s vs %s\n",passhash,ret)
	if passhash == ret {
		return true
	}

	return false
}

func (self *Muck) newacct(u string, p string) bool {
	passhash := shapass(p)

	q := "select name from users where name = '"+u+"';\n"
	self.db.dbsend <- q
	r := <-self.db.dbrecv

	if r.Next() {
		return false
	}

	q = "insert into users (name,passhash) values ('"+u+"','"+passhash+"');\n"
	self.db.dbsend <- q
	r = <-self.db.dbrecv

	if ! r.Next() {
		return false
	}

	return true;
}

func shapass(p string) string {
	h := sha256.New()
	io.WriteString(h, p)
	return fmt.Sprintf("%x", h.Sum(nil))
}
