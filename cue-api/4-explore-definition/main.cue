package main

import (
	"test.com/hello"
)

#test: string
test: #test & { "bar" }

"foo": "world"

ref: hello.#SayHi
ref2: hello.#SayBonjour

// This field is hidden
_hidden: true

optional?: string

attribute: number @test()

struct: {
	_hidden: false
}