package main

import (
	"fmt"

	_ "cuelang.org/go/cue/ast"
	"cuelang.org/go/cue/format"
	cueload "cuelang.org/go/cue/load"
)

// Purpose: Find a way to retrieve a cue file and display his content
func main() {
	// Create a new context
	// Required to load a cue files
	//	cuectx := cuecontext.New()

	// Simple cue configuration
	// We will need it later
	config := &cueload.Config{}

	// Load a cue instance from the current directory
	instances := cueload.Instances([]string{"."}, config)

	instance := instances[0]
	file := instance.Files[0]

	decl := file.Decls[3]

	// EXPLORE https://github.com/cue-lang/cue/blob/v0.4.3/cue/ast/walk.go

	// Format beautiful display from ast
	display1, err := format.Node(decl)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(display1))

	// Retrieve first instance
	// There can be more if there are multiples packets in the directory (if I have understood well)
	// It returns a usable value.
	//	v := cuectx.BuildInstance(instance)
	//
	//	// Generate AST from cue instance
	//	ast := v.Syntax(
	//		// Uncomment to close struct and lists
	//		// Comment to get theoretical plan
	//		//cue.Final(),
	//		cue.ResolveReferences(true),
	//
	//		// Display docs
	//		cue.Docs(true),
	//		cue.Attributes(true),
	//		cue.Optional(true),
	//		cue.Definitions(false),
	//		// cue.Raw(),
	//		// cue.Schema(),
	//	)
	//
	//	// Format beautiful display from ast
	//	display, err := format.Node(ast)
	//	if err != nil {
	//		panic(err)
	//	}
	//
	//	fmt.Println(string(display))
}
