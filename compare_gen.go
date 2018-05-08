package main

// THIS FILE WAS PRODUCED BY THE MSGP CODE GENERATION TOOL (github.com/dchenk/msgp).
// DO NOT EDIT.

import (
	"github.com/dchenk/msgp/msgp"
)

// MarshalMsg implements msgp.Marshaler
func (z Large) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendArrayHeader(o, uint32(len(z)))
	for za0001 := range z {
		// map header, size 7
		// string "hello"
		o = append(o, 0x87, 0xa5, 0x68, 0x65, 0x6c, 0x6c, 0x6f)
		o = msgp.AppendString(o, z[za0001].Hello)
		// string "name"
		o = append(o, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
		o = msgp.AppendString(o, z[za0001].Name)
		// string "age"
		o = append(o, 0xa3, 0x61, 0x67, 0x65)
		o = msgp.AppendInt(o, z[za0001].Age)
		// string "weight"
		o = append(o, 0xa6, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74)
		o = msgp.AppendFloat64(o, z[za0001].Weight)
		// string "height"
		o = append(o, 0xa6, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74)
		o = msgp.AppendFloat64(o, z[za0001].Height)
		// string "hobbies"
		o = append(o, 0xa7, 0x68, 0x6f, 0x62, 0x62, 0x69, 0x65, 0x73)
		o = msgp.AppendArrayHeader(o, uint32(len(z[za0001].Hobbies)))
		for za0002 := range z[za0001].Hobbies {
			o = msgp.AppendString(o, z[za0001].Hobbies[za0002])
		}
		// string "extra"
		// map header, size 4
		// string "location"
		o = append(o, 0xa5, 0x65, 0x78, 0x74, 0x72, 0x61, 0x84, 0xa8, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e)
		o = msgp.AppendString(o, z[za0001].Extra.Location)
		// string "bio"
		o = append(o, 0xa3, 0x62, 0x69, 0x6f)
		o = msgp.AppendString(o, z[za0001].Extra.Bio)
		// string "number"
		o = append(o, 0xa6, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72)
		o = msgp.AppendInt(o, z[za0001].Extra.Number)
		// string "negative_num"
		o = append(o, 0xac, 0x6e, 0x65, 0x67, 0x61, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x6e, 0x75, 0x6d)
		o = msgp.AppendInt(o, z[za0001].Extra.NegativeNum)
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Large) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0003 uint32
	zb0003, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		return
	}
	if cap((*z)) >= int(zb0003) {
		(*z) = (*z)[:zb0003]
	} else {
		(*z) = make(Large, zb0003)
	}
	for zb0001 := range *z {
		var field []byte
		_ = field
		var zb0004 uint32
		zb0004, bts, err = msgp.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
		for zb0004 > 0 {
			zb0004--
			field, bts, err = msgp.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			switch string(field) {
			case "hello":
				(*z)[zb0001].Hello, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
			case "name":
				(*z)[zb0001].Name, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
			case "age":
				(*z)[zb0001].Age, bts, err = msgp.ReadIntBytes(bts)
				if err != nil {
					return
				}
			case "weight":
				(*z)[zb0001].Weight, bts, err = msgp.ReadFloat64Bytes(bts)
				if err != nil {
					return
				}
			case "height":
				(*z)[zb0001].Height, bts, err = msgp.ReadFloat64Bytes(bts)
				if err != nil {
					return
				}
			case "hobbies":
				var zb0005 uint32
				zb0005, bts, err = msgp.ReadArrayHeaderBytes(bts)
				if err != nil {
					return
				}
				if cap((*z)[zb0001].Hobbies) >= int(zb0005) {
					(*z)[zb0001].Hobbies = ((*z)[zb0001].Hobbies)[:zb0005]
				} else {
					(*z)[zb0001].Hobbies = make([]string, zb0005)
				}
				for zb0002 := range (*z)[zb0001].Hobbies {
					(*z)[zb0001].Hobbies[zb0002], bts, err = msgp.ReadStringBytes(bts)
					if err != nil {
						return
					}
				}
			case "extra":
				var zb0006 uint32
				zb0006, bts, err = msgp.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				for zb0006 > 0 {
					zb0006--
					field, bts, err = msgp.ReadMapKeyZC(bts)
					if err != nil {
						return
					}
					switch string(field) {
					case "location":
						(*z)[zb0001].Extra.Location, bts, err = msgp.ReadStringBytes(bts)
						if err != nil {
							return
						}
					case "bio":
						(*z)[zb0001].Extra.Bio, bts, err = msgp.ReadStringBytes(bts)
						if err != nil {
							return
						}
					case "number":
						(*z)[zb0001].Extra.Number, bts, err = msgp.ReadIntBytes(bts)
						if err != nil {
							return
						}
					case "negative_num":
						(*z)[zb0001].Extra.NegativeNum, bts, err = msgp.ReadIntBytes(bts)
						if err != nil {
							return
						}
					default:
						bts, err = msgp.Skip(bts)
						if err != nil {
							return
						}
					}
				}
			default:
				bts, err = msgp.Skip(bts)
				if err != nil {
					return
				}
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z Large) Msgsize() (s int) {
	s = msgp.ArrayHeaderSize
	for zb0007 := range z {
		s += 1 + 6 + msgp.StringPrefixSize + len(z[zb0007].Hello) + 5 + msgp.StringPrefixSize + len(z[zb0007].Name) + 4 + msgp.IntSize + 7 + msgp.Float64Size + 7 + msgp.Float64Size + 8 + msgp.ArrayHeaderSize
		for zb0008 := range z[zb0007].Hobbies {
			s += msgp.StringPrefixSize + len(z[zb0007].Hobbies[zb0008])
		}
		s += 6 + 1 + 9 + msgp.StringPrefixSize + len(z[zb0007].Extra.Location) + 4 + msgp.StringPrefixSize + len(z[zb0007].Extra.Bio) + 7 + msgp.IntSize + 13 + msgp.IntSize
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z Medium) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendArrayHeader(o, uint32(len(z)))
	for za0001 := range z {
		// map header, size 8
		// string "hello"
		o = append(o, 0x88, 0xa5, 0x68, 0x65, 0x6c, 0x6c, 0x6f)
		o = msgp.AppendString(o, z[za0001].Hello)
		// string "name"
		o = append(o, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
		o = msgp.AppendString(o, z[za0001].Name)
		// string "age"
		o = append(o, 0xa3, 0x61, 0x67, 0x65)
		o = msgp.AppendInt(o, z[za0001].Age)
		// string "weight"
		o = append(o, 0xa6, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74)
		o = msgp.AppendFloat64(o, z[za0001].Weight)
		// string "height"
		o = append(o, 0xa6, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74)
		o = msgp.AppendFloat64(o, z[za0001].Height)
		// string "hobbies"
		o = append(o, 0xa7, 0x68, 0x6f, 0x62, 0x62, 0x69, 0x65, 0x73)
		o = msgp.AppendArrayHeader(o, uint32(len(z[za0001].Hobbies)))
		for za0002 := range z[za0001].Hobbies {
			o = msgp.AppendString(o, z[za0001].Hobbies[za0002])
		}
		// string "cool"
		o = append(o, 0xa4, 0x63, 0x6f, 0x6f, 0x6c)
		o = msgp.AppendBool(o, z[za0001].Cool)
		// string "extra"
		// map header, size 2
		// string "location"
		o = append(o, 0xa5, 0x65, 0x78, 0x74, 0x72, 0x61, 0x82, 0xa8, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e)
		o = msgp.AppendString(o, z[za0001].Extra.Location)
		// string "bio"
		o = append(o, 0xa3, 0x62, 0x69, 0x6f)
		o = msgp.AppendString(o, z[za0001].Extra.Bio)
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Medium) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0003 uint32
	zb0003, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		return
	}
	if cap((*z)) >= int(zb0003) {
		(*z) = (*z)[:zb0003]
	} else {
		(*z) = make(Medium, zb0003)
	}
	for zb0001 := range *z {
		var field []byte
		_ = field
		var zb0004 uint32
		zb0004, bts, err = msgp.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
		for zb0004 > 0 {
			zb0004--
			field, bts, err = msgp.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			switch string(field) {
			case "hello":
				(*z)[zb0001].Hello, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
			case "name":
				(*z)[zb0001].Name, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
			case "age":
				(*z)[zb0001].Age, bts, err = msgp.ReadIntBytes(bts)
				if err != nil {
					return
				}
			case "weight":
				(*z)[zb0001].Weight, bts, err = msgp.ReadFloat64Bytes(bts)
				if err != nil {
					return
				}
			case "height":
				(*z)[zb0001].Height, bts, err = msgp.ReadFloat64Bytes(bts)
				if err != nil {
					return
				}
			case "hobbies":
				var zb0005 uint32
				zb0005, bts, err = msgp.ReadArrayHeaderBytes(bts)
				if err != nil {
					return
				}
				if cap((*z)[zb0001].Hobbies) >= int(zb0005) {
					(*z)[zb0001].Hobbies = ((*z)[zb0001].Hobbies)[:zb0005]
				} else {
					(*z)[zb0001].Hobbies = make([]string, zb0005)
				}
				for zb0002 := range (*z)[zb0001].Hobbies {
					(*z)[zb0001].Hobbies[zb0002], bts, err = msgp.ReadStringBytes(bts)
					if err != nil {
						return
					}
				}
			case "cool":
				(*z)[zb0001].Cool, bts, err = msgp.ReadBoolBytes(bts)
				if err != nil {
					return
				}
			case "extra":
				var zb0006 uint32
				zb0006, bts, err = msgp.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				for zb0006 > 0 {
					zb0006--
					field, bts, err = msgp.ReadMapKeyZC(bts)
					if err != nil {
						return
					}
					switch string(field) {
					case "location":
						(*z)[zb0001].Extra.Location, bts, err = msgp.ReadStringBytes(bts)
						if err != nil {
							return
						}
					case "bio":
						(*z)[zb0001].Extra.Bio, bts, err = msgp.ReadStringBytes(bts)
						if err != nil {
							return
						}
					default:
						bts, err = msgp.Skip(bts)
						if err != nil {
							return
						}
					}
				}
			default:
				bts, err = msgp.Skip(bts)
				if err != nil {
					return
				}
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z Medium) Msgsize() (s int) {
	s = msgp.ArrayHeaderSize
	for zb0007 := range z {
		s += 1 + 6 + msgp.StringPrefixSize + len(z[zb0007].Hello) + 5 + msgp.StringPrefixSize + len(z[zb0007].Name) + 4 + msgp.IntSize + 7 + msgp.Float64Size + 7 + msgp.Float64Size + 8 + msgp.ArrayHeaderSize
		for zb0008 := range z[zb0007].Hobbies {
			s += msgp.StringPrefixSize + len(z[zb0007].Hobbies[zb0008])
		}
		s += 5 + msgp.BoolSize + 6 + 1 + 9 + msgp.StringPrefixSize + len(z[zb0007].Extra.Location) + 4 + msgp.StringPrefixSize + len(z[zb0007].Extra.Bio)
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Small) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 5
	// string "name"
	o = append(o, 0x85, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.Name)
	// string "age"
	o = append(o, 0xa3, 0x61, 0x67, 0x65)
	o = msgp.AppendInt(o, z.Age)
	// string "weight"
	o = append(o, 0xa6, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74)
	o = msgp.AppendFloat64(o, z.Weight)
	// string "boring"
	o = append(o, 0xa6, 0x62, 0x6f, 0x72, 0x69, 0x6e, 0x67)
	o = msgp.AppendBool(o, z.Boring)
	// string "hobbies"
	o = append(o, 0xa7, 0x68, 0x6f, 0x62, 0x62, 0x69, 0x65, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Hobbies)))
	for za0001 := range z.Hobbies {
		o = msgp.AppendString(o, z.Hobbies[za0001])
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Small) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch string(field) {
		case "name":
			z.Name, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "age":
			z.Age, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "weight":
			z.Weight, bts, err = msgp.ReadFloat64Bytes(bts)
			if err != nil {
				return
			}
		case "boring":
			z.Boring, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				return
			}
		case "hobbies":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Hobbies) >= int(zb0002) {
				z.Hobbies = (z.Hobbies)[:zb0002]
			} else {
				z.Hobbies = make([]string, zb0002)
			}
			for za0001 := range z.Hobbies {
				z.Hobbies[za0001], bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Small) Msgsize() (s int) {
	s = 1 + 5 + msgp.StringPrefixSize + len(z.Name) + 4 + msgp.IntSize + 7 + msgp.Float64Size + 7 + msgp.BoolSize + 8 + msgp.ArrayHeaderSize
	for za0001 := range z.Hobbies {
		s += msgp.StringPrefixSize + len(z.Hobbies[za0001])
	}
	return
}
