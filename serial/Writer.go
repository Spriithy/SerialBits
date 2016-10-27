package serial

import (
	"math"
)

type Data []byte

func (d Data) WriteBytes(ptr int, a []byte) int {
	assertBounds(len(d) >= ptr + len(a))
	for _, v := range a {
		ptr = d.WriteByte(ptr, v)
	}
	return ptr
}

// 1 bit boolean
func (d Data) WriteBoolean(ptr int, v bool) int {
	assertBounds(len(d) >= ptr + GetSize(Boolean))
	if v {
		return d.WriteByte(ptr, 1)
	}
	return d.WriteByte(ptr, 0)
}

// 8 bits integer
func (d Data) WriteInt8(ptr int, v int8) int {
	assertBounds(len(d) >= ptr + GetSize(Int8))
	return d.WriteByte(ptr, byte(v))
}

func (d Data) WriteUInt8(ptr int, v uint8) int {
	assertBounds(len(d) >= ptr + GetSize(UInt8))
	return d.WriteByte(ptr, byte(v))
}

// 16 bits integer
func (d Data) WriteInt16(ptr int, v int16) int {
	assertBounds(len(d) >= ptr + GetSize(Int16))
	d[ptr] = byte((v >> 8) & 0xFF); ptr++
	d[ptr] = byte((v >> 0) & 0xFF); ptr++
	return ptr
}

func (d Data) WriteUInt16(ptr int, v uint16) int {
	assertBounds(len(d) >= ptr + GetSize(UInt16))
	d[ptr] = byte((v >> 8) & 0xFF); ptr++
	d[ptr] = byte((v >> 0) & 0xFF); ptr++
	return ptr
}

// 32 bits integer
func (d Data) WriteInt32(ptr int, v int32) int {
	assertBounds(len(d) >= ptr + GetSize(Int32))
	d[ptr] = byte((v >> 24) & 0xFF); ptr++
	d[ptr] = byte((v >> 16) & 0xFF); ptr++
	d[ptr] = byte((v >> 8) & 0xFF); ptr++
	d[ptr] = byte((v >> 0) & 0xFF); ptr++
	return ptr
}

func (d Data) WriteUInt32(ptr int, v uint32) int {
	assertBounds(len(d) >= ptr + GetSize(UInt32))
	d[ptr] = byte((v >> 24) & 0xFF); ptr++
	d[ptr] = byte((v >> 16) & 0xFF); ptr++
	d[ptr] = byte((v >> 8) & 0xFF); ptr++
	d[ptr] = byte((v >> 0) & 0xFF); ptr++
	return ptr
}

// 64 bits integer
func (d Data) WriteInt64(ptr int, v int64) int {
	assertBounds(len(d) >= ptr + GetSize(Int64))
	d[ptr] = byte((v >> 56) & 0xFF); ptr++
	d[ptr] = byte((v >> 48) & 0xFF); ptr++
	d[ptr] = byte((v >> 40) & 0xFF); ptr++
	d[ptr] = byte((v >> 32) & 0xFF); ptr++
	d[ptr] = byte((v >> 24) & 0xFF); ptr++
	d[ptr] = byte((v >> 16) & 0xFF); ptr++
	d[ptr] = byte((v >> 8) & 0xFF); ptr++
	d[ptr] = byte((v >> 0) & 0xFF); ptr++
	return ptr
}

func (d Data) WriteUInt64(ptr int, v uint64) int {
	assertBounds(len(d) >= ptr + GetSize(UInt64))
	d[ptr] = byte((v >> 56) & 0xFF); ptr++
	d[ptr] = byte((v >> 48) & 0xFF); ptr++
	d[ptr] = byte((v >> 40) & 0xFF); ptr++
	d[ptr] = byte((v >> 32) & 0xFF); ptr++
	d[ptr] = byte((v >> 24) & 0xFF); ptr++
	d[ptr] = byte((v >> 16) & 0xFF); ptr++
	d[ptr] = byte((v >> 8) & 0xFF); ptr++
	d[ptr] = byte((v >> 0) & 0xFF); ptr++
	return ptr
}

// 32 bits floating point
func (d Data) WriteFloat32(ptr int, v float32) int {
	// Assertion is being made upon calling WriteUInt32
	u := math.Float32bits(v)
	return d.WriteUInt32(ptr, u)
}

func (d Data) WriteFloat64(ptr int, v float64) int {
	// Assertion is being made upon calling WriteUInt32
	u := math.Float64bits(v)
	return d.WriteUInt64(ptr, u)
}

// 64 bits floating point
func (d Data) WriteComplex64(ptr int, v complex64) int {
	// Assertion is being made upon calling WriteFloat32
	ptr = d.WriteFloat32(ptr, real(v))
	return d.WriteFloat32(ptr, imag(v))
}

func (d Data) WriteComplex128(ptr int, v complex128) int {
	// Assertion is being made upon calling WriteFloat64
	ptr = d.WriteFloat64(ptr, real(v))
	return d.WriteFloat64(ptr, imag(v))
}

// Aliases
func (d Data) WriteByte(ptr int, v byte) int {
	assertBounds(len(d) >= ptr + GetSize(Byte))
	d[ptr] = v; ptr++
	return ptr
}

func (d Data) WriteRune(ptr int, v rune) int {
	assertBounds(len(d) >= ptr + GetSize(Rune)) // Even though it could be omitted
	return d.WriteInt32(ptr, int32(v))
}

// Strings
// Written using length indicator (int) at ptr destination
func (d Data) WriteString(ptr int, v string) int {
	// Here again assertion is being made when calling WriteUInt16 and WriteBytes
	ptr = d.WriteUInt16(ptr, uint16(len(v)))
	return d.WriteBytes(ptr, ([]byte)(v))
}

// One function to store any array kind, rather than implementing the same over and over
// for each type
func (d Data) WriteArrayOf(ptr int, a []interface{}) int {
	switch a[0].(type) {
	case bool:
		for _, e := range a {
			ptr = d.WriteBoolean(ptr, e.(bool))
		}
	case int8:
		for _, e := range a {
			ptr = d.WriteInt8(ptr, e.(int8))
		}
	case uint8:
		for _, e := range a {
			ptr = d.WriteUInt8(ptr, e.(uint8))
		}
	case int16:
		for _, e := range a {
			ptr = d.WriteInt16(ptr, e.(int16))
		}
	case uint16:
		for _, e := range a {
			ptr = d.WriteUInt16(ptr, e.(uint16))
		}
	case int32:
		for _, e := range a {
			ptr = d.WriteInt32(ptr, e.(int32))
		}
	case uint32:
		for _, e := range a {
			ptr = d.WriteUInt32(ptr, e.(uint32))
		}
	case int64:
		for _, e := range a {
			ptr = d.WriteInt64(ptr, e.(int64))
		}
	case uint64:
		for _, e := range a {
			ptr = d.WriteUInt64(ptr, e.(uint64))
		}
	case float32:
		for _, e := range a {
			ptr = d.WriteFloat32(ptr, e.(float32))
		}
	case float64:
		for _, e := range a {
			ptr = d.WriteFloat64(ptr, e.(float64))
		}
	case complex64:
		for _, e := range a {
			ptr = d.WriteComplex64(ptr, e.(complex64))
		}
	case complex128:
		for _, e := range a {
			ptr = d.WriteComplex128(ptr, e.(complex128))
		}
	default:
		panic("Unsupported type used in `WriteArrayOf` function!")
		return -1
	}
	return ptr
}