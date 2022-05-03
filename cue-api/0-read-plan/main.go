package main

import (
	"fmt"

	"cuelang.org/go/cue/cuecontext"
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

	// Marshal cue value to display it
	json, err := v.MarshalJSON()
	if err != nil {
		panic(err)
	}

	fmt.Println(string(json))
}
