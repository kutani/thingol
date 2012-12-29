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
	rootUser.SetName("Admin")
	rootUser.setId(1)

	rootRoom.SetName("Room 0")
	rootRoom.AddThing(rootUser)

	tlist := rootRoom.GetContents()

	fmt.Println("Room Contains:")
	l := len(tlist)
	for i := 0; i < l; i++ {
		fmt.Println(tlist[i].GetName())
	}
}
