package GraniteON

type any = interface{}

const cpuHas32bits = uint64(^uint(0)) <= uint64(^uint32(0))

const typeBits = byte(7 << 5)
const typeSizeBits = byte(3 << 3)
const effSizeBits = byte(7)

const typeNil = byte(0)
const typeBool = byte(1 << 5)
const typeUInt = byte(2 << 5)
const typeInt = byte(3 << 5)
const typeFloat = byte(4 << 5)
const typeString = byte(5 << 5)
const typeArray = byte(6 << 5)
const typeDict = byte(7 << 5)

const typeSize8 = byte(0)
const typeSize16 = byte(1 << 3)
const typeSize32 = byte(2 << 3)
const typeSize64 = byte(3 << 3)

const effSize8 = byte(0)
const effSize16 = byte(1)
const effSize24 = byte(2)
const effSize32 = byte(3)
const effSize40 = byte(4)
const effSize48 = byte(5)
const effSize56 = byte(6)
const effSize64 = byte(7)

const scalarNil = typeNil    // | 0
const scalarFalse = typeBool // | 0
const scalarTrue = typeBool | 1
