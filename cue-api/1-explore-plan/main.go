package main

import (
	"log"

	"cuelang.org/go/cue/cuecontext"
	cueload "cuelang.org/go/cue/load"
)

// Purpose: Read a cue file and explore all keys in it with their value and their type.
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

	// Retrieve the kind of our current value (root value)
	log.Println("Main value kind:", v.Kind())

	// Create an iterator through all values in the current one
	i, err := v.Fields()
	if err != nil {
		panic(err)
	}

	// Loop through it and display kind
	log.Println("Loop through value")
	for i.Next() != false {
		log.Println("Path: '", i.Value().Path(), "Kind: '", i.Value().Kind(), "'Undefined kind: '", i.Value().IncompleteKind(), "'")
	}
}
