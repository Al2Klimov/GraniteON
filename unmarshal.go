package GraniteON

import (
	"bytes"
	"io"
	"math"
)

func unmarshalAny(r io.Reader) (o any, n int64, err error) {
	var buf [1]byte
	var m int

	m, err = r.Read(buf[:])
	n = int64(m)

	if err != nil {
		return
	}

	switch buf[0] & typeBits {
	case typeNil:
		if buf[0]&(^typeBits) == 0 {
			o = nil
		} else {
			err = NotUnmarshalable{buf[0]}
		}
	case typeBool:
		switch buf[0] & (^typeBits) {
		case 0:
			o = false
		case 1:
			o = true
		default:
			err = NotUnmarshalable{buf[0]}
		}
	case typeUInt, typeInt:
		var i uint64
		i, m, err = unpackUIntBE(buf[0], r)
		n += int64(m)

		if err != nil {
			return
		}

		switch buf[0] & (typeBits | typeSizeBits) {
		case typeUInt | typeSize8:
			o = uint8(i)
		case typeUInt | typeSize16:
			o = uint16(i)
		case typeUInt | typeSize32:
			o = uint32(i)
		case typeUInt | typeSize64:
			o = i
		case typeInt | typeSize8:
			o = int8(uint8(i))
		case typeInt | typeSize16:
			o = int16(uint16(i))
		case typeInt | typeSize32:
			o = int32(uint32(i))
		case typeInt | typeSize64:
			o = int64(i)
		}
	case typeFloat:
		switch buf[0] & (typeBits | typeSizeBits) {
		case typeFloat | typeSize32, typeFloat | typeSize64:
			var i uint64
			i, m, err = unpackFloatBE(buf[0], r)
			n += int64(m)

			if err != nil {
				return
			}

			switch buf[0] & (typeBits | typeSizeBits) {
			case typeFloat | typeSize32:
				o = math.Float32frombits(uint32(i >> 32))
			default:
				o = math.Float64frombits(i)
			}
		default:
			err = NotUnmarshalable{buf[0]}
		}
	case typeString:
		if buf[0]&typeSizeBits == 0 {
			var l uint64
			l, m, err = unpackUIntBE(buf[0], r)
			n += int64(m)

			if err != nil {
				return
			}

			buf := make([]byte, l)
			m, err = r.Read(buf)
			n += int64(m)

			if err != nil {
				return
			}

			o = string(buf)
		} else {
			err = NotUnmarshalable{buf[0]}
		}
	case typeArray:
		if buf[0]&typeSizeBits == 0 {
			var l uint64
			l, m, err = unpackUIntBE(buf[0], r)
			n += int64(m)

			if err != nil {
				return
			}

			var m int64
			a := make([]any, l)

			for i := uint64(0); i < l; i++ {
				a[i], m, err = unmarshalAny(r)
				n += m

				if err != nil {
					return
				}
			}

			o = a
		} else {
			err = NotUnmarshalable{buf[0]}
		}
	case typeDict:
		if buf[0]&typeSizeBits == 0 {
			var l uint64
			l, m, err = unpackUIntBE(buf[0], r)
			n += int64(m)

			if err != nil {
				return
			}

			d := make(map[string]any, l)
			var m int64
			var k any
			var kkk string

			for i := uint64(0); i < l; i++ {
				k, m, err = unmarshalAny(r)
				n += m

				if err != nil {
					return
				}

				switch kk := k.(type) {
				case string:
					kkk = kk
				default:
					buf := &bytes.Buffer{}
					GraniteON{k}.WriteTo(buf)
					kkk = buf.String()
				}

				d[kkk], m, err = unmarshalAny(r)
				n += m

				if err != nil {
					return
				}
			}

			o = d
		} else {
			err = NotUnmarshalable{buf[0]}
		}
	}

	return
}

func unpackUIntBE(typ byte, r io.Reader) (i uint64, n int, err error) {
	var buf [8]byte
	if n, err = r.Read(buf[7-(typ&effSizeBits):]); err != nil {
		return
	}

	i = (uint64(buf[0]) << 56) | (uint64(buf[1]) << 48) | (uint64(buf[2]) << 40) | (uint64(buf[3]) << 32) |
		(uint64(buf[4]) << 24) | (uint64(buf[5]) << 16) | (uint64(buf[6]) << 8) | uint64(buf[7])
	return
}

func unpackFloatBE(typ byte, r io.Reader) (i uint64, n int, err error) {
	var buf [8]byte
	if n, err = r.Read(buf[:(typ&effSizeBits)+1]); err != nil {
		return
	}

	i = (uint64(buf[0]) << 56) | (uint64(buf[1]) << 48) | (uint64(buf[2]) << 40) | (uint64(buf[3]) << 32) |
		(uint64(buf[4]) << 24) | (uint64(buf[5]) << 16) | (uint64(buf[6]) << 8) | uint64(buf[7])
	return
}
