package GraniteON

import (
	"bufio"
	"bytes"
	"encoding/hex"
	"encoding/json"
	. "github.com/Al2Klimov/go-test-utils"
	"io"
	"math"
	"math/rand"
	"testing"
)

type cat struct {
	desc string
}

func (c cat) String() string {
	return c.desc
}

type lolcat struct {
	hasCheezBurger bool
}

func (l lolcat) MarshalText() (text []byte, err error) {
	if l.hasCheezBurger {
		return []byte("I has cheezburger!"), nil
	}

	return nil, hex.ErrLength
}

type grumpycat struct {
	motd string
}

func (g grumpycat) MarshalJSON() ([]byte, error) {
	if g.motd == "" {
		return nil, hex.ErrLength
	}

	return json.Marshal(map[string]any{"motd": g.motd})
}

type brokenJson struct {
	brokenJson []byte
}

func (b brokenJson) MarshalJSON() ([]byte, error) {
	return b.brokenJson, nil
}

type fullDisk struct {
	cap int
}

func (d *fullDisk) Write(p []byte) (n int, err error) {
	if len(p) > d.cap {
		n = d.cap
		err = bufio.ErrBufferFull

		d.cap = 0

		return
	}

	d.cap -= len(p)

	return len(p), nil
}

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
	if maxUInt <= maxUInt32 {
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

func TestString(t *testing.T) {
	assertMarshalAndUnmarshal(t, "", []byte{0x50, 0x00}, "")
	assertMarshalAndUnmarshal(t, []byte{}, []byte{0x50, 0x00}, "")
	assertMarshalAndUnmarshal(t, "x", []byte{0x50, 0x01, 'x'}, "x")
	assertMarshalAndUnmarshal(t, []byte{'x'}, []byte{0x50, 0x01, 'x'}, "x")

	{
		buf := [256]byte{0x50, 0xfe}
		rand.Read(buf[2:])

		s := string(buf[2:])
		assertMarshalAndUnmarshal(t, s, buf[:], s)
		assertMarshalAndUnmarshal(t, buf[2:], buf[:], s)
	}

	{
		buf := [257]byte{0x50, 0xff}
		rand.Read(buf[2:])

		s := string(buf[2:])
		assertMarshalAndUnmarshal(t, s, buf[:], s)
		assertMarshalAndUnmarshal(t, buf[2:], buf[:], s)
	}

	{
		buf := [259]byte{0x51, 0x01, 0x00}
		rand.Read(buf[3:])

		s := string(buf[3:])
		assertMarshalAndUnmarshal(t, s, buf[:], s)
		assertMarshalAndUnmarshal(t, buf[3:], buf[:], s)
	}

	{
		buf := [260]byte{0x51, 0x01, 0x01}
		rand.Read(buf[3:])

		s := string(buf[3:])
		assertMarshalAndUnmarshal(t, s, buf[:], s)
		assertMarshalAndUnmarshal(t, buf[3:], buf[:], s)
	}

	{
		buf := [65537]byte{0x51, 0xff, 0xfe}
		rand.Read(buf[3:])

		s := string(buf[3:])
		assertMarshalAndUnmarshal(t, s, buf[:], s)
		assertMarshalAndUnmarshal(t, buf[3:], buf[:], s)
	}

	{
		buf := [65538]byte{0x51, 0xff, 0xff}
		rand.Read(buf[3:])

		s := string(buf[3:])
		assertMarshalAndUnmarshal(t, s, buf[:], s)
		assertMarshalAndUnmarshal(t, buf[3:], buf[:], s)
	}

	{
		buf := [65541]byte{0x52, 0x00, 0x01, 0x00, 0x00}
		rand.Read(buf[5:])

		s := string(buf[5:])
		assertMarshalAndUnmarshal(t, s, buf[:], s)
		assertMarshalAndUnmarshal(t, buf[5:], buf[:], s)
	}

	{
		buf := [65542]byte{0x52, 0x00, 0x01, 0x00, 0x01}
		rand.Read(buf[5:])

		s := string(buf[5:])
		assertMarshalAndUnmarshal(t, s, buf[:], s)
		assertMarshalAndUnmarshal(t, buf[5:], buf[:], s)
	}
}

func TestArray(t *testing.T) {
	assertMarshalAndUnmarshal(t, []any{}, []byte{0x60, 0x00}, []any{})
	assertMarshalAndUnmarshal(t, []any{nil}, []byte{0x60, 0x01, 0x00}, []any{nil})
	assertMarshalAndUnmarshal(t, []any{false}, []byte{0x60, 0x01, 0x10}, []any{false})
	assertMarshalAndUnmarshal(t, []any{true}, []byte{0x60, 0x01, 0x11}, []any{true})
	assertMarshalAndUnmarshal(t, []any{uint16(1)}, []byte{0x60, 0x01, 0x21, 0x00, 0x01}, []any{uint16(1)})
	assertMarshalAndUnmarshal(t, []any{int32(2)}, []byte{0x60, 0x01, 0x32, 0x00, 0x00, 0x00, 0x02}, []any{int32(2)})
	assertMarshalAndUnmarshal(t, []any{float64(3)}, []byte{0x60, 0x01, 0x43, 0x40, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, []any{float64(3)})
	assertMarshalAndUnmarshal(t, []any{"x"}, []byte{0x60, 0x01, 0x50, 0x01, 'x'}, []any{"x"})
	assertMarshalAndUnmarshal(t, []any{[]any{}}, []byte{0x60, 0x01, 0x60, 0x00}, []any{[]any{}})
	assertMarshalAndUnmarshal(t, []any{map[string]any{}}, []byte{0x60, 0x01, 0x70, 0x00}, []any{map[string]any{}})

	assertMarshalAndUnmarshal(
		t,
		[]any{map[string]any{}, []any{}, "x", float64(3), int32(2), uint16(1), true, false, nil},
		[]byte{
			0x60, 0x09,
			0x70, 0x00,
			0x60, 0x00,
			0x50, 0x01, 'x',
			0x43, 0x40, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x32, 0x00, 0x00, 0x00, 0x02,
			0x21, 0x00, 0x01,
			0x11,
			0x10,
			0x00,
		},
		[]any{map[string]any{}, []any{}, "x", float64(3), int32(2), uint16(1), true, false, nil},
	)
}

func TestDict(t *testing.T) {
	assertMarshalAndUnmarshal(t, map[string]any{}, []byte{0x70, 0x00}, map[string]any{})
	assertMarshalAndUnmarshal(t, map[string]any{"nil": nil}, []byte{0x70, 0x01, 0x50, 0x03, 'n', 'i', 'l', 0x00}, map[string]any{"nil": nil})
	assertMarshalAndUnmarshal(t, map[string]any{"false": false}, []byte{0x70, 0x01, 0x50, 0x05, 'f', 'a', 'l', 's', 'e', 0x10}, map[string]any{"false": false})
	assertMarshalAndUnmarshal(t, map[string]any{"true": true}, []byte{0x70, 0x01, 0x50, 0x04, 't', 'r', 'u', 'e', 0x11}, map[string]any{"true": true})
	assertMarshalAndUnmarshal(t, map[string]any{"uint16(1)": uint16(1)}, []byte{0x70, 0x01, 0x50, 0x09, 'u', 'i', 'n', 't', '1', '6', '(', '1', ')', 0x21, 0x00, 0x01}, map[string]any{"uint16(1)": uint16(1)})
	assertMarshalAndUnmarshal(t, map[string]any{"int32(2)": int32(2)}, []byte{0x70, 0x01, 0x50, 0x08, 'i', 'n', 't', '3', '2', '(', '2', ')', 0x32, 0x00, 0x00, 0x00, 0x02}, map[string]any{"int32(2)": int32(2)})
	assertMarshalAndUnmarshal(t, map[string]any{"float64(3)": float64(3)}, []byte{0x70, 0x01, 0x50, 0x0a, 'f', 'l', 'o', 'a', 't', '6', '4', '(', '3', ')', 0x43, 0x40, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, map[string]any{"float64(3)": float64(3)})
	assertMarshalAndUnmarshal(t, map[string]any{`"x"`: "x"}, []byte{0x70, 0x01, 0x50, 0x03, '"', 'x', '"', 0x50, 0x01, 'x'}, map[string]any{`"x"`: "x"})
	assertMarshalAndUnmarshal(t, map[string]any{"[]any{}": []any{}}, []byte{0x70, 0x01, 0x50, 0x07, '[', ']', 'a', 'n', 'y', '{', '}', 0x60, 0x00}, map[string]any{"[]any{}": []any{}})
	assertMarshalAndUnmarshal(t, map[string]any{"map[string]any{}": map[string]any{}}, []byte{0x70, 0x01, 0x50, 0x10, 'm', 'a', 'p', '[', 's', 't', 'r', 'i', 'n', 'g', ']', 'a', 'n', 'y', '{', '}', 0x70, 0x00}, map[string]any{"map[string]any{}": map[string]any{}})

	assertMarshalAndUnmarshal(
		t,
		map[string]any{
			"map[string]any{}": map[string]any{},
			"[]any{}":          []any{},
			`"x"`:              "x",
			"float64(3)":       float64(3),
			"int32(2)":         int32(2),
			"uint16(1)":        uint16(1),
			"true":             true,
			"false":            false,
			"nil":              nil,
		},
		[]byte{
			0x70, 0x09,
			0x50, 0x03, '"', 'x', '"', 0x50, 0x01, 'x',
			0x50, 0x07, '[', ']', 'a', 'n', 'y', '{', '}', 0x60, 0x00,
			0x50, 0x05, 'f', 'a', 'l', 's', 'e', 0x10,
			0x50, 0x0a, 'f', 'l', 'o', 'a', 't', '6', '4', '(', '3', ')', 0x43, 0x40, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x50, 0x08, 'i', 'n', 't', '3', '2', '(', '2', ')', 0x32, 0x00, 0x00, 0x00, 0x02,
			0x50, 0x10, 'm', 'a', 'p', '[', 's', 't', 'r', 'i', 'n', 'g', ']', 'a', 'n', 'y', '{', '}', 0x70, 0x00,
			0x50, 0x03, 'n', 'i', 'l', 0x00,
			0x50, 0x04, 't', 'r', 'u', 'e', 0x11,
			0x50, 0x09, 'u', 'i', 'n', 't', '1', '6', '(', '1', ')', 0x21, 0x00, 0x01,
		},
		map[string]any{
			"map[string]any{}": map[string]any{},
			"[]any{}":          []any{},
			`"x"`:              "x",
			"float64(3)":       float64(3),
			"int32(2)":         int32(2),
			"uint16(1)":        uint16(1),
			"true":             true,
			"false":            false,
			"nil":              nil,
		},
	)
}

func TestSpecial(t *testing.T) {
	assertMarshalAndUnmarshal(t, cat{"has mouse"}, []byte{0x50, 0x09, 'h', 'a', 's', ' ', 'm', 'o', 'u', 's', 'e'}, "has mouse")
	assertMarshalAndUnmarshal(t, lolcat{true}, []byte{0x50, 0x12, 'I', ' ', 'h', 'a', 's', ' ', 'c', 'h', 'e', 'e', 'z', 'b', 'u', 'r', 'g', 'e', 'r', '!'}, "I has cheezburger!")
	assertMarshalAndUnmarshal(t, grumpycat{`-.-"`}, []byte{0x70, 0x01, 0x50, 0x04, 'm', 'o', 't', 'd', 0x50, 0x04, '-', '.', '-', '"'}, map[string]any{"motd": `-.-"`})

	assertUnmarshal(t, []byte{0x70, 0x01, 0x00, 0x00}, map[string]any{string([]byte{0x00}): nil})
}

func TestError(t *testing.T) {
	{
		n, err := GraniteON{struct{}{}}.WriteTo(nil)

		AssertCallResult(
			t,
			"GraniteON{struct{}{}}.WriteTo(nil)",
			[]any{},
			[]any{int64(0), NotMarshalable{struct{}{}}},
			[]any{n, err},
			true,
		)

		AssertCallResult(
			t,
			"NotMarshalable{struct{}{}}.Error()",
			[]any{},
			[]any{"non-marshalable value: (struct {}{})"},
			[]any{NotMarshalable{struct{}{}}.Error()},
			true,
		)

		AssertCallResult(
			t,
			"NotMarshalable{io.Writer(nil)}.Error()",
			[]any{},
			[]any{"non-marshalable value: nil(<nil>)"},
			[]any{NotMarshalable{io.Writer(nil)}.Error()},
			true,
		)
	}

	assertMarshalToFullDisk(t, "", 0)
	assertMarshalToFullDisk(t, []byte{}, 0)

	assertMarshalToFullDisk(t, []any{}, 0)
	assertMarshalToFullDisk(t, []any{nil}, 2)

	assertMarshalToFullDisk(t, map[string]any{}, 0)
	assertMarshalToFullDisk(t, map[string]any{"": nil}, 2)
	assertMarshalToFullDisk(t, map[string]any{"": nil}, 4)

	assertUnmarshalFromIncomplete(t, []byte{})

	assertUnmarshalFromIncomplete(t, []byte{0x20})
	assertUnmarshalFromIncomplete(t, []byte{0x21})
	assertUnmarshalFromIncomplete(t, []byte{0x22})
	assertUnmarshalFromIncomplete(t, []byte{0x23})

	assertUnmarshalFromIncomplete(t, []byte{0x30})
	assertUnmarshalFromIncomplete(t, []byte{0x31})
	assertUnmarshalFromIncomplete(t, []byte{0x32})
	assertUnmarshalFromIncomplete(t, []byte{0x33})

	assertUnmarshalFromIncomplete(t, []byte{0x42})
	assertUnmarshalFromIncomplete(t, []byte{0x43})

	assertUnmarshalFromIncomplete(t, []byte{0x50})
	assertUnmarshalFromIncomplete(t, []byte{0x51})
	assertUnmarshalFromIncomplete(t, []byte{0x52})
	assertUnmarshalFromIncomplete(t, []byte{0x53})
	assertUnmarshalFromIncomplete(t, []byte{0x50, 0x01})
	assertUnmarshalFromIncomplete(t, []byte{0x51, 0x00, 0x01})
	assertUnmarshalFromIncomplete(t, []byte{0x52, 0x00, 0x00, 0x00, 0x01})
	assertUnmarshalFromIncomplete(t, []byte{0x53, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01})

	assertUnmarshalFromIncomplete(t, []byte{0x60})
	assertUnmarshalFromIncomplete(t, []byte{0x61})
	assertUnmarshalFromIncomplete(t, []byte{0x62})
	assertUnmarshalFromIncomplete(t, []byte{0x63})
	assertUnmarshalFromIncomplete(t, []byte{0x60, 0x01})
	assertUnmarshalFromIncomplete(t, []byte{0x61, 0x00, 0x01})
	assertUnmarshalFromIncomplete(t, []byte{0x62, 0x00, 0x00, 0x00, 0x01})
	assertUnmarshalFromIncomplete(t, []byte{0x63, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01})

	assertUnmarshalFromIncomplete(t, []byte{0x70})
	assertUnmarshalFromIncomplete(t, []byte{0x71})
	assertUnmarshalFromIncomplete(t, []byte{0x72})
	assertUnmarshalFromIncomplete(t, []byte{0x73})
	assertUnmarshalFromIncomplete(t, []byte{0x70, 0x01})
	assertUnmarshalFromIncomplete(t, []byte{0x71, 0x00, 0x01})
	assertUnmarshalFromIncomplete(t, []byte{0x72, 0x00, 0x00, 0x00, 0x01})
	assertUnmarshalFromIncomplete(t, []byte{0x73, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01})
	assertUnmarshalFromIncomplete(t, []byte{0x70, 0x01, 0x50, 0x00})
	assertUnmarshalFromIncomplete(t, []byte{0x71, 0x00, 0x01, 0x50, 0x00})
	assertUnmarshalFromIncomplete(t, []byte{0x72, 0x00, 0x00, 0x00, 0x01, 0x50, 0x00})
	assertUnmarshalFromIncomplete(t, []byte{0x73, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x50, 0x00})

	{
		buf := &bytes.Buffer{}
		buf.Write([]byte{0xff})
		n, err := (&GraniteON{}).ReadFrom(buf)

		AssertCallResult(
			t,
			"(&GraniteON{}).ReadFrom(&bytes.Buffer{[]byte{0xff}})",
			[]any{},
			[]any{int64(1), NotUnmarshalable{0xff}},
			[]any{n, err},
			true,
		)

		AssertCallResult(
			t,
			"NotUnmarshalable{0xff}.Error()",
			[]any{},
			[]any{"non-unmarshalable type: byte(255)"},
			[]any{NotUnmarshalable{0xff}.Error()},
			true,
		)
	}

	{
		n, err := GraniteON{lolcat{false}}.WriteTo(nil)

		AssertCallResult(
			t,
			"GraniteON{lolcat{false}}.WriteTo(nil)",
			[]any{},
			[]any{int64(0), hex.ErrLength},
			[]any{n, err},
			true,
		)
	}

	{
		n, err := GraniteON{grumpycat{""}}.WriteTo(nil)

		AssertCallResult(
			t,
			`GraniteON{grumpycat{""}}.WriteTo(nil)`,
			[]any{},
			[]any{int64(0), hex.ErrLength},
			[]any{n, err},
			true,
		)
	}

	{
		n, err := GraniteON{brokenJson{[]byte{'{'}}}.WriteTo(nil)

		AssertCallResult(
			t,
			"GraniteON{brokenJson{[]byte{'{'}}}.WriteTo(nil)",
			[]any{},
			[]any{int64(0), nil},
			[]any{n, err},
			false,
		)
	}
}

func assertMarshal(t *testing.T, object any, marshaled []byte) {
	t.Helper()

	buf := &bytes.Buffer{}
	n, err := GraniteON{object}.WriteTo(buf)

	AssertCallResult(
		t,
		"GraniteON{%#v}.WriteTo(&bytes.Buffer{})",
		[]any{object},
		[]any{int64(len(marshaled)), nil},
		[]any{n, err},
		true,
	)

	AssertCallResult(
		t,
		"GraniteON{%#v}.WriteTo(&bytes.Buffer{}); buf.Bytes()",
		[]any{object},
		[]any{marshaled},
		[]any{buf.Bytes()},
		true,
	)
}

func assertUnmarshal(t *testing.T, marshaled []byte, unmarshaled any) {
	t.Helper()

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
		true,
	)

	AssertCallResult(
		t,
		"um := &GraniteON{}; um.ReadFrom(&bytes.Buffer{%#v}); um.Object",
		[]any{marshaled},
		[]any{unmarshaled},
		[]any{um.Object},
		true,
	)
}

func assertMarshalAndUnmarshal(t *testing.T, object any, marshaled []byte, unmarshaled any) {
	t.Helper()

	assertMarshal(t, object, marshaled)
	assertUnmarshal(t, marshaled, unmarshaled)
}

func assertMarshalToFullDisk(t *testing.T, object any, cap int) {
	t.Helper()

	n, err := GraniteON{object}.WriteTo(&fullDisk{cap})

	AssertCallResult(
		t,
		"GraniteON{%#v}.WriteTo(&fullDisk{%d})",
		[]any{object, cap},
		[]any{int64(cap), bufio.ErrBufferFull},
		[]any{n, err},
		true,
	)
}

func assertUnmarshalFromIncomplete(t *testing.T, incomplete []byte) {
	t.Helper()

	buf := &bytes.Buffer{}
	buf.Write(incomplete)
	n, err := (&GraniteON{}).ReadFrom(buf)

	AssertCallResult(
		t,
		"(&GraniteON{}).ReadFrom(&bytes.Buffer{%#v})",
		[]any{incomplete},
		[]any{int64(len(incomplete)), io.EOF},
		[]any{n, err},
		true,
	)
}
