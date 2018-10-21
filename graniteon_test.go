package GraniteON

import (
	"bytes"
	. "github.com/Al2Klimov/go-test-utils"
	"math"
	"testing"
)

func TestNil(t *testing.T) {
	assertMarshalAndUnmarshal(t, nil, []byte{0x00}, nil)
}

func TestBool(t *testing.T) {
	assertMarshalAndUnmarshal(t, false, []byte{0x10}, false)
	assertMarshalAndUnmarshal(t, true, []byte{0x11}, true)
}

func TestUInt(t *testing.T) {
	if maxUInt <= maxUInt32 {
		assertMarshalAndUnmarshal(t, uint(0), []byte{0x22, 0x00, 0x00, 0x00, 0x00}, uint32(0))
		assertMarshalAndUnmarshal(t, uint(1), []byte{0x22, 0x00, 0x00, 0x00, 0x01}, uint32(1))
		assertMarshalAndUnmarshal(t, uint(4294967294), []byte{0x22, 0xff, 0xff, 0xff, 0xfe}, uint32(4294967294))
		assertMarshalAndUnmarshal(t, uint(4294967295), []byte{0x22, 0xff, 0xff, 0xff, 0xff}, uint32(4294967295))
	} else {
		assertMarshalAndUnmarshal(t, uint(0), []byte{0x23, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, uint64(0))
		assertMarshalAndUnmarshal(t, uint(1), []byte{0x23, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01}, uint64(1))
		assertMarshalAndUnmarshal(t, uint(18446744073709551614), []byte{0x23, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xfe}, uint64(18446744073709551614))
		assertMarshalAndUnmarshal(t, uint(18446744073709551615), []byte{0x23, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}, uint64(18446744073709551615))
	}

	assertMarshalAndUnmarshal(t, uint8(0), []byte{0x20, 0x00}, uint8(0))
	assertMarshalAndUnmarshal(t, uint8(1), []byte{0x20, 0x01}, uint8(1))
	assertMarshalAndUnmarshal(t, uint8(254), []byte{0x20, 0xfe}, uint8(254))
	assertMarshalAndUnmarshal(t, uint8(255), []byte{0x20, 0xff}, uint8(255))

	assertMarshalAndUnmarshal(t, uint16(0), []byte{0x21, 0x00, 0x00}, uint16(0))
	assertMarshalAndUnmarshal(t, uint16(1), []byte{0x21, 0x00, 0x01}, uint16(1))
	assertMarshalAndUnmarshal(t, uint16(65534), []byte{0x21, 0xff, 0xfe}, uint16(65534))
	assertMarshalAndUnmarshal(t, uint16(65535), []byte{0x21, 0xff, 0xff}, uint16(65535))

	assertMarshalAndUnmarshal(t, uint32(0), []byte{0x22, 0x00, 0x00, 0x00, 0x00}, uint32(0))
	assertMarshalAndUnmarshal(t, uint32(1), []byte{0x22, 0x00, 0x00, 0x00, 0x01}, uint32(1))
	assertMarshalAndUnmarshal(t, uint32(4294967294), []byte{0x22, 0xff, 0xff, 0xff, 0xfe}, uint32(4294967294))
	assertMarshalAndUnmarshal(t, uint32(4294967295), []byte{0x22, 0xff, 0xff, 0xff, 0xff}, uint32(4294967295))

	assertMarshalAndUnmarshal(t, uint64(0), []byte{0x23, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, uint64(0))
	assertMarshalAndUnmarshal(t, uint64(1), []byte{0x23, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01}, uint64(1))
	assertMarshalAndUnmarshal(t, uint64(18446744073709551614), []byte{0x23, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xfe}, uint64(18446744073709551614))
	assertMarshalAndUnmarshal(t, uint64(18446744073709551615), []byte{0x23, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}, uint64(18446744073709551615))
}

func TestInt(t *testing.T) {
	if maxInt <= maxInt32 {
		assertMarshalAndUnmarshal(t, int(-2147483648), []byte{0x32, 0x80, 0x00, 0x00, 0x00}, int32(-2147483648))
		assertMarshalAndUnmarshal(t, int(-2147483647), []byte{0x32, 0x80, 0x00, 0x00, 0x01}, int32(-2147483647))
		assertMarshalAndUnmarshal(t, int(-2), []byte{0x32, 0xff, 0xff, 0xff, 0xfe}, int32(-2))
		assertMarshalAndUnmarshal(t, int(-1), []byte{0x32, 0xff, 0xff, 0xff, 0xff}, int32(-1))
		assertMarshalAndUnmarshal(t, int(0), []byte{0x32, 0x00, 0x00, 0x00, 0x00}, int32(0))
		assertMarshalAndUnmarshal(t, int(1), []byte{0x32, 0x00, 0x00, 0x00, 0x01}, int32(1))
		assertMarshalAndUnmarshal(t, int(2147483646), []byte{0x32, 0x7f, 0xff, 0xff, 0xfe}, int32(2147483646))
		assertMarshalAndUnmarshal(t, int(2147483647), []byte{0x32, 0x7f, 0xff, 0xff, 0xff}, int32(2147483647))
	} else {
		assertMarshalAndUnmarshal(t, int(-9223372036854775808), []byte{0x33, 0x80, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, int64(-9223372036854775808))
		assertMarshalAndUnmarshal(t, int(-9223372036854775807), []byte{0x33, 0x80, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01}, int64(-9223372036854775807))
		assertMarshalAndUnmarshal(t, int(-2), []byte{0x33, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xfe}, int64(-2))
		assertMarshalAndUnmarshal(t, int(-1), []byte{0x33, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}, int64(-1))
		assertMarshalAndUnmarshal(t, int(0), []byte{0x33, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, int64(0))
		assertMarshalAndUnmarshal(t, int(1), []byte{0x33, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01}, int64(1))
		assertMarshalAndUnmarshal(t, int(9223372036854775806), []byte{0x33, 0x7f, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xfe}, int64(9223372036854775806))
		assertMarshalAndUnmarshal(t, int(9223372036854775807), []byte{0x33, 0x7f, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}, int64(9223372036854775807))
	}

	assertMarshalAndUnmarshal(t, int8(-128), []byte{0x30, 0x80}, int8(-128))
	assertMarshalAndUnmarshal(t, int8(-127), []byte{0x30, 0x81}, int8(-127))
	assertMarshalAndUnmarshal(t, int8(-2), []byte{0x30, 0xfe}, int8(-2))
	assertMarshalAndUnmarshal(t, int8(-1), []byte{0x30, 0xff}, int8(-1))
	assertMarshalAndUnmarshal(t, int8(0), []byte{0x30, 0x00}, int8(0))
	assertMarshalAndUnmarshal(t, int8(1), []byte{0x30, 0x01}, int8(1))
	assertMarshalAndUnmarshal(t, int8(126), []byte{0x30, 0x7e}, int8(126))
	assertMarshalAndUnmarshal(t, int8(127), []byte{0x30, 0x7f}, int8(127))

	assertMarshalAndUnmarshal(t, int16(-32768), []byte{0x31, 0x80, 0x00}, int16(-32768))
	assertMarshalAndUnmarshal(t, int16(-32767), []byte{0x31, 0x80, 0x01}, int16(-32767))
	assertMarshalAndUnmarshal(t, int16(-2), []byte{0x31, 0xff, 0xfe}, int16(-2))
	assertMarshalAndUnmarshal(t, int16(-1), []byte{0x31, 0xff, 0xff}, int16(-1))
	assertMarshalAndUnmarshal(t, int16(0), []byte{0x31, 0x00, 0x00}, int16(0))
	assertMarshalAndUnmarshal(t, int16(1), []byte{0x31, 0x00, 0x01}, int16(1))
	assertMarshalAndUnmarshal(t, int16(32766), []byte{0x31, 0x7f, 0xfe}, int16(32766))
	assertMarshalAndUnmarshal(t, int16(32767), []byte{0x31, 0x7f, 0xff}, int16(32767))

	assertMarshalAndUnmarshal(t, int32(-2147483648), []byte{0x32, 0x80, 0x00, 0x00, 0x00}, int32(-2147483648))
	assertMarshalAndUnmarshal(t, int32(-2147483647), []byte{0x32, 0x80, 0x00, 0x00, 0x01}, int32(-2147483647))
	assertMarshalAndUnmarshal(t, int32(-2), []byte{0x32, 0xff, 0xff, 0xff, 0xfe}, int32(-2))
	assertMarshalAndUnmarshal(t, int32(-1), []byte{0x32, 0xff, 0xff, 0xff, 0xff}, int32(-1))
	assertMarshalAndUnmarshal(t, int32(0), []byte{0x32, 0x00, 0x00, 0x00, 0x00}, int32(0))
	assertMarshalAndUnmarshal(t, int32(1), []byte{0x32, 0x00, 0x00, 0x00, 0x01}, int32(1))
	assertMarshalAndUnmarshal(t, int32(2147483646), []byte{0x32, 0x7f, 0xff, 0xff, 0xfe}, int32(2147483646))
	assertMarshalAndUnmarshal(t, int32(2147483647), []byte{0x32, 0x7f, 0xff, 0xff, 0xff}, int32(2147483647))

	assertMarshalAndUnmarshal(t, int64(-9223372036854775808), []byte{0x33, 0x80, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, int64(-9223372036854775808))
	assertMarshalAndUnmarshal(t, int64(-9223372036854775807), []byte{0x33, 0x80, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01}, int64(-9223372036854775807))
	assertMarshalAndUnmarshal(t, int64(-2), []byte{0x33, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xfe}, int64(-2))
	assertMarshalAndUnmarshal(t, int64(-1), []byte{0x33, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}, int64(-1))
	assertMarshalAndUnmarshal(t, int64(0), []byte{0x33, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, int64(0))
	assertMarshalAndUnmarshal(t, int64(1), []byte{0x33, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01}, int64(1))
	assertMarshalAndUnmarshal(t, int64(9223372036854775806), []byte{0x33, 0x7f, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xfe}, int64(9223372036854775806))
	assertMarshalAndUnmarshal(t, int64(9223372036854775807), []byte{0x33, 0x7f, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}, int64(9223372036854775807))
}

func TestFloat(t *testing.T) {
	assertMarshalAndUnmarshal(t, float32(-math.MaxFloat32), []byte{0x42, 0xff, 0x7f, 0xff, 0xff}, float32(-math.MaxFloat32))
	assertMarshalAndUnmarshal(t, float32(-math.MaxFloat32/2), []byte{0x42, 0xfe, 0xff, 0xff, 0xff}, float32(-math.MaxFloat32/2))
	assertMarshalAndUnmarshal(t, float32(-1), []byte{0x42, 0xbf, 0x80, 0x00, 0x00}, float32(-1))
	assertMarshalAndUnmarshal(t, float32(0), []byte{0x42, 0x00, 0x00, 0x00, 0x00}, float32(0))
	assertMarshalAndUnmarshal(t, float32(1), []byte{0x42, 0x3f, 0x80, 0x00, 0x00}, float32(1))
	assertMarshalAndUnmarshal(t, float32(math.MaxFloat32/2), []byte{0x42, 0x7e, 0xff, 0xff, 0xff}, float32(math.MaxFloat32/2))
	assertMarshalAndUnmarshal(t, float32(math.MaxFloat32), []byte{0x42, 0x7f, 0x7f, 0xff, 0xff}, float32(math.MaxFloat32))

	assertMarshalAndUnmarshal(t, float64(-math.MaxFloat64), []byte{0x43, 0xff, 0xef, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}, float64(-math.MaxFloat64))
	assertMarshalAndUnmarshal(t, float64(-math.MaxFloat64/2), []byte{0x43, 0xff, 0xdf, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}, float64(-math.MaxFloat64/2))
	assertMarshalAndUnmarshal(t, float64(-1), []byte{0x43, 0xbf, 0xf0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, float64(-1))
	assertMarshalAndUnmarshal(t, float64(0), []byte{0x43, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, float64(0))
	assertMarshalAndUnmarshal(t, float64(1), []byte{0x43, 0x3f, 0xf0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, float64(1))
	assertMarshalAndUnmarshal(t, float64(math.MaxFloat64/2), []byte{0x43, 0x7f, 0xdf, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}, float64(math.MaxFloat64/2))
	assertMarshalAndUnmarshal(t, float64(math.MaxFloat64), []byte{0x43, 0x7f, 0xef, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}, float64(math.MaxFloat64))
}

func assertMarshalAndUnmarshal(t *testing.T, object any, marshaled []byte, unmarshaled any) {
	t.Helper()

	{
		buf := &bytes.Buffer{}
		n, err := GraniteON{object}.WriteTo(buf)

		AssertCallResult(
			t,
			"GraniteON{%#v}.WriteTo(&bytes.Buffer{})",
			[]any{object},
			[]any{int64(len(marshaled)), nil},
			[]any{n, err},
		)

		AssertCallResult(
			t,
			"GraniteON{%#v}.WriteTo(&bytes.Buffer{}); buf.Bytes()",
			[]any{object},
			[]any{marshaled},
			[]any{buf.Bytes()},
		)
	}

	buf := &bytes.Buffer{}
	buf.Write(marshaled)

	um := &GraniteON{}
	n, err := um.ReadFrom(buf)

	AssertCallResult(
		t,
		"(&GraniteON{}).ReadFrom(&bytes.Buffer{%#v})",
		[]any{marshaled},
		[]any{int64(len(marshaled)), nil},
		[]any{n, err},
	)

	AssertCallResult(
		t,
		"um := &GraniteON{}; um.ReadFrom(&bytes.Buffer{%#v}); um.Object",
		[]any{marshaled},
		[]any{unmarshaled},
		[]any{um.Object},
	)
}
