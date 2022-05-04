package main

import (
	"fmt"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/cuecontext"
	"cuelang.org/go/cue/errors"
	cueload "cuelang.org/go/cue/load"
)

// Purpose: Find a way to retrieve a cue file and display his content
func main() {
	// Create a new context
	// Required to load a cue files
	cuectx := cuecontext.New()

	// Simple cue configuration
	// We will need it later
	config := &cueload.Config{}

	// Load a cue instance from the current directory
	instances := cueload.Instances([]string{"."}, config)

	// Retrieve first instance
	// There can be more if there are multiples packets in the directory (if I have understood well)
	// It returns a usable value.
	v := cuectx.BuildInstance(instances[0])
	if v.Err() != nil {
		fmt.Println(errors.Details(v.Err(), nil))
	}

	// Walk through all def
	Walk(v, func(v cue.Value) bool {
		fmt.Printf("Field %v is type of %v with value %v at position %v\n", v.Path(), v.IncompleteKind(), v, v.Pos())
		return true
	}, nil)
}

// Walk is an alternative to cue.Value.Walk which handles more field types
// You can customize this with your own options
func Walk(v cue.Value, before func(cue.Value) bool, after func(cue.Value)) {

	// call before and possibly stop recursion
	if before != nil && !before(v) {
		return
	}

	options := []cue.Option{
		cue.Definitions(true),
		cue.Hidden(true),
		cue.Optional(true),
	}

	// possibly recurse
	switch v.IncompleteKind() {
	case cue.StructKind:
		s, _ := v.Fields(options...)

		for s.Next() {
			Walk(s.Value(), before, after)
		}

	case cue.ListKind:
		l, _ := v.List()
		for l.Next() {
			Walk(l.Value(), before, after)
		}
	}

	if after != nil {
		after(v)
	}

}
