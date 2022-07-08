package main

import (
	"test.com/hello"
)

#Foo: string

#Def: {
	// test
	foo: string
	bar: number
	baz: [...string]
	map: [string]: string
	_foo: string
	bar_foo: {
		_
	}
}

def: {
	test: #Def
}

"foo": "world"

ref: hello.#SayHi

// This field is hidden
_hidden: true

optional?: string

attribute: number @test()