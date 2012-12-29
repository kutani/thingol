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

type User struct {
	id int
	name string
	desc string
	inventory []Thing
	parent Container
}

func (self *User) setId(i int) {
	self.id = i
}

func (self *User) getId() int {
	return self.id
}

func (self *User) SetName(n string) {
	self.name = n
}

func (self *User) SetDesc(n string) {
	self.desc = n
}

func (self *User) GetName() string {
	return self.name
}

func (self *User) GetDesc() string {
	return self.desc
}

func (self *User) SetParent(c Container) {
	self.parent = c
}

func (self *User) GetContents() []Thing {
	return self.inventory
}

func (self *User) AddThing(t Thing) {
	l := 0
	if self.inventory != nil {
		l = len(self.inventory)
	}
	n_iarry := make([]Thing, l+1)
	copy(n_iarry, self.inventory)

	self.inventory = n_iarry
	self.inventory[l] = t
	t.SetParent(self)
}

func (self *User) GetThingByName(n string) Thing {
	l := 0
	if self.inventory != nil {
		l = len(self.inventory)
	}

	for i := 0; i < l; i++ {
		if self.inventory[i].GetName() == n {
			return self.inventory[i]
		}
	}

	return nil
}
