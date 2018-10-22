package GraniteON

import (
	"encoding"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"sort"
)

func marshalAny(o any, w io.Writer) (n int64, err error) {
	switch oo := o.(type) {
	case nil:
		buf := [1]byte{scalarNil}
		n, errWr := w.Write(buf[:])
		return int64(n), errWr
	case bool:
		var buf [1]byte
		if oo {
			buf[0] = scalarTrue
		} else {
			buf[0] = scalarFalse
		}

		n, errWr := w.Write(buf[:])
		return int64(n), errWr
	case uint:
		if cpuHas32bits {
			return packUIntBE(uint64(oo), typeUInt|typeSize32, w)
		} else {
			return packUIntBE(uint64(oo), typeUInt|typeSize64, w)
		}
	case uint8:
		return packUIntBE(uint64(oo), typeUInt|typeSize8, w)
	case uint16:
		return packUIntBE(uint64(oo), typeUInt|typeSize16, w)
	case uint32:
		return packUIntBE(uint64(oo), typeUInt|typeSize32, w)
	case uint64:
		return packUIntBE(oo, typeUInt|typeSize64, w)
	case int:
		if cpuHas32bits {
			return packUIntBE(uint64(uint(oo)), typeInt|typeSize32, w)
		} else {
			return packUIntBE(uint64(uint(oo)), typeInt|typeSize64, w)
		}
	case int8:
		return packUIntBE(uint64(uint8(oo)), typeInt|typeSize8, w)
	case int16:
		return packUIntBE(uint64(uint16(oo)), typeInt|typeSize16, w)
	case int32:
		return packUIntBE(uint64(uint32(oo)), typeInt|typeSize32, w)
	case int64:
		return packUIntBE(uint64(oo), typeInt|typeSize64, w)
	case float32:
		return packFloatBE(uint64(math.Float32bits(oo))<<32, typeFloat|typeSize32, w)
	case float64:
		return packFloatBE(math.Float64bits(oo), typeFloat|typeSize64, w)
	case string:
		return marshalAny([]byte(oo), w)
	case []byte:
		n, err = packUIntBE(uint64(len(oo)), typeString, w)
		if err != nil {
			return
		}

		var m int
		m, err = w.Write(oo)
		n += int64(m)

		return
	case []any:
		n, err = packUIntBE(uint64(len(oo)), typeArray, w)
		if err != nil {
			return
		}

		for _, v := range oo {
			m, errWr := marshalAny(v, w)
			n += m

			if errWr != nil {
				return n, errWr
			}
		}

		return
	case map[string]any:
		sortedKeys := make([]string, len(oo))

		{
			i := 0

			for k := range oo {
				sortedKeys[i] = k
				i++
			}
		}

		sort.Strings(sortedKeys)

		n, err = packUIntBE(uint64(len(oo)), typeDict, w)
		if err != nil {
			return
		}

		for _, k := range sortedKeys {
			{
				m, errWr := marshalAny(k, w)
				n += m

				if errWr != nil {
					return n, errWr
				}
			}

			m, errWr := marshalAny(oo[k], w)
			n += m

			if errWr != nil {
				return n, errWr
			}
		}

		return
	case json.Marshaler:
		jsn, errMJ := oo.MarshalJSON()
		if errMJ != nil {
			return 0, errMJ
		}

		var unJson any
		if errUJ := json.Unmarshal(jsn, &unJson); errUJ != nil {
			return 0, errUJ
		}

		return marshalAny(unJson, w)
	case encoding.TextMarshaler:
		txt, errMT := oo.MarshalText()
		if errMT != nil {
			return 0, errMT
		}

		return marshalAny(txt, w)
	case fmt.Stringer:
		return marshalAny([]byte(oo.String()), w)
	default:
		return 0, NotMarshalable{o}
	}
}

func packUIntBE(i uint64, typ byte, w io.Writer) (n int64, err error) {
	packedInt := [8]byte{
		byte(i >> 56),
		byte(i >> 48),
		byte(i >> 40),
		byte(i >> 32),
		byte(i >> 24),
		byte(i >> 16),
		byte(i >> 8),
		byte(i),
	}

	off := 7

	for i := 0; i < 8; i++ {
		if packedInt[i] != 0 {
			off = i
			break
		}
	}

	buf := [9]byte{typ | byte(7-off)}
	copy(buf[1:], packedInt[off:])

	m, errWr := w.Write(buf[:9-off])
	return int64(m), errWr
}

func packFloatBE(f uint64, typ byte, w io.Writer) (n int64, err error) {
	packedFloat := [8]byte{
		byte(f >> 56),
		byte(f >> 48),
		byte(f >> 40),
		byte(f >> 32),
		byte(f >> 24),
		byte(f >> 16),
		byte(f >> 8),
		byte(f),
	}

	cap := 0

	for i := 7; i > -1; i-- {
		if packedFloat[i] != 0 {
			cap = i
			break
		}
	}

	buf := [9]byte{typ | byte(cap)}
	copy(buf[1:], packedFloat[:cap+1])

	m, errWr := w.Write(buf[:cap+2])
	return int64(m), errWr
}
