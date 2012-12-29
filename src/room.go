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

type Room struct {
	id int
	name string
	desc string
	owner User
	contents []Thing
	parent Container
}

func (self *Room) setId(i int) {
	self.id = i
}

func (self *Room) getId() int {
	return self.id
}

func (self *Room) SetName(n string) {
	self.name = n
}

func (self *Room) SetDesc(n string) {
	self.desc = n
}

func (self *Room) GetName() string {
	return self.name
}

func (self *Room) GetDesc() string {
	return self.desc
}

func (self *Room) SetParent(c Container) {
	self.parent = c
}

func (self *Room) GetContents() []Thing {
	return self.contents
}

func (self *Room) AddThing(t Thing) {
	l := 0
	if self.contents != nil {
		l = len(self.contents)
	}

	n_carry := make([]Thing, l+1)
	copy(n_carry, self.contents)

	self.contents = n_carry
	self.contents[l] = t

	t.SetParent(self)
}

func (self *Room) GetThingByName(n string) Thing {
	l := 0
	if self.contents != nil {
		l = len(self.contents)
	}

	for i := 0; i < l; i++ {
		if self.contents[i].GetName() == n {
			return self.contents[i]
		}
	}

	return nil
}
