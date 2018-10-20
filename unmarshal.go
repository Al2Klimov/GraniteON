package GraniteON

import (
	"bytes"
	"fmt"
	"io"
	"math"
	"strconv"
)

type NotUnmarshalable struct {
	Type byte
}

func (e NotUnmarshalable) Error() string {
	return fmt.Sprintf("non-unmarshalable type: byte(%d)", e.Type)
}

type Unmarshaler struct {
	Object any
}

func (u *Unmarshaler) ReadFrom(r io.Reader) (n int64, err error) {
	u.Object, n, err = unmarshalAny(r)
	return
}

func unmarshalAny(r io.Reader) (o any, n int64, err error) {
	var buf [1]byte
	var m int

	m, err = r.Read(buf[:])
	n = int64(m)

	if err != nil {
		return
	}

	switch buf[0] {
	case scalarNil:
		o = nil
	case scalarFalse:
		o = false
	case scalarTrue:
		o = true
	case scalarUInt8:
		var i uint8
		i, m, err = unpackUInt8(r)
		n += int64(m)

		if err != nil {
			return
		}

		o = i
	case scalarUInt16:
		var i uint16
		i, m, err = unpackUInt16BE(r)
		n += int64(m)

		if err != nil {
			return
		}

		o = i
	case scalarUInt32:
		var i uint32
		i, m, err = unpackUInt32BE(r)
		n += int64(m)

		if err != nil {
			return
		}

		o = i
	case scalarUInt64:
		var i uint64
		i, m, err = unpackUInt64BE(r)
		n += int64(m)

		if err != nil {
			return
		}

		o = i
	case scalarInt8:
		var i uint8
		i, m, err = unpackUInt8(r)
		n += int64(m)

		if err != nil {
			return
		}

		o = int8(i)
	case scalarInt16:
		var i uint16
		i, m, err = unpackUInt16BE(r)
		n += int64(m)

		if err != nil {
			return
		}

		o = int16(i)
	case scalarInt32:
		var i uint32
		i, m, err = unpackUInt32BE(r)
		n += int64(m)

		if err != nil {
			return
		}

		o = int32(i)
	case scalarInt64:
		var i uint64
		i, m, err = unpackUInt64BE(r)
		n += int64(m)

		if err != nil {
			return
		}

		o = int64(i)
	case scalarFloat32:
		var i uint32
		i, m, err = unpackUInt32BE(r)
		n += int64(m)

		if err != nil {
			return
		}

		o = math.Float32frombits(i)
	case scalarFloat64:
		var i uint64
		i, m, err = unpackUInt64BE(r)
		n += int64(m)

		if err != nil {
			return
		}

		o = math.Float64frombits(i)
	case scalarString8:
		var l uint8
		l, m, err = unpackUInt8(r)
		n += int64(m)

		if err != nil {
			return
		}

		o, m, err = unpackString(r, int(l))
		n += int64(m)
	case scalarString16:
		var l uint16
		l, m, err = unpackUInt16BE(r)
		n += int64(m)

		if err != nil {
			return
		}

		o, m, err = unpackString(r, int(l))
		n += int64(m)
	case scalarString32:
		var l uint32
		l, m, err = unpackUInt32BE(r)
		n += int64(m)

		if err != nil {
			return
		}

		o, m, err = unpackString(r, int(l))
		n += int64(m)
	case scalarString64:
		var l uint64
		l, m, err = unpackUInt64BE(r)
		n += int64(m)

		if err != nil {
			return
		}

		o, m, err = unpackString(r, int(l))
		n += int64(m)
	case array8:
		var l uint8
		l, m, err = unpackUInt8(r)
		n += int64(m)

		if err != nil {
			return
		}

		var m int64
		o, m, err = unpackArray(r, int(l))
		n += m
	case array16:
		var l uint16
		l, m, err = unpackUInt16BE(r)
		n += int64(m)

		if err != nil {
			return
		}

		var m int64
		o, m, err = unpackArray(r, int(l))
		n += m
	case array32:
		var l uint32
		l, m, err = unpackUInt32BE(r)
		n += int64(m)

		if err != nil {
			return
		}

		var m int64
		o, m, err = unpackArray(r, int(l))
		n += m
	case array64:
		var l uint64
		l, m, err = unpackUInt64BE(r)
		n += int64(m)

		if err != nil {
			return
		}

		var m int64
		o, m, err = unpackArray(r, int(l))
		n += m
	case dict8:
		var l uint8
		l, m, err = unpackUInt8(r)
		n += int64(m)

		if err != nil {
			return
		}

		var m int64
		o, m, err = unpackDict(r, int(l))
		n += m
	case dict16:
		var l uint16
		l, m, err = unpackUInt16BE(r)
		n += int64(m)

		if err != nil {
			return
		}

		var m int64
		o, m, err = unpackDict(r, int(l))
		n += m
	case dict32:
		var l uint32
		l, m, err = unpackUInt32BE(r)
		n += int64(m)

		if err != nil {
			return
		}

		var m int64
		o, m, err = unpackDict(r, int(l))
		n += m
	case dict64:
		var l uint64
		l, m, err = unpackUInt64BE(r)
		n += int64(m)

		if err != nil {
			return
		}

		var m int64
		o, m, err = unpackDict(r, int(l))
		n += m
	default:
		err = NotUnmarshalable{buf[0]}
	}

	return
}

func unpackUInt8(r io.Reader) (i uint8, n int, err error) {
	var buf [1]byte
	if n, err = r.Read(buf[:]); err != nil {
		return
	}

	i = buf[0]
	return
}

func unpackUInt16BE(r io.Reader) (i uint16, n int, err error) {
	var buf [2]byte
	if n, err = r.Read(buf[:]); err != nil {
		return
	}

	i = (uint16(buf[0]) << 8) | uint16(buf[1])
	return
}

func unpackUInt32BE(r io.Reader) (i uint32, n int, err error) {
	var buf [4]byte
	if n, err = r.Read(buf[:]); err != nil {
		return
	}

	i = (uint32(buf[0]) << 24) | (uint32(buf[1]) << 16) | (uint32(buf[2]) << 8) | uint32(buf[3])
	return
}

func unpackUInt64BE(r io.Reader) (i uint64, n int, err error) {
	var buf [8]byte
	if n, err = r.Read(buf[:]); err != nil {
		return
	}

	i = (uint64(buf[0]) << 56) | (uint64(buf[1]) << 48) | (uint64(buf[2]) << 40) | (uint64(buf[3]) << 32) |
		(uint64(buf[4]) << 24) | (uint64(buf[5]) << 16) | (uint64(buf[6]) << 8) | uint64(buf[7])
	return
}

func unpackString(r io.Reader, l int) (s any, n int, err error) {
	buf := make([]byte, l)
	if n, err = r.Read(buf); err != nil {
		return
	}

	s = string(buf)
	return
}

func unpackArray(r io.Reader, l int) (a any, n int64, err error) {
	aa := make([]any, l)
	var m int64

	for i := 0; i < l; i++ {
		aa[i], m, err = unmarshalAny(r)
		n += m

		if err != nil {
			return
		}
	}

	a = aa
	return
}

func unpackDict(r io.Reader, l int) (d any, n int64, err error) {
	dd := make(map[string]any, l)
	var m int64
	var k any
	var kkk string

	for i := 0; i < l; i++ {
		k, m, err = unmarshalAny(r)
		n += m

		if err != nil {
			return
		}

		switch kk := k.(type) {
		case nil:
			kkk = ""
		case bool:
			kkk = strconv.FormatBool(kk)
		case uint8:
			kkk = strconv.FormatUint(uint64(kk), 10)
		case uint16:
			kkk = strconv.FormatUint(uint64(kk), 10)
		case uint32:
			kkk = strconv.FormatUint(uint64(kk), 10)
		case uint64:
			kkk = strconv.FormatUint(kk, 10)
		case int8:
			kkk = strconv.FormatInt(int64(kk), 10)
		case int16:
			kkk = strconv.FormatInt(int64(kk), 10)
		case int32:
			kkk = strconv.FormatInt(int64(kk), 10)
		case int64:
			kkk = strconv.FormatInt(kk, 10)
		case float32:
			kkk = strconv.FormatFloat(float64(kk), 'g', -1, 64)
		case float64:
			kkk = strconv.FormatFloat(kk, 'g', -1, 64)
		case string:
			kkk = kk
		default:
			buf := &bytes.Buffer{}
			Marshaler{k}.WriteTo(buf)
			kkk = buf.String()
		}

		dd[kkk], m, err = unmarshalAny(r)
		n += m

		if err != nil {
			return
		}
	}

	d = dd
	return
}
