package GraniteON

type any = interface{}

const maxUInt = uint64(^uint(0))
const maxUInt32 = uint64(^uint32(0))

const typeBits = byte(7 << 5)
const typeSizeBits = byte(3 << 3)
const effSizeBits = byte(7)

const sizeBits = typeSizeBits | effSizeBits
const typeAndSizeBits = typeBits | typeSizeBits

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

const typeUInt8 = typeUInt | typeSize8
const typeUInt16 = typeUInt | typeSize16
const typeUInt32 = typeUInt | typeSize32
const typeUInt64 = typeUInt | typeSize64

const scalar8UInt8 = typeUInt8 | effSize8
const scalar8UInt16 = typeUInt16 | effSize8
const scalar16UInt16 = typeUInt16 | effSize16
const scalar8UInt32 = typeUInt32 | effSize8
const scalar32UInt32 = typeUInt32 | effSize32
const scalar8UInt64 = typeUInt64 | effSize8
const scalar64UInt64 = typeUInt64 | effSize64

const typeInt8 = typeInt | typeSize8
const typeInt16 = typeInt | typeSize16
const typeInt32 = typeInt | typeSize32
const typeInt64 = typeInt | typeSize64

const scalar8Int8 = typeInt8 | effSize8
const scalar8Int16 = typeInt16 | effSize8
const scalar16Int16 = typeInt16 | effSize16
const scalar8Int32 = typeInt32 | effSize8
const scalar32Int32 = typeInt32 | effSize32
const scalar8Int64 = typeInt64 | effSize8
const scalar64Int64 = typeInt64 | effSize64

const typeFloat32 = typeFloat | typeSize32
const typeFloat64 = typeFloat | typeSize64

const scalar8Float32 = typeFloat32 | effSize8
const scalar16Float32 = typeFloat32 | effSize16
const scalar32Float32 = typeFloat32 | effSize32
const scalar8Float64 = typeFloat64 | effSize8
const scalar16Float64 = typeFloat64 | effSize16
const scalar64Float64 = typeFloat64 | effSize64

const string8 = typeString | effSize8
const string16 = typeString | effSize16
const string24 = typeString | effSize24
const string32 = typeString | effSize32
const string64 = typeString | effSize64

const array8 = typeArray | effSize8
const array16 = typeArray | effSize16
const array32 = typeArray | effSize32
const array64 = typeArray | effSize64

const dict8 = typeDict | effSize8
const dict16 = typeDict | effSize16
const dict32 = typeDict | effSize32
const dict64 = typeDict | effSize64
