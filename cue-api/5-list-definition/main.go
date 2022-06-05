package main

import (
	"fmt"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/build"
	"cuelang.org/go/cue/cuecontext"
	"cuelang.org/go/cue/errors"
	cueload "cuelang.org/go/cue/load"
	"cuelang.org/go/pkg/strings"
	"github.com/tomchv/cueapi/definitions"
)

type Value struct {
	cue.Value
}

func (v *Value) IsDefinition() bool {
	p := v.Path().String()

	return strings.HasPrefix(p, "#") && !strings.Contains(p, ".")
}

var defs = definitions.New()

func explorePackage(instance *build.Instance) {
	for _, i := range instance.Imports {
		explorePackage(i)
	}

	cuectx := cuecontext.New()

	v := cuectx.BuildInstance(instance)
	if v.Err() != nil {
		fmt.Println(errors.Details(v.Err(), nil))
		return
	}

	Walk(v, func(v cue.Value) bool {
		cv := &Value{v}
		if !cv.IsDefinition() {
			return true
		}

		defs.Append(v)

		return true
	}, nil)
}

// Purpose: Find a way to retrieve a cue file and display his content
func main() {
	// Create a new context
	// Required to load a cue files

	// Simple cue configuration
	// We will need it later
	config := &cueload.Config{}

	// Load a cue instance from the current directory
	instances := cueload.Instances([]string{"."}, config)

	explorePackage(instances[0])
	fmt.Println(defs)
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
		cue.ResolveReferences(true),
		cue.All(),
		cue.Docs(true),
		cue.Schema(),
		cue.Definitions(true),
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
