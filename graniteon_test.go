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
	assertMarshalAndUnmarshal(t, nil, []byte{scalarNil}, nil)
}

func TestBool(t *testing.T) {
	assertMarshalAndUnmarshal(t, false, []byte{scalarFalse}, false)
	assertMarshalAndUnmarshal(t, true, []byte{scalarTrue}, true)
}

func TestUInt(t *testing.T) {
	if cpuHas32bits {
		assertMarshalAndUnmarshal(t, uint(0), []byte{typeUInt | typeSize32 | effSize8, 0x00}, uint32(0))
		assertMarshalAndUnmarshal(t, uint(1), []byte{typeUInt | typeSize32 | effSize8, 0x01}, uint32(1))
		assertMarshalAndUnmarshal(t, uint(254), []byte{typeUInt | typeSize32 | effSize8, 0xfe}, uint32(254))
		assertMarshalAndUnmarshal(t, uint(255), []byte{typeUInt | typeSize32 | effSize8, 0xff}, uint32(255))
		assertMarshalAndUnmarshal(t, uint(256), []byte{typeUInt | typeSize32 | effSize16, 0x01, 0x00}, uint32(256))
		assertMarshalAndUnmarshal(t, uint(257), []byte{typeUInt | typeSize32 | effSize16, 0x01, 0x01}, uint32(257))
		assertMarshalAndUnmarshal(t, uint(65534), []byte{typeUInt | typeSize32 | effSize16, 0xff, 0xfe}, uint32(65534))
		assertMarshalAndUnmarshal(t, uint(65535), []byte{typeUInt | typeSize32 | effSize16, 0xff, 0xff}, uint32(65535))
		assertMarshalAndUnmarshal(t, uint(65536), []byte{typeUInt | typeSize32 | effSize24, 0x01, 0x00, 0x00}, uint32(65536))
		assertMarshalAndUnmarshal(t, uint(65537), []byte{typeUInt | typeSize32 | effSize24, 0x01, 0x00, 0x01}, uint32(65537))
		assertMarshalAndUnmarshal(t, uint(16777214), []byte{typeUInt | typeSize32 | effSize24, 0xff, 0xff, 0xfe}, uint32(16777214))
		assertMarshalAndUnmarshal(t, uint(16777215), []byte{typeUInt | typeSize32 | effSize24, 0xff, 0xff, 0xff}, uint32(16777215))
		assertMarshalAndUnmarshal(t, uint(16777216), []byte{typeUInt | typeSize32 | effSize32, 0x01, 0x00, 0x00, 0x00}, uint32(16777216))
		assertMarshalAndUnmarshal(t, uint(16777217), []byte{typeUInt | typeSize32 | effSize32, 0x01, 0x00, 0x00, 0x01}, uint32(16777217))
		assertMarshalAndUnmarshal(t, uint(4294967294), []byte{typeUInt | typeSize32 | effSize32, 0xff, 0xff, 0xff, 0xfe}, uint32(4294967294))
		assertMarshalAndUnmarshal(t, uint(4294967295), []byte{typeUInt | typeSize32 | effSize32, 0xff, 0xff, 0xff, 0xff}, uint32(4294967295))
	} else {
		assertMarshalAndUnmarshal(t, uint(0), []byte{typeUInt | typeSize64 | effSize8, 0x00}, uint64(0))
		assertMarshalAndUnmarshal(t, uint(1), []byte{typeUInt | typeSize64 | effSize8, 0x01}, uint64(1))
		assertMarshalAndUnmarshal(t, uint(254), []byte{typeUInt | typeSize64 | effSize8, 0xfe}, uint64(254))
		assertMarshalAndUnmarshal(t, uint(255), []byte{typeUInt | typeSize64 | effSize8, 0xff}, uint64(255))
		assertMarshalAndUnmarshal(t, uint(256), []byte{typeUInt | typeSize64 | effSize16, 0x01, 0x00}, uint64(256))
		assertMarshalAndUnmarshal(t, uint(257), []byte{typeUInt | typeSize64 | effSize16, 0x01, 0x01}, uint64(257))
		assertMarshalAndUnmarshal(t, uint(65534), []byte{typeUInt | typeSize64 | effSize16, 0xff, 0xfe}, uint64(65534))
		assertMarshalAndUnmarshal(t, uint(65535), []byte{typeUInt | typeSize64 | effSize16, 0xff, 0xff}, uint64(65535))
		assertMarshalAndUnmarshal(t, uint(65536), []byte{typeUInt | typeSize64 | effSize24, 0x01, 0x00, 0x00}, uint64(65536))
		assertMarshalAndUnmarshal(t, uint(65537), []byte{typeUInt | typeSize64 | effSize24, 0x01, 0x00, 0x01}, uint64(65537))
		assertMarshalAndUnmarshal(t, uint(16777214), []byte{typeUInt | typeSize64 | effSize24, 0xff, 0xff, 0xfe}, uint64(16777214))
		assertMarshalAndUnmarshal(t, uint(16777215), []byte{typeUInt | typeSize64 | effSize24, 0xff, 0xff, 0xff}, uint64(16777215))
		assertMarshalAndUnmarshal(t, uint(16777216), []byte{typeUInt | typeSize64 | effSize32, 0x01, 0x00, 0x00, 0x00}, uint64(16777216))
		assertMarshalAndUnmarshal(t, uint(16777217), []byte{typeUInt | typeSize64 | effSize32, 0x01, 0x00, 0x00, 0x01}, uint64(16777217))
		assertMarshalAndUnmarshal(t, uint(4294967294), []byte{typeUInt | typeSize64 | effSize32, 0xff, 0xff, 0xff, 0xfe}, uint64(4294967294))
		assertMarshalAndUnmarshal(t, uint(4294967295), []byte{typeUInt | typeSize64 | effSize32, 0xff, 0xff, 0xff, 0xff}, uint64(4294967295))
		assertMarshalAndUnmarshal(t, uint(4294967296), []byte{typeUInt | typeSize64 | effSize40, 0x01, 0x00, 0x00, 0x00, 0x00}, uint64(4294967296))
		assertMarshalAndUnmarshal(t, uint(4294967297), []byte{typeUInt | typeSize64 | effSize40, 0x01, 0x00, 0x00, 0x00, 0x01}, uint64(4294967297))
		assertMarshalAndUnmarshal(t, uint(1099511627774), []byte{typeUInt | typeSize64 | effSize40, 0xff, 0xff, 0xff, 0xff, 0xfe}, uint64(1099511627774))
		assertMarshalAndUnmarshal(t, uint(1099511627775), []byte{typeUInt | typeSize64 | effSize40, 0xff, 0xff, 0xff, 0xff, 0xff}, uint64(1099511627775))
		assertMarshalAndUnmarshal(t, uint(1099511627776), []byte{typeUInt | typeSize64 | effSize48, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00}, uint64(1099511627776))
		assertMarshalAndUnmarshal(t, uint(1099511627777), []byte{typeUInt | typeSize64 | effSize48, 0x01, 0x00, 0x00, 0x00, 0x00, 0x01}, uint64(1099511627777))
		assertMarshalAndUnmarshal(t, uint(281474976710654), []byte{typeUInt | typeSize64 | effSize48, 0xff, 0xff, 0xff, 0xff, 0xff, 0xfe}, uint64(281474976710654))
		assertMarshalAndUnmarshal(t, uint(281474976710655), []byte{typeUInt | typeSize64 | effSize48, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}, uint64(281474976710655))
		assertMarshalAndUnmarshal(t, uint(281474976710656), []byte{typeUInt | typeSize64 | effSize56, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, uint64(281474976710656))
		assertMarshalAndUnmarshal(t, uint(281474976710657), []byte{typeUInt | typeSize64 | effSize56, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01}, uint64(281474976710657))
		assertMarshalAndUnmarshal(t, uint(72057594037927934), []byte{typeUInt | typeSize64 | effSize56, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xfe}, uint64(72057594037927934))
		assertMarshalAndUnmarshal(t, uint(72057594037927935), []byte{typeUInt | typeSize64 | effSize56, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}, uint64(72057594037927935))
		assertMarshalAndUnmarshal(t, uint(72057594037927936), []byte{typeUInt | typeSize64 | effSize64, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, uint64(72057594037927936))
		assertMarshalAndUnmarshal(t, uint(72057594037927937), []byte{typeUInt | typeSize64 | effSize64, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01}, uint64(72057594037927937))
		assertMarshalAndUnmarshal(t, uint(18446744073709551614), []byte{typeUInt | typeSize64 | effSize64, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xfe}, uint64(18446744073709551614))
		assertMarshalAndUnmarshal(t, uint(18446744073709551615), []byte{typeUInt | typeSize64 | effSize64, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}, uint64(18446744073709551615))
	}

	assertMarshalAndUnmarshal(t, uint8(0), []byte{typeUInt | typeSize8 | effSize8, 0x00}, uint8(0))
	assertMarshalAndUnmarshal(t, uint8(1), []byte{typeUInt | typeSize8 | effSize8, 0x01}, uint8(1))
	assertMarshalAndUnmarshal(t, uint8(254), []byte{typeUInt | typeSize8 | effSize8, 0xfe}, uint8(254))
	assertMarshalAndUnmarshal(t, uint8(255), []byte{typeUInt | typeSize8 | effSize8, 0xff}, uint8(255))

	assertMarshalAndUnmarshal(t, uint16(0), []byte{typeUInt | typeSize16 | effSize8, 0x00}, uint16(0))
	assertMarshalAndUnmarshal(t, uint16(1), []byte{typeUInt | typeSize16 | effSize8, 0x01}, uint16(1))
	assertMarshalAndUnmarshal(t, uint16(65534), []byte{typeUInt | typeSize16 | effSize16, 0xff, 0xfe}, uint16(65534))
	assertMarshalAndUnmarshal(t, uint16(65535), []byte{typeUInt | typeSize16 | effSize16, 0xff, 0xff}, uint16(65535))

	assertMarshalAndUnmarshal(t, uint32(0), []byte{typeUInt | typeSize32 | effSize8, 0x00}, uint32(0))
	assertMarshalAndUnmarshal(t, uint32(1), []byte{typeUInt | typeSize32 | effSize8, 0x01}, uint32(1))
	assertMarshalAndUnmarshal(t, uint32(254), []byte{typeUInt | typeSize32 | effSize8, 0xfe}, uint32(254))
	assertMarshalAndUnmarshal(t, uint32(255), []byte{typeUInt | typeSize32 | effSize8, 0xff}, uint32(255))
	assertMarshalAndUnmarshal(t, uint32(256), []byte{typeUInt | typeSize32 | effSize16, 0x01, 0x00}, uint32(256))
	assertMarshalAndUnmarshal(t, uint32(257), []byte{typeUInt | typeSize32 | effSize16, 0x01, 0x01}, uint32(257))
	assertMarshalAndUnmarshal(t, uint32(65534), []byte{typeUInt | typeSize32 | effSize16, 0xff, 0xfe}, uint32(65534))
	assertMarshalAndUnmarshal(t, uint32(65535), []byte{typeUInt | typeSize32 | effSize16, 0xff, 0xff}, uint32(65535))
	assertMarshalAndUnmarshal(t, uint32(65536), []byte{typeUInt | typeSize32 | effSize24, 0x01, 0x00, 0x00}, uint32(65536))
	assertMarshalAndUnmarshal(t, uint32(65537), []byte{typeUInt | typeSize32 | effSize24, 0x01, 0x00, 0x01}, uint32(65537))
	assertMarshalAndUnmarshal(t, uint32(16777214), []byte{typeUInt | typeSize32 | effSize24, 0xff, 0xff, 0xfe}, uint32(16777214))
	assertMarshalAndUnmarshal(t, uint32(16777215), []byte{typeUInt | typeSize32 | effSize24, 0xff, 0xff, 0xff}, uint32(16777215))
	assertMarshalAndUnmarshal(t, uint32(16777216), []byte{typeUInt | typeSize32 | effSize32, 0x01, 0x00, 0x00, 0x00}, uint32(16777216))
	assertMarshalAndUnmarshal(t, uint32(16777217), []byte{typeUInt | typeSize32 | effSize32, 0x01, 0x00, 0x00, 0x01}, uint32(16777217))
	assertMarshalAndUnmarshal(t, uint32(4294967294), []byte{typeUInt | typeSize32 | effSize32, 0xff, 0xff, 0xff, 0xfe}, uint32(4294967294))
	assertMarshalAndUnmarshal(t, uint32(4294967295), []byte{typeUInt | typeSize32 | effSize32, 0xff, 0xff, 0xff, 0xff}, uint32(4294967295))

	assertMarshalAndUnmarshal(t, uint64(0), []byte{typeUInt | typeSize64 | effSize8, 0x00}, uint64(0))
	assertMarshalAndUnmarshal(t, uint64(1), []byte{typeUInt | typeSize64 | effSize8, 0x01}, uint64(1))
	assertMarshalAndUnmarshal(t, uint64(254), []byte{typeUInt | typeSize64 | effSize8, 0xfe}, uint64(254))
	assertMarshalAndUnmarshal(t, uint64(255), []byte{typeUInt | typeSize64 | effSize8, 0xff}, uint64(255))
	assertMarshalAndUnmarshal(t, uint64(256), []byte{typeUInt | typeSize64 | effSize16, 0x01, 0x00}, uint64(256))
	assertMarshalAndUnmarshal(t, uint64(257), []byte{typeUInt | typeSize64 | effSize16, 0x01, 0x01}, uint64(257))
	assertMarshalAndUnmarshal(t, uint64(65534), []byte{typeUInt | typeSize64 | effSize16, 0xff, 0xfe}, uint64(65534))
	assertMarshalAndUnmarshal(t, uint64(65535), []byte{typeUInt | typeSize64 | effSize16, 0xff, 0xff}, uint64(65535))
	assertMarshalAndUnmarshal(t, uint64(65536), []byte{typeUInt | typeSize64 | effSize24, 0x01, 0x00, 0x00}, uint64(65536))
	assertMarshalAndUnmarshal(t, uint64(65537), []byte{typeUInt | typeSize64 | effSize24, 0x01, 0x00, 0x01}, uint64(65537))
	assertMarshalAndUnmarshal(t, uint64(16777214), []byte{typeUInt | typeSize64 | effSize24, 0xff, 0xff, 0xfe}, uint64(16777214))
	assertMarshalAndUnmarshal(t, uint64(16777215), []byte{typeUInt | typeSize64 | effSize24, 0xff, 0xff, 0xff}, uint64(16777215))
	assertMarshalAndUnmarshal(t, uint64(16777216), []byte{typeUInt | typeSize64 | effSize32, 0x01, 0x00, 0x00, 0x00}, uint64(16777216))
	assertMarshalAndUnmarshal(t, uint64(16777217), []byte{typeUInt | typeSize64 | effSize32, 0x01, 0x00, 0x00, 0x01}, uint64(16777217))
	assertMarshalAndUnmarshal(t, uint64(4294967294), []byte{typeUInt | typeSize64 | effSize32, 0xff, 0xff, 0xff, 0xfe}, uint64(4294967294))
	assertMarshalAndUnmarshal(t, uint64(4294967295), []byte{typeUInt | typeSize64 | effSize32, 0xff, 0xff, 0xff, 0xff}, uint64(4294967295))
	assertMarshalAndUnmarshal(t, uint64(4294967296), []byte{typeUInt | typeSize64 | effSize40, 0x01, 0x00, 0x00, 0x00, 0x00}, uint64(4294967296))
	assertMarshalAndUnmarshal(t, uint64(4294967297), []byte{typeUInt | typeSize64 | effSize40, 0x01, 0x00, 0x00, 0x00, 0x01}, uint64(4294967297))
	assertMarshalAndUnmarshal(t, uint64(1099511627774), []byte{typeUInt | typeSize64 | effSize40, 0xff, 0xff, 0xff, 0xff, 0xfe}, uint64(1099511627774))
	assertMarshalAndUnmarshal(t, uint64(1099511627775), []byte{typeUInt | typeSize64 | effSize40, 0xff, 0xff, 0xff, 0xff, 0xff}, uint64(1099511627775))
	assertMarshalAndUnmarshal(t, uint64(1099511627776), []byte{typeUInt | typeSize64 | effSize48, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00}, uint64(1099511627776))
	assertMarshalAndUnmarshal(t, uint64(1099511627777), []byte{typeUInt | typeSize64 | effSize48, 0x01, 0x00, 0x00, 0x00, 0x00, 0x01}, uint64(1099511627777))
	assertMarshalAndUnmarshal(t, uint64(281474976710654), []byte{typeUInt | typeSize64 | effSize48, 0xff, 0xff, 0xff, 0xff, 0xff, 0xfe}, uint64(281474976710654))
	assertMarshalAndUnmarshal(t, uint64(281474976710655), []byte{typeUInt | typeSize64 | effSize48, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}, uint64(281474976710655))
	assertMarshalAndUnmarshal(t, uint64(281474976710656), []byte{typeUInt | typeSize64 | effSize56, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, uint64(281474976710656))
	assertMarshalAndUnmarshal(t, uint64(281474976710657), []byte{typeUInt | typeSize64 | effSize56, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01}, uint64(281474976710657))
	assertMarshalAndUnmarshal(t, uint64(72057594037927934), []byte{typeUInt | typeSize64 | effSize56, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xfe}, uint64(72057594037927934))
	assertMarshalAndUnmarshal(t, uint64(72057594037927935), []byte{typeUInt | typeSize64 | effSize56, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}, uint64(72057594037927935))
	assertMarshalAndUnmarshal(t, uint64(72057594037927936), []byte{typeUInt | typeSize64 | effSize64, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, uint64(72057594037927936))
	assertMarshalAndUnmarshal(t, uint64(72057594037927937), []byte{typeUInt | typeSize64 | effSize64, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01}, uint64(72057594037927937))
	assertMarshalAndUnmarshal(t, uint64(18446744073709551614), []byte{typeUInt | typeSize64 | effSize64, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xfe}, uint64(18446744073709551614))
	assertMarshalAndUnmarshal(t, uint64(18446744073709551615), []byte{typeUInt | typeSize64 | effSize64, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}, uint64(18446744073709551615))
}

func TestInt(t *testing.T) {
	if cpuHas32bits {
		assertMarshalAndUnmarshal(t, int(-2147483648), []byte{typeInt | typeSize32 | effSize32, 0x80, 0x00, 0x00, 0x00}, int32(-2147483648))
		assertMarshalAndUnmarshal(t, int(-2147483647), []byte{typeInt | typeSize32 | effSize32, 0x80, 0x00, 0x00, 0x01}, int32(-2147483647))
		assertMarshalAndUnmarshal(t, int(-2), []byte{typeInt | typeSize32 | effSize32, 0xff, 0xff, 0xff, 0xfe}, int32(-2))
		assertMarshalAndUnmarshal(t, int(-1), []byte{typeInt | typeSize32 | effSize32, 0xff, 0xff, 0xff, 0xff}, int32(-1))
		assertMarshalAndUnmarshal(t, int(0), []byte{typeInt | typeSize32 | effSize8, 0x00}, int32(0))
		assertMarshalAndUnmarshal(t, int(1), []byte{typeInt | typeSize32 | effSize8, 0x01}, int32(1))
		assertMarshalAndUnmarshal(t, int(254), []byte{typeInt | typeSize32 | effSize8, 0xfe}, int32(254))
		assertMarshalAndUnmarshal(t, int(255), []byte{typeInt | typeSize32 | effSize8, 0xff}, int32(255))
		assertMarshalAndUnmarshal(t, int(256), []byte{typeInt | typeSize32 | effSize16, 0x01, 0x00}, int32(256))
		assertMarshalAndUnmarshal(t, int(257), []byte{typeInt | typeSize32 | effSize16, 0x01, 0x01}, int32(257))
		assertMarshalAndUnmarshal(t, int(65534), []byte{typeInt | typeSize32 | effSize16, 0xff, 0xfe}, int32(65534))
		assertMarshalAndUnmarshal(t, int(65535), []byte{typeInt | typeSize32 | effSize16, 0xff, 0xff}, int32(65535))
		assertMarshalAndUnmarshal(t, int(65536), []byte{typeInt | typeSize32 | effSize24, 0x01, 0x00, 0x00}, int32(65536))
		assertMarshalAndUnmarshal(t, int(65537), []byte{typeInt | typeSize32 | effSize24, 0x01, 0x00, 0x01}, int32(65537))
		assertMarshalAndUnmarshal(t, int(16777214), []byte{typeInt | typeSize32 | effSize24, 0xff, 0xff, 0xfe}, int32(16777214))
		assertMarshalAndUnmarshal(t, int(16777215), []byte{typeInt | typeSize32 | effSize24, 0xff, 0xff, 0xff}, int32(16777215))
		assertMarshalAndUnmarshal(t, int(16777216), []byte{typeInt | typeSize32 | effSize32, 0x01, 0x00, 0x00, 0x00}, int32(16777216))
		assertMarshalAndUnmarshal(t, int(16777217), []byte{typeInt | typeSize32 | effSize32, 0x01, 0x00, 0x00, 0x01}, int32(16777217))
		assertMarshalAndUnmarshal(t, int(2147483646), []byte{typeInt | typeSize32 | effSize32, 0x7f, 0xff, 0xff, 0xfe}, int32(2147483646))
		assertMarshalAndUnmarshal(t, int(2147483647), []byte{typeInt | typeSize32 | effSize32, 0x7f, 0xff, 0xff, 0xff}, int32(2147483647))
	} else {
		assertMarshalAndUnmarshal(t, int(-9223372036854775808), []byte{typeInt | typeSize64 | effSize64, 0x80, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, int64(-9223372036854775808))
		assertMarshalAndUnmarshal(t, int(-9223372036854775807), []byte{typeInt | typeSize64 | effSize64, 0x80, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01}, int64(-9223372036854775807))
		assertMarshalAndUnmarshal(t, int(-2), []byte{typeInt | typeSize64 | effSize64, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xfe}, int64(-2))
		assertMarshalAndUnmarshal(t, int(-1), []byte{typeInt | typeSize64 | effSize64, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}, int64(-1))
		assertMarshalAndUnmarshal(t, int(0), []byte{typeInt | typeSize64 | effSize8, 0x00}, int64(0))
		assertMarshalAndUnmarshal(t, int(1), []byte{typeInt | typeSize64 | effSize8, 0x01}, int64(1))
		assertMarshalAndUnmarshal(t, int(254), []byte{typeInt | typeSize64 | effSize8, 0xfe}, int64(254))
		assertMarshalAndUnmarshal(t, int(255), []byte{typeInt | typeSize64 | effSize8, 0xff}, int64(255))
		assertMarshalAndUnmarshal(t, int(256), []byte{typeInt | typeSize64 | effSize16, 0x01, 0x00}, int64(256))
		assertMarshalAndUnmarshal(t, int(257), []byte{typeInt | typeSize64 | effSize16, 0x01, 0x01}, int64(257))
		assertMarshalAndUnmarshal(t, int(65534), []byte{typeInt | typeSize64 | effSize16, 0xff, 0xfe}, int64(65534))
		assertMarshalAndUnmarshal(t, int(65535), []byte{typeInt | typeSize64 | effSize16, 0xff, 0xff}, int64(65535))
		assertMarshalAndUnmarshal(t, int(65536), []byte{typeInt | typeSize64 | effSize24, 0x01, 0x00, 0x00}, int64(65536))
		assertMarshalAndUnmarshal(t, int(65537), []byte{typeInt | typeSize64 | effSize24, 0x01, 0x00, 0x01}, int64(65537))
		assertMarshalAndUnmarshal(t, int(16777214), []byte{typeInt | typeSize64 | effSize24, 0xff, 0xff, 0xfe}, int64(16777214))
		assertMarshalAndUnmarshal(t, int(16777215), []byte{typeInt | typeSize64 | effSize24, 0xff, 0xff, 0xff}, int64(16777215))
		assertMarshalAndUnmarshal(t, int(16777216), []byte{typeInt | typeSize64 | effSize32, 0x01, 0x00, 0x00, 0x00}, int64(16777216))
		assertMarshalAndUnmarshal(t, int(16777217), []byte{typeInt | typeSize64 | effSize32, 0x01, 0x00, 0x00, 0x01}, int64(16777217))
		assertMarshalAndUnmarshal(t, int(4294967294), []byte{typeInt | typeSize64 | effSize32, 0xff, 0xff, 0xff, 0xfe}, int64(4294967294))
		assertMarshalAndUnmarshal(t, int(4294967295), []byte{typeInt | typeSize64 | effSize32, 0xff, 0xff, 0xff, 0xff}, int64(4294967295))
		assertMarshalAndUnmarshal(t, int(4294967296), []byte{typeInt | typeSize64 | effSize40, 0x01, 0x00, 0x00, 0x00, 0x00}, int64(4294967296))
		assertMarshalAndUnmarshal(t, int(4294967297), []byte{typeInt | typeSize64 | effSize40, 0x01, 0x00, 0x00, 0x00, 0x01}, int64(4294967297))
		assertMarshalAndUnmarshal(t, int(1099511627774), []byte{typeInt | typeSize64 | effSize40, 0xff, 0xff, 0xff, 0xff, 0xfe}, int64(1099511627774))
		assertMarshalAndUnmarshal(t, int(1099511627775), []byte{typeInt | typeSize64 | effSize40, 0xff, 0xff, 0xff, 0xff, 0xff}, int64(1099511627775))
		assertMarshalAndUnmarshal(t, int(1099511627776), []byte{typeInt | typeSize64 | effSize48, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00}, int64(1099511627776))
		assertMarshalAndUnmarshal(t, int(1099511627777), []byte{typeInt | typeSize64 | effSize48, 0x01, 0x00, 0x00, 0x00, 0x00, 0x01}, int64(1099511627777))
		assertMarshalAndUnmarshal(t, int(281474976710654), []byte{typeInt | typeSize64 | effSize48, 0xff, 0xff, 0xff, 0xff, 0xff, 0xfe}, int64(281474976710654))
		assertMarshalAndUnmarshal(t, int(281474976710655), []byte{typeInt | typeSize64 | effSize48, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}, int64(281474976710655))
		assertMarshalAndUnmarshal(t, int(281474976710656), []byte{typeInt | typeSize64 | effSize56, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, int64(281474976710656))
		assertMarshalAndUnmarshal(t, int(281474976710657), []byte{typeInt | typeSize64 | effSize56, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01}, int64(281474976710657))
		assertMarshalAndUnmarshal(t, int(72057594037927934), []byte{typeInt | typeSize64 | effSize56, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xfe}, int64(72057594037927934))
		assertMarshalAndUnmarshal(t, int(72057594037927935), []byte{typeInt | typeSize64 | effSize56, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}, int64(72057594037927935))
		assertMarshalAndUnmarshal(t, int(72057594037927936), []byte{typeInt | typeSize64 | effSize64, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, int64(72057594037927936))
		assertMarshalAndUnmarshal(t, int(72057594037927937), []byte{typeInt | typeSize64 | effSize64, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01}, int64(72057594037927937))
		assertMarshalAndUnmarshal(t, int(9223372036854775806), []byte{typeInt | typeSize64 | effSize64, 0x7f, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xfe}, int64(9223372036854775806))
		assertMarshalAndUnmarshal(t, int(9223372036854775807), []byte{typeInt | typeSize64 | effSize64, 0x7f, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}, int64(9223372036854775807))
	}

	assertMarshalAndUnmarshal(t, int8(-128), []byte{typeInt | typeSize8 | effSize8, 0x80}, int8(-128))
	assertMarshalAndUnmarshal(t, int8(-127), []byte{typeInt | typeSize8 | effSize8, 0x81}, int8(-127))
	assertMarshalAndUnmarshal(t, int8(-2), []byte{typeInt | typeSize8 | effSize8, 0xfe}, int8(-2))
	assertMarshalAndUnmarshal(t, int8(-1), []byte{typeInt | typeSize8 | effSize8, 0xff}, int8(-1))
	assertMarshalAndUnmarshal(t, int8(0), []byte{typeInt | typeSize8 | effSize8, 0x00}, int8(0))
	assertMarshalAndUnmarshal(t, int8(1), []byte{typeInt | typeSize8 | effSize8, 0x01}, int8(1))
	assertMarshalAndUnmarshal(t, int8(126), []byte{typeInt | typeSize8 | effSize8, 0x7e}, int8(126))
	assertMarshalAndUnmarshal(t, int8(127), []byte{typeInt | typeSize8 | effSize8, 0x7f}, int8(127))

	assertMarshalAndUnmarshal(t, int16(-32768), []byte{typeInt | typeSize16 | effSize16, 0x80, 0x00}, int16(-32768))
	assertMarshalAndUnmarshal(t, int16(-32767), []byte{typeInt | typeSize16 | effSize16, 0x80, 0x01}, int16(-32767))
	assertMarshalAndUnmarshal(t, int16(-2), []byte{typeInt | typeSize16 | effSize16, 0xff, 0xfe}, int16(-2))
	assertMarshalAndUnmarshal(t, int16(-1), []byte{typeInt | typeSize16 | effSize16, 0xff, 0xff}, int16(-1))
	assertMarshalAndUnmarshal(t, int16(0), []byte{typeInt | typeSize16 | effSize8, 0x00}, int16(0))
	assertMarshalAndUnmarshal(t, int16(1), []byte{typeInt | typeSize16 | effSize8, 0x01}, int16(1))
	assertMarshalAndUnmarshal(t, int16(32766), []byte{typeInt | typeSize16 | effSize16, 0x7f, 0xfe}, int16(32766))
	assertMarshalAndUnmarshal(t, int16(32767), []byte{typeInt | typeSize16 | effSize16, 0x7f, 0xff}, int16(32767))

	assertMarshalAndUnmarshal(t, int32(-2147483648), []byte{typeInt | typeSize32 | effSize32, 0x80, 0x00, 0x00, 0x00}, int32(-2147483648))
	assertMarshalAndUnmarshal(t, int32(-2147483647), []byte{typeInt | typeSize32 | effSize32, 0x80, 0x00, 0x00, 0x01}, int32(-2147483647))
	assertMarshalAndUnmarshal(t, int32(-2), []byte{typeInt | typeSize32 | effSize32, 0xff, 0xff, 0xff, 0xfe}, int32(-2))
	assertMarshalAndUnmarshal(t, int32(-1), []byte{typeInt | typeSize32 | effSize32, 0xff, 0xff, 0xff, 0xff}, int32(-1))
	assertMarshalAndUnmarshal(t, int32(0), []byte{typeInt | typeSize32 | effSize8, 0x00}, int32(0))
	assertMarshalAndUnmarshal(t, int32(1), []byte{typeInt | typeSize32 | effSize8, 0x01}, int32(1))
	assertMarshalAndUnmarshal(t, int32(254), []byte{typeInt | typeSize32 | effSize8, 0xfe}, int32(254))
	assertMarshalAndUnmarshal(t, int32(255), []byte{typeInt | typeSize32 | effSize8, 0xff}, int32(255))
	assertMarshalAndUnmarshal(t, int32(256), []byte{typeInt | typeSize32 | effSize16, 0x01, 0x00}, int32(256))
	assertMarshalAndUnmarshal(t, int32(257), []byte{typeInt | typeSize32 | effSize16, 0x01, 0x01}, int32(257))
	assertMarshalAndUnmarshal(t, int32(65534), []byte{typeInt | typeSize32 | effSize16, 0xff, 0xfe}, int32(65534))
	assertMarshalAndUnmarshal(t, int32(65535), []byte{typeInt | typeSize32 | effSize16, 0xff, 0xff}, int32(65535))
	assertMarshalAndUnmarshal(t, int32(65536), []byte{typeInt | typeSize32 | effSize24, 0x01, 0x00, 0x00}, int32(65536))
	assertMarshalAndUnmarshal(t, int32(65537), []byte{typeInt | typeSize32 | effSize24, 0x01, 0x00, 0x01}, int32(65537))
	assertMarshalAndUnmarshal(t, int32(16777214), []byte{typeInt | typeSize32 | effSize24, 0xff, 0xff, 0xfe}, int32(16777214))
	assertMarshalAndUnmarshal(t, int32(16777215), []byte{typeInt | typeSize32 | effSize24, 0xff, 0xff, 0xff}, int32(16777215))
	assertMarshalAndUnmarshal(t, int32(16777216), []byte{typeInt | typeSize32 | effSize32, 0x01, 0x00, 0x00, 0x00}, int32(16777216))
	assertMarshalAndUnmarshal(t, int32(16777217), []byte{typeInt | typeSize32 | effSize32, 0x01, 0x00, 0x00, 0x01}, int32(16777217))
	assertMarshalAndUnmarshal(t, int32(2147483646), []byte{typeInt | typeSize32 | effSize32, 0x7f, 0xff, 0xff, 0xfe}, int32(2147483646))
	assertMarshalAndUnmarshal(t, int32(2147483647), []byte{typeInt | typeSize32 | effSize32, 0x7f, 0xff, 0xff, 0xff}, int32(2147483647))

	assertMarshalAndUnmarshal(t, int64(-9223372036854775808), []byte{typeInt | typeSize64 | effSize64, 0x80, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, int64(-9223372036854775808))
	assertMarshalAndUnmarshal(t, int64(-9223372036854775807), []byte{typeInt | typeSize64 | effSize64, 0x80, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01}, int64(-9223372036854775807))
	assertMarshalAndUnmarshal(t, int64(-2), []byte{typeInt | typeSize64 | effSize64, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xfe}, int64(-2))
	assertMarshalAndUnmarshal(t, int64(-1), []byte{typeInt | typeSize64 | effSize64, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}, int64(-1))
	assertMarshalAndUnmarshal(t, int64(0), []byte{typeInt | typeSize64 | effSize8, 0x00}, int64(0))
	assertMarshalAndUnmarshal(t, int64(1), []byte{typeInt | typeSize64 | effSize8, 0x01}, int64(1))
	assertMarshalAndUnmarshal(t, int64(254), []byte{typeInt | typeSize64 | effSize8, 0xfe}, int64(254))
	assertMarshalAndUnmarshal(t, int64(255), []byte{typeInt | typeSize64 | effSize8, 0xff}, int64(255))
	assertMarshalAndUnmarshal(t, int64(256), []byte{typeInt | typeSize64 | effSize16, 0x01, 0x00}, int64(256))
	assertMarshalAndUnmarshal(t, int64(257), []byte{typeInt | typeSize64 | effSize16, 0x01, 0x01}, int64(257))
	assertMarshalAndUnmarshal(t, int64(65534), []byte{typeInt | typeSize64 | effSize16, 0xff, 0xfe}, int64(65534))
	assertMarshalAndUnmarshal(t, int64(65535), []byte{typeInt | typeSize64 | effSize16, 0xff, 0xff}, int64(65535))
	assertMarshalAndUnmarshal(t, int64(65536), []byte{typeInt | typeSize64 | effSize24, 0x01, 0x00, 0x00}, int64(65536))
	assertMarshalAndUnmarshal(t, int64(65537), []byte{typeInt | typeSize64 | effSize24, 0x01, 0x00, 0x01}, int64(65537))
	assertMarshalAndUnmarshal(t, int64(16777214), []byte{typeInt | typeSize64 | effSize24, 0xff, 0xff, 0xfe}, int64(16777214))
	assertMarshalAndUnmarshal(t, int64(16777215), []byte{typeInt | typeSize64 | effSize24, 0xff, 0xff, 0xff}, int64(16777215))
	assertMarshalAndUnmarshal(t, int64(16777216), []byte{typeInt | typeSize64 | effSize32, 0x01, 0x00, 0x00, 0x00}, int64(16777216))
	assertMarshalAndUnmarshal(t, int64(16777217), []byte{typeInt | typeSize64 | effSize32, 0x01, 0x00, 0x00, 0x01}, int64(16777217))
	assertMarshalAndUnmarshal(t, int64(4294967294), []byte{typeInt | typeSize64 | effSize32, 0xff, 0xff, 0xff, 0xfe}, int64(4294967294))
	assertMarshalAndUnmarshal(t, int64(4294967295), []byte{typeInt | typeSize64 | effSize32, 0xff, 0xff, 0xff, 0xff}, int64(4294967295))
	assertMarshalAndUnmarshal(t, int64(4294967296), []byte{typeInt | typeSize64 | effSize40, 0x01, 0x00, 0x00, 0x00, 0x00}, int64(4294967296))
	assertMarshalAndUnmarshal(t, int64(4294967297), []byte{typeInt | typeSize64 | effSize40, 0x01, 0x00, 0x00, 0x00, 0x01}, int64(4294967297))
	assertMarshalAndUnmarshal(t, int64(1099511627774), []byte{typeInt | typeSize64 | effSize40, 0xff, 0xff, 0xff, 0xff, 0xfe}, int64(1099511627774))
	assertMarshalAndUnmarshal(t, int64(1099511627775), []byte{typeInt | typeSize64 | effSize40, 0xff, 0xff, 0xff, 0xff, 0xff}, int64(1099511627775))
	assertMarshalAndUnmarshal(t, int64(1099511627776), []byte{typeInt | typeSize64 | effSize48, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00}, int64(1099511627776))
	assertMarshalAndUnmarshal(t, int64(1099511627777), []byte{typeInt | typeSize64 | effSize48, 0x01, 0x00, 0x00, 0x00, 0x00, 0x01}, int64(1099511627777))
	assertMarshalAndUnmarshal(t, int64(281474976710654), []byte{typeInt | typeSize64 | effSize48, 0xff, 0xff, 0xff, 0xff, 0xff, 0xfe}, int64(281474976710654))
	assertMarshalAndUnmarshal(t, int64(281474976710655), []byte{typeInt | typeSize64 | effSize48, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}, int64(281474976710655))
	assertMarshalAndUnmarshal(t, int64(281474976710656), []byte{typeInt | typeSize64 | effSize56, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, int64(281474976710656))
	assertMarshalAndUnmarshal(t, int64(281474976710657), []byte{typeInt | typeSize64 | effSize56, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01}, int64(281474976710657))
	assertMarshalAndUnmarshal(t, int64(72057594037927934), []byte{typeInt | typeSize64 | effSize56, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xfe}, int64(72057594037927934))
	assertMarshalAndUnmarshal(t, int64(72057594037927935), []byte{typeInt | typeSize64 | effSize56, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}, int64(72057594037927935))
	assertMarshalAndUnmarshal(t, int64(72057594037927936), []byte{typeInt | typeSize64 | effSize64, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, int64(72057594037927936))
	assertMarshalAndUnmarshal(t, int64(72057594037927937), []byte{typeInt | typeSize64 | effSize64, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01}, int64(72057594037927937))
	assertMarshalAndUnmarshal(t, int64(9223372036854775806), []byte{typeInt | typeSize64 | effSize64, 0x7f, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xfe}, int64(9223372036854775806))
	assertMarshalAndUnmarshal(t, int64(9223372036854775807), []byte{typeInt | typeSize64 | effSize64, 0x7f, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}, int64(9223372036854775807))
}

func TestFloat(t *testing.T) {
	assertMarshalAndUnmarshal(t, float32(-math.MaxFloat32), []byte{typeFloat | typeSize32 | effSize32, 0xff, 0x7f, 0xff, 0xff}, float32(-math.MaxFloat32))
	assertMarshalAndUnmarshal(t, float32(-math.MaxFloat32/2), []byte{typeFloat | typeSize32 | effSize32, 0xfe, 0xff, 0xff, 0xff}, float32(-math.MaxFloat32/2))
	assertMarshalAndUnmarshal(t, float32(-1), []byte{typeFloat | typeSize32 | effSize16, 0xbf, 0x80}, float32(-1))
	assertMarshalAndUnmarshal(t, float32(0), []byte{typeFloat | typeSize32 | effSize8, 0x00}, float32(0))
	assertMarshalAndUnmarshal(t, float32(1), []byte{typeFloat | typeSize32 | effSize16, 0x3f, 0x80}, float32(1))
	assertMarshalAndUnmarshal(t, float32(math.MaxFloat32/2), []byte{typeFloat | typeSize32 | effSize32, 0x7e, 0xff, 0xff, 0xff}, float32(math.MaxFloat32/2))
	assertMarshalAndUnmarshal(t, float32(math.MaxFloat32), []byte{typeFloat | typeSize32 | effSize32, 0x7f, 0x7f, 0xff, 0xff}, float32(math.MaxFloat32))

	assertMarshalAndUnmarshal(t, float64(-math.MaxFloat64), []byte{typeFloat | typeSize64 | effSize64, 0xff, 0xef, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}, float64(-math.MaxFloat64))
	assertMarshalAndUnmarshal(t, float64(-math.MaxFloat64/2), []byte{typeFloat | typeSize64 | effSize64, 0xff, 0xdf, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}, float64(-math.MaxFloat64/2))
	assertMarshalAndUnmarshal(t, float64(-1), []byte{typeFloat | typeSize64 | effSize16, 0xbf, 0xf0}, float64(-1))
	assertMarshalAndUnmarshal(t, float64(0), []byte{typeFloat | typeSize64 | effSize8, 0x00}, float64(0))
	assertMarshalAndUnmarshal(t, float64(1), []byte{typeFloat | typeSize64 | effSize16, 0x3f, 0xf0}, float64(1))
	assertMarshalAndUnmarshal(t, float64(math.MaxFloat64/2), []byte{typeFloat | typeSize64 | effSize64, 0x7f, 0xdf, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}, float64(math.MaxFloat64/2))
	assertMarshalAndUnmarshal(t, float64(math.MaxFloat64), []byte{typeFloat | typeSize64 | effSize64, 0x7f, 0xef, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}, float64(math.MaxFloat64))
}

func TestString(t *testing.T) {
	assertMarshalAndUnmarshal(t, "", []byte{typeString | effSize8, 0x00}, "")
	assertMarshalAndUnmarshal(t, []byte{}, []byte{typeString | effSize8, 0x00}, "")
	assertMarshalAndUnmarshal(t, "x", []byte{typeString | effSize8, 0x01, 'x'}, "x")
	assertMarshalAndUnmarshal(t, []byte{'x'}, []byte{typeString | effSize8, 0x01, 'x'}, "x")

	{
		buf := [256]byte{typeString | effSize8, 0xfe}
		rand.Read(buf[2:])

		s := string(buf[2:])
		assertMarshalAndUnmarshal(t, s, buf[:], s)
		assertMarshalAndUnmarshal(t, buf[2:], buf[:], s)
	}

	{
		buf := [257]byte{typeString | effSize8, 0xff}
		rand.Read(buf[2:])

		s := string(buf[2:])
		assertMarshalAndUnmarshal(t, s, buf[:], s)
		assertMarshalAndUnmarshal(t, buf[2:], buf[:], s)
	}

	{
		buf := [259]byte{typeString | effSize16, 0x01, 0x00}
		rand.Read(buf[3:])

		s := string(buf[3:])
		assertMarshalAndUnmarshal(t, s, buf[:], s)
		assertMarshalAndUnmarshal(t, buf[3:], buf[:], s)
	}

	{
		buf := [260]byte{typeString | effSize16, 0x01, 0x01}
		rand.Read(buf[3:])

		s := string(buf[3:])
		assertMarshalAndUnmarshal(t, s, buf[:], s)
		assertMarshalAndUnmarshal(t, buf[3:], buf[:], s)
	}

	{
		buf := [65537]byte{typeString | effSize16, 0xff, 0xfe}
		rand.Read(buf[3:])

		s := string(buf[3:])
		assertMarshalAndUnmarshal(t, s, buf[:], s)
		assertMarshalAndUnmarshal(t, buf[3:], buf[:], s)
	}

	{
		buf := [65538]byte{typeString | effSize16, 0xff, 0xff}
		rand.Read(buf[3:])

		s := string(buf[3:])
		assertMarshalAndUnmarshal(t, s, buf[:], s)
		assertMarshalAndUnmarshal(t, buf[3:], buf[:], s)
	}

	{
		buf := [65540]byte{typeString | effSize24, 0x01, 0x00, 0x00}
		rand.Read(buf[4:])

		s := string(buf[4:])
		assertMarshalAndUnmarshal(t, s, buf[:], s)
		assertMarshalAndUnmarshal(t, buf[4:], buf[:], s)
	}

	{
		buf := [65541]byte{typeString | effSize24, 0x01, 0x00, 0x01}
		rand.Read(buf[4:])

		s := string(buf[4:])
		assertMarshalAndUnmarshal(t, s, buf[:], s)
		assertMarshalAndUnmarshal(t, buf[4:], buf[:], s)
	}
}

func TestArray(t *testing.T) {
	assertMarshalAndUnmarshal(t, []any{}, []byte{typeArray | effSize8, 0x00}, []any{})
	assertMarshalAndUnmarshal(t, []any{nil}, []byte{typeArray | effSize8, 0x01, scalarNil}, []any{nil})
	assertMarshalAndUnmarshal(t, []any{false}, []byte{typeArray | effSize8, 0x01, scalarFalse}, []any{false})
	assertMarshalAndUnmarshal(t, []any{true}, []byte{typeArray | effSize8, 0x01, scalarTrue}, []any{true})
	assertMarshalAndUnmarshal(t, []any{uint16(1)}, []byte{typeArray | effSize8, 0x01, typeUInt | typeSize16 | effSize8, 0x01}, []any{uint16(1)})
	assertMarshalAndUnmarshal(t, []any{int32(2)}, []byte{typeArray | effSize8, 0x01, typeInt | typeSize32 | effSize8, 0x02}, []any{int32(2)})
	assertMarshalAndUnmarshal(t, []any{float64(3)}, []byte{typeArray | effSize8, 0x01, typeFloat | typeSize64 | effSize16, 0x40, 0x08}, []any{float64(3)})
	assertMarshalAndUnmarshal(t, []any{"x"}, []byte{typeArray | effSize8, 0x01, typeString | effSize8, 0x01, 'x'}, []any{"x"})
	assertMarshalAndUnmarshal(t, []any{[]any{}}, []byte{typeArray | effSize8, 0x01, typeArray | effSize8, 0x00}, []any{[]any{}})
	assertMarshalAndUnmarshal(t, []any{map[string]any{}}, []byte{typeArray | effSize8, 0x01, typeDict | effSize8, 0x00}, []any{map[string]any{}})

	assertMarshalAndUnmarshal(
		t,
		[]any{map[string]any{}, []any{}, "x", float64(3), int32(2), uint16(1), true, false, nil},
		[]byte{
			typeArray | effSize8, 0x09,
			typeDict | effSize8, 0x00,
			typeArray | effSize8, 0x00,
			typeString | effSize8, 0x01, 'x',
			typeFloat | typeSize64 | effSize16, 0x40, 0x08,
			typeInt | typeSize32 | effSize8, 0x02,
			typeUInt | typeSize16 | effSize8, 0x01,
			scalarTrue,
			scalarFalse,
			scalarNil,
		},
		[]any{map[string]any{}, []any{}, "x", float64(3), int32(2), uint16(1), true, false, nil},
	)
}

func TestDict(t *testing.T) {
	assertMarshalAndUnmarshal(t, map[string]any{}, []byte{typeDict | effSize8, 0x00}, map[string]any{})
	assertMarshalAndUnmarshal(t, map[string]any{"nil": nil}, []byte{typeDict | effSize8, 0x01, typeString | effSize8, 0x03, 'n', 'i', 'l', scalarNil}, map[string]any{"nil": nil})
	assertMarshalAndUnmarshal(t, map[string]any{"false": false}, []byte{typeDict | effSize8, 0x01, typeString | effSize8, 0x05, 'f', 'a', 'l', 's', 'e', scalarFalse}, map[string]any{"false": false})
	assertMarshalAndUnmarshal(t, map[string]any{"true": true}, []byte{typeDict | effSize8, 0x01, typeString | effSize8, 0x04, 't', 'r', 'u', 'e', scalarTrue}, map[string]any{"true": true})
	assertMarshalAndUnmarshal(t, map[string]any{"uint16(1)": uint16(1)}, []byte{typeDict | effSize8, 0x01, typeString | effSize8, 0x09, 'u', 'i', 'n', 't', '1', '6', '(', '1', ')', typeUInt | typeSize16 | effSize8, 0x01}, map[string]any{"uint16(1)": uint16(1)})
	assertMarshalAndUnmarshal(t, map[string]any{"int32(2)": int32(2)}, []byte{typeDict | effSize8, 0x01, typeString | effSize8, 0x08, 'i', 'n', 't', '3', '2', '(', '2', ')', typeInt | typeSize32 | effSize8, 0x02}, map[string]any{"int32(2)": int32(2)})
	assertMarshalAndUnmarshal(t, map[string]any{"float64(3)": float64(3)}, []byte{typeDict | effSize8, 0x01, typeString | effSize8, 0x0a, 'f', 'l', 'o', 'a', 't', '6', '4', '(', '3', ')', typeFloat | typeSize64 | effSize16, 0x40, 0x08}, map[string]any{"float64(3)": float64(3)})
	assertMarshalAndUnmarshal(t, map[string]any{`"x"`: "x"}, []byte{typeDict | effSize8, 0x01, typeString | effSize8, 0x03, '"', 'x', '"', typeString | effSize8, 0x01, 'x'}, map[string]any{`"x"`: "x"})
	assertMarshalAndUnmarshal(t, map[string]any{"[]any{}": []any{}}, []byte{typeDict | effSize8, 0x01, typeString | effSize8, 0x07, '[', ']', 'a', 'n', 'y', '{', '}', typeArray | effSize8, 0x00}, map[string]any{"[]any{}": []any{}})
	assertMarshalAndUnmarshal(t, map[string]any{"map[string]any{}": map[string]any{}}, []byte{typeDict | effSize8, 0x01, typeString | effSize8, 0x10, 'm', 'a', 'p', '[', 's', 't', 'r', 'i', 'n', 'g', ']', 'a', 'n', 'y', '{', '}', typeDict | effSize8, 0x00}, map[string]any{"map[string]any{}": map[string]any{}})

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
			typeDict | effSize8, 0x09,
			typeString | effSize8, 0x03, '"', 'x', '"', typeString | effSize8, 0x01, 'x',
			typeString | effSize8, 0x07, '[', ']', 'a', 'n', 'y', '{', '}', typeArray | effSize8, 0x00,
			typeString | effSize8, 0x05, 'f', 'a', 'l', 's', 'e', scalarFalse,
			typeString | effSize8, 0x0a, 'f', 'l', 'o', 'a', 't', '6', '4', '(', '3', ')', typeFloat | typeSize64 | effSize16, 0x40, 0x08,
			typeString | effSize8, 0x08, 'i', 'n', 't', '3', '2', '(', '2', ')', typeInt | typeSize32 | effSize8, 0x02,
			typeString | effSize8, 0x10, 'm', 'a', 'p', '[', 's', 't', 'r', 'i', 'n', 'g', ']', 'a', 'n', 'y', '{', '}', typeDict | effSize8, 0x00,
			typeString | effSize8, 0x03, 'n', 'i', 'l', scalarNil,
			typeString | effSize8, 0x04, 't', 'r', 'u', 'e', scalarTrue,
			typeString | effSize8, 0x09, 'u', 'i', 'n', 't', '1', '6', '(', '1', ')', typeUInt | typeSize16 | effSize8, 0x01,
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
	assertMarshalAndUnmarshal(t, cat{"has mouse"}, []byte{typeString | effSize8, 0x09, 'h', 'a', 's', ' ', 'm', 'o', 'u', 's', 'e'}, "has mouse")
	assertMarshalAndUnmarshal(t, lolcat{true}, []byte{typeString | effSize8, 0x12, 'I', ' ', 'h', 'a', 's', ' ', 'c', 'h', 'e', 'e', 'z', 'b', 'u', 'r', 'g', 'e', 'r', '!'}, "I has cheezburger!")
	assertMarshalAndUnmarshal(t, grumpycat{`-.-"`}, []byte{typeDict | effSize8, 0x01, typeString | effSize8, 0x04, 'm', 'o', 't', 'd', typeString | effSize8, 0x04, '-', '.', '-', '"'}, map[string]any{"motd": `-.-"`})

	assertUnmarshal(t, []byte{typeDict | effSize8, 0x01, scalarNil, scalarNil}, map[string]any{string([]byte{0x00}): nil})
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

	assertUnmarshalFromIncomplete(t, []byte{typeUInt | typeSize8 | effSize8})
	assertUnmarshalFromIncomplete(t, []byte{typeUInt | typeSize16 | effSize16})
	assertUnmarshalFromIncomplete(t, []byte{typeUInt | typeSize32 | effSize32})
	assertUnmarshalFromIncomplete(t, []byte{typeUInt | typeSize64 | effSize64})

	assertUnmarshalFromIncomplete(t, []byte{typeInt | typeSize8 | effSize8})
	assertUnmarshalFromIncomplete(t, []byte{typeInt | typeSize16 | effSize16})
	assertUnmarshalFromIncomplete(t, []byte{typeInt | typeSize32 | effSize32})
	assertUnmarshalFromIncomplete(t, []byte{typeInt | typeSize64 | effSize64})

	assertUnmarshalFromIncomplete(t, []byte{typeFloat | typeSize32 | effSize32})
	assertUnmarshalFromIncomplete(t, []byte{typeFloat | typeSize64 | effSize64})

	assertUnmarshalFromIncomplete(t, []byte{typeString | effSize8})
	assertUnmarshalFromIncomplete(t, []byte{typeString | effSize16})
	assertUnmarshalFromIncomplete(t, []byte{typeString | effSize32})
	assertUnmarshalFromIncomplete(t, []byte{typeString | effSize64})
	assertUnmarshalFromIncomplete(t, []byte{typeString | effSize8, 0x01})
	assertUnmarshalFromIncomplete(t, []byte{typeString | effSize16, 0x00, 0x01})
	assertUnmarshalFromIncomplete(t, []byte{typeString | effSize32, 0x00, 0x00, 0x00, 0x01})
	assertUnmarshalFromIncomplete(t, []byte{typeString | effSize64, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01})

	assertUnmarshalFromIncomplete(t, []byte{typeArray | effSize8})
	assertUnmarshalFromIncomplete(t, []byte{typeArray | effSize16})
	assertUnmarshalFromIncomplete(t, []byte{typeArray | effSize32})
	assertUnmarshalFromIncomplete(t, []byte{typeArray | effSize64})
	assertUnmarshalFromIncomplete(t, []byte{typeArray | effSize8, 0x01})
	assertUnmarshalFromIncomplete(t, []byte{typeArray | effSize16, 0x00, 0x01})
	assertUnmarshalFromIncomplete(t, []byte{typeArray | effSize32, 0x00, 0x00, 0x00, 0x01})
	assertUnmarshalFromIncomplete(t, []byte{typeArray | effSize64, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01})

	assertUnmarshalFromIncomplete(t, []byte{typeDict | effSize8})
	assertUnmarshalFromIncomplete(t, []byte{typeDict | effSize16})
	assertUnmarshalFromIncomplete(t, []byte{typeDict | effSize32})
	assertUnmarshalFromIncomplete(t, []byte{typeDict | effSize64})
	assertUnmarshalFromIncomplete(t, []byte{typeDict | effSize8, 0x01})
	assertUnmarshalFromIncomplete(t, []byte{typeDict | effSize16, 0x00, 0x01})
	assertUnmarshalFromIncomplete(t, []byte{typeDict | effSize32, 0x00, 0x00, 0x00, 0x01})
	assertUnmarshalFromIncomplete(t, []byte{typeDict | effSize64, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01})
	assertUnmarshalFromIncomplete(t, []byte{typeDict | effSize8, 0x01, typeString | effSize8, 0x00})
	assertUnmarshalFromIncomplete(t, []byte{typeDict | effSize16, 0x00, 0x01, typeString | effSize8, 0x00})
	assertUnmarshalFromIncomplete(t, []byte{typeDict | effSize32, 0x00, 0x00, 0x00, 0x01, typeString | effSize8, 0x00})
	assertUnmarshalFromIncomplete(t, []byte{typeDict | effSize64, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, typeString | effSize8, 0x00})

	assertUnmarshalNotUnmarshalable(t, typeNil|1)
	assertUnmarshalNotUnmarshalable(t, typeBool|2)
	assertUnmarshalNotUnmarshalable(t, typeFloat|typeSize8|effSize8)
	assertUnmarshalNotUnmarshalable(t, typeString|typeSize16|effSize8)
	assertUnmarshalNotUnmarshalable(t, typeArray|typeSize16|effSize8)
	assertUnmarshalNotUnmarshalable(t, typeDict|typeSize16|effSize8)

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

func assertUnmarshalNotUnmarshalable(t *testing.T, notUnmarshalable byte) {
	t.Helper()

	buf := &bytes.Buffer{}
	buf.Write([]byte{notUnmarshalable})
	n, err := (&GraniteON{}).ReadFrom(buf)

	AssertCallResult(
		t,
		"(&GraniteON{}).ReadFrom(&bytes.Buffer{[]byte{%d}})",
		[]any{notUnmarshalable},
		[]any{int64(1), NotUnmarshalable{notUnmarshalable}},
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
