package GraniteON

type any = interface{}

const maxUInt = uint64(^uint(0))
const maxUInt8 = uint64(^uint8(0))
const maxUInt16 = uint64(^uint16(0))
const maxUInt32 = uint64(^uint32(0))

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

const scalarNil = typeNil // | 0
const scalarFalse = typeBool // | 0
const scalarTrue = typeBool | 1

const scalarUInt8 = typeUInt | typeSize8 | effSize8
const scalarUInt16 = typeUInt | typeSize16 | effSize16
const scalarUInt32 = typeUInt | typeSize32 | effSize32
const scalarUInt64 = typeUInt | typeSize64 | effSize64

const scalarInt8 = typeInt | typeSize8 | effSize8
const scalarInt16 = typeInt | typeSize16 | effSize16
const scalarInt32 = typeInt | typeSize32 | effSize32
const scalarInt64 = typeInt | typeSize64 | effSize64

const scalarFloat32 = typeFloat | typeSize32 | effSize32
const scalarFloat64 = typeFloat | typeSize64 | effSize64

const scalarString8 = typeString | effSize8
const scalarString16 = typeString | effSize16
const scalarString32 = typeString | effSize32
const scalarString64 = typeString | effSize64

const array8 = typeArray | effSize8
const array16 = typeArray | effSize16
const array32 = typeArray | effSize32
const array64 = typeArray | effSize64

const dict8 = typeDict | effSize8
const dict16 = typeDict | effSize16
const dict32 = typeDict | effSize32
const dict64 = typeDict | effSize64
