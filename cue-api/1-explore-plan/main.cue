package main

import (
	"test.com/hello"
)


"hello": number
"hello": 1

ref: hello.#SayHi

nested: dep1: "baz"

undefinedString: string

undefinedMultiChoice: string | { "foo": "bar" }