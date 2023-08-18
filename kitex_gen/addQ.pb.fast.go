// Code generated by Fastpb v0.0.2. DO NOT EDIT.

package kitex_gen

import (
	fmt "fmt"
	fastpb "github.com/cloudwego/fastpb"
)

var (
	_ = fmt.Errorf
	_ = fastpb.Skip
)

func (x *AddRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_AddRequest[number], err)
}

func (x *AddRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.X, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *AddRequest) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Y, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *AddResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_AddResponse[number], err)
}

func (x *AddResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Res, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *AddRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *AddRequest) fastWriteField1(buf []byte) (offset int) {
	if x.X == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 1, x.GetX())
	return offset
}

func (x *AddRequest) fastWriteField2(buf []byte) (offset int) {
	if x.Y == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 2, x.GetY())
	return offset
}

func (x *AddResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *AddResponse) fastWriteField1(buf []byte) (offset int) {
	if x.Res == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 1, x.GetRes())
	return offset
}

func (x *AddRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *AddRequest) sizeField1() (n int) {
	if x.X == 0 {
		return n
	}
	n += fastpb.SizeInt32(1, x.GetX())
	return n
}

func (x *AddRequest) sizeField2() (n int) {
	if x.Y == 0 {
		return n
	}
	n += fastpb.SizeInt32(2, x.GetY())
	return n
}

func (x *AddResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *AddResponse) sizeField1() (n int) {
	if x.Res == 0 {
		return n
	}
	n += fastpb.SizeInt32(1, x.GetRes())
	return n
}

var fieldIDToName_AddRequest = map[int32]string{
	1: "X",
	2: "Y",
}

var fieldIDToName_AddResponse = map[int32]string{
	1: "Res",
}
