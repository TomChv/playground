package definitions

import (
	"fmt"
	"sync"

	"cuelang.org/go/cue"
)

type Definitions struct {
	vs []cue.Value
	m  sync.Mutex
}

func New() *Definitions {
	return &Definitions{}
}

func (d *Definitions) Append(v cue.Value) {
	d.m.Lock()
	defer d.m.Unlock()
	d.vs = append(d.vs, v)
}

func (d *Definitions) Get(path string) (cue.Value, error) {
	for _, v := range d.vs {
		if v.Path().String() == path {
			return v, nil
		}
	}

	return cue.Value{}, fmt.Errorf("value not found")
}

func (d *Definitions) String() string {
	str := ""
	for _, v := range d.vs {
		str += fmt.Sprintf("path: %s\nat %s\nvalue: %s\n\n", v.Path(), v.Pos(), v)
	}

	return str
}
