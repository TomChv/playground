package main

import (
	"fmt"
	"io/fs"
	"log"
	"path"
	"path/filepath"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/cuecontext"
	cueload "cuelang.org/go/cue/load"
)

// Build a cue configuration tree from the files in fs.
func Build(src string, overlays map[string]fs.FS, args ...string) (cue.Value, error) {
	buildConfig := &cueload.Config{
		Dir:     src,
		Overlay: map[string]cueload.Source{},
	}

	// Map the source files into the overlay
	for mnt, f := range overlays {
		f := f
		mnt := mnt
		err := fs.WalkDir(f, ".", func(p string, entry fs.DirEntry, err error) error {
			if err != nil {
				return err
			}

			if !entry.Type().IsRegular() {
				return nil
			}

			if filepath.Ext(entry.Name()) != ".cue" {
				return nil
			}

			contents, err := fs.ReadFile(f, p)
			if err != nil {
				return fmt.Errorf("%s: %w", p, err)
			}

			overlayPath := path.Join(buildConfig.Dir, mnt, p)
			buildConfig.Overlay[overlayPath] = cueload.FromBytes(contents)
			return nil
		})
		if err != nil {
			return cue.Value{}, err
		}
	}
	instances := cueload.Instances(args, buildConfig)

	instance := instances[0]
	if err := instance.Err; err != nil {
		return cue.Value{}, err
	}

	cuectx := cuecontext.New()
	v := cuectx.BuildInstance(instance)
	if err := v.Err(); err != nil {
		return cue.Value{}, err
	}
	if err := v.Validate(); err != nil {
		return cue.Value{}, err
	}
	return v, nil
}

func main() {
	v, err := Build("/private/tmp/test/foo", nil, []string{"."}...)
	if err != nil {
		panic(err)
	}

	log.Println(v)
}
