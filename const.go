package GraniteON

type any = interface{}

const maxInt = int64(^int(0) >> 1)
const maxInt32 = int64(^int32(0) >> 1)
const maxUInt = uint64(^uint(0))
const maxUInt8 = uint64(^uint8(0))
const maxUInt16 = uint64(^uint16(0))
const maxUInt32 = uint64(^uint32(0))

const typeNil = byte(0x00)
const typeBool = byte(0x10)
const typeUInt = byte(0x20)
const typeInt = byte(0x30)
const typeFloat = byte(0x40)
const typeString = byte(0x50)
const typeArray = byte(0x60)
const typeDict = byte(0x70)

const size8 = byte(0x00)
const size16 = byte(0x01)
const size32 = byte(0x02)
const size64 = byte(0x03)

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
