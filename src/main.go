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
	"database/sql"
	"fmt"
	"net"
)

type Thing interface {
	setId(i int)
	getId() int
	SetName(n string)
	SetDesc(n string)
	GetName() string
	GetDesc() string
	SetParent(c Container)
}

type Container interface {
	GetContents() []Thing
	AddThing(t Thing)
	GetThingByName(n string) Thing
}

type Command interface {
	Help() string
	LongHelp() string
	Do()
}

// The top-level room of the world heirarchy
var rootRoom *Room = new(Room)

// The Admin
var rootUser *User = new(User)

func main() {

	/* 
	 * Do our init stuff here; connecting to DB, etc
	 */

	// Start the connection stack manager
	go ConnectionStackManager()

	// Start the DB handler
	db := new(dbHandler)
	dberr := db.Connect()

	if dberr != nil {
		fmt.Println(dberr)
		return
	}

	var dbsend chan string = make(chan string)
	var dbrecv chan *sql.Rows = make(chan *sql.Rows)

	db.dbsend = dbsend
	db.dbrecv = dbrecv

	go db.handleDB()

	defer db.Close()

	/*
	 * Init done, now start listening for connections
	 */

	ln, err := net.Listen("tcp", ":8067")
	if err != nil {
		fmt.Println("Failed listening: ", err)
		return
	}

	for {
		conn, err := ln.Accept()

		if err != nil {
			fmt.Println("Failed accepting: ", err)
			continue
		}

		// We've got a new connection; send it off to have fun

		c := new(UserConnection)
		c.C = conn

		c.awake = true
		append_to_cStack(c)

		go c.handleConnection()
	}
}
