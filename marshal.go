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
		if maxUInt <= maxUInt32 {
			return marshalAny(uint32(oo), w)
		} else {
			return marshalAny(uint64(oo), w)
		}
	case uint8:
		buf := [2]byte{scalarUInt8, oo}
		n, errWr := w.Write(buf[:])
		return int64(n), errWr
	case uint16:
		buf := [3]byte{scalarUInt16}
		i := packUInt16BE(oo)
		copy(buf[1:], i[:])
		n, errWr := w.Write(buf[:])
		return int64(n), errWr
	case uint32:
		buf := [5]byte{scalarUInt32}
		i := packUInt32BE(oo)
		copy(buf[1:], i[:])
		n, errWr := w.Write(buf[:])
		return int64(n), errWr
	case uint64:
		buf := [9]byte{scalarUInt64}
		i := packUInt64BE(oo)
		copy(buf[1:], i[:])
		n, errWr := w.Write(buf[:])
		return int64(n), errWr
	case int:
		if maxInt <= maxInt32 {
			return marshalAny(int32(oo), w)
		} else {
			return marshalAny(int64(oo), w)
		}
	case int8:
		buf := [2]byte{scalarInt8, byte(oo)}
		n, errWr := w.Write(buf[:])
		return int64(n), errWr
	case int16:
		buf := [3]byte{scalarInt16}
		i := packUInt16BE(uint16(oo))
		copy(buf[1:], i[:])
		n, errWr := w.Write(buf[:])
		return int64(n), errWr
	case int32:
		buf := [5]byte{scalarInt32}
		i := packUInt32BE(uint32(oo))
		copy(buf[1:], i[:])
		n, errWr := w.Write(buf[:])
		return int64(n), errWr
	case int64:
		buf := [9]byte{scalarInt64}
		i := packUInt64BE(uint64(oo))
		copy(buf[1:], i[:])
		n, errWr := w.Write(buf[:])
		return int64(n), errWr
	case float32:
		buf := [5]byte{scalarFloat32}
		f := packUInt32BE(math.Float32bits(oo))
		copy(buf[1:], f[:])
		n, errWr := w.Write(buf[:])
		return int64(n), errWr
	case float64:
		buf := [9]byte{scalarFloat64}
		f := packUInt64BE(math.Float64bits(oo))
		copy(buf[1:], f[:])
		n, errWr := w.Write(buf[:])
		return int64(n), errWr
	case string:
		return marshalAny([]byte(oo), w)
	case []byte:
		buf, bufLen := packLen(typeString, len(oo))
		n, errWr1 := w.Write(buf[:bufLen])
		if errWr1 != nil {
			return int64(n), errWr1
		}

		m, errWr2 := w.Write(oo)
		return int64(n + m), errWr2
	case []any:
		var n int64

		{
			buf, bufLen := packLen(typeArray, len(oo))
			m, errWr := w.Write(buf[:bufLen])
			if errWr != nil {
				return int64(m), errWr
			}

			n = int64(m)
		}

		for _, v := range oo {
			m, errWr := marshalAny(v, w)
			n += m

			if errWr != nil {
				return n, errWr
			}
		}

		return n, nil
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

		var n int64

		{
			buf, bufLen := packLen(typeDict, len(oo))
			m, errWr := w.Write(buf[:bufLen])
			if errWr != nil {
				return int64(m), errWr
			}

			n = int64(m)
		}

		for _, k := range sortedKeys {
			{
				buf, bufLen := packLen(typeString, len(k))
				m, errWr := w.Write(buf[:bufLen])
				n += int64(m)

				if errWr != nil {
					return n, errWr
				}
			}

			{
				m, errWr := w.Write([]byte(k))
				n += int64(m)

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

		return n, nil
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

func packUInt16BE(i uint16) [2]byte {
	return [2]byte{
		byte(i >> 8),
		byte(i),
	}
}

func packUInt32BE(i uint32) [4]byte {
	return [4]byte{
		byte(i >> 24),
		byte(i >> 16),
		byte(i >> 8),
		byte(i),
	}
}

func packUInt64BE(i uint64) [8]byte {
	return [8]byte{
		byte(i >> 56),
		byte(i >> 48),
		byte(i >> 40),
		byte(i >> 32),
		byte(i >> 24),
		byte(i >> 16),
		byte(i >> 8),
		byte(i),
	}
}

func packLen(typ byte, l int) (buf [9]byte, bufLen uint8) {
	i := uint64(l)

	if i <= maxUInt8 {
		buf[0] = typ | size8
		buf[1] = byte(i)
		bufLen = 2
	} else if i <= maxUInt16 {
		buf[0] = typ | size16
		ii := packUInt16BE(uint16(i))
		copy(buf[1:], ii[:])
		bufLen = 3
	} else if i <= maxUInt32 {
		buf[0] = typ | size32
		ii := packUInt32BE(uint32(i))
		copy(buf[1:], ii[:])
		bufLen = 5
	} else {
		buf[0] = typ | size64
		ii := packUInt64BE(i)
		copy(buf[1:], ii[:])
		bufLen = 9
	}

	return
}
