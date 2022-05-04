package main

import (
	"test.com/hello"
)


"foo": "world"

ref: hello.#SayHi

// This field is hidden
_hidden: true

optional?: string

attribute: number @test()

struct: {
	_hidden: false
}