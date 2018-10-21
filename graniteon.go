package GraniteON

import (
	"fmt"
	"io"
	"reflect"
)

type GraniteON struct {
	Object any
}

func (g GraniteON) WriteTo(w io.Writer) (n int64, err error) {
	return marshalAny(g.Object, w)
}

func (g *GraniteON) ReadFrom(r io.Reader) (n int64, err error) {
	g.Object, n, err = unmarshalAny(r)
	return
}

type NotMarshalable struct {
	Object any
}

func (e NotMarshalable) Error() string {
	var typeName string
	if typ := reflect.TypeOf(e.Object); typ == nil {
		typeName = "nil"
	} else {
		typeName = typ.Name()
	}

	return fmt.Sprintf("non-marshalable value: %s(%#v)", typeName, e.Object)
}

type NotUnmarshalable struct {
	Type byte
}

func (e NotUnmarshalable) Error() string {
	return fmt.Sprintf("non-unmarshalable type: byte(%d)", e.Type)
}
