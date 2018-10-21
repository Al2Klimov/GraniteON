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

const size8 = byte(0)
const size16 = byte(1 << 3)
const size32 = byte(2 << 3)
const size64 = byte(3 << 3)

const scalarNil = typeNil | size8
const scalarFalse = typeBool | size8
const scalarTrue = typeBool | size16

const scalarUInt8 = typeUInt | size8
const scalarUInt16 = typeUInt | size16
const scalarUInt32 = typeUInt | size32
const scalarUInt64 = typeUInt | size64

const scalarInt8 = typeInt | size8
const scalarInt16 = typeInt | size16
const scalarInt32 = typeInt | size32
const scalarInt64 = typeInt | size64

const scalarFloat32 = typeFloat | size32
const scalarFloat64 = typeFloat | size64

const scalarString8 = typeString | size8
const scalarString16 = typeString | size16
const scalarString32 = typeString | size32
const scalarString64 = typeString | size64

const array8 = typeArray | size8
const array16 = typeArray | size16
const array32 = typeArray | size32
const array64 = typeArray | size64

const dict8 = typeDict | size8
const dict16 = typeDict | size16
const dict32 = typeDict | size32
const dict64 = typeDict | size64
