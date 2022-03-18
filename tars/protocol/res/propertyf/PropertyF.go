// Package propertyf comment
// This file was generated by tars2go 1.1.6
// Generated from PropertyF.tars
package propertyf

import (
	"fmt"

	"github.com/TarsCloud/TarsGo/tars/protocol/codec"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = fmt.Errorf
var _ = codec.FromInt8

// StatPropMsgHead struct implement
type StatPropMsgHead struct {
	ModuleName   string `json:"moduleName"`
	Ip           string `json:"ip"`
	PropertyName string `json:"propertyName"`
	SetName      string `json:"setName"`
	SetArea      string `json:"setArea"`
	SetID        string `json:"setID"`
	SContainer   string `json:"sContainer"`
	IPropertyVer int32  `json:"iPropertyVer"`
}

func (st *StatPropMsgHead) ResetDefault() {
	st.IPropertyVer = 1
}

// ReadFrom reads  from readBuf and put into struct.
func (st *StatPropMsgHead) ReadFrom(readBuf *codec.Reader) error {
	var (
		err    error
		length int32
		have   bool
		ty     byte
	)
	st.ResetDefault()

	err = readBuf.ReadString(&st.ModuleName, 0, true)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.Ip, 1, true)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.PropertyName, 2, true)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.SetName, 3, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.SetArea, 4, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.SetID, 5, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.SContainer, 6, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadInt32(&st.IPropertyVer, 7, false)
	if err != nil {
		return err
	}

	_ = err
	_ = length
	_ = have
	_ = ty
	return nil
}

// ReadBlock reads struct from the given tag , require or optional.
func (st *StatPropMsgHead) ReadBlock(readBuf *codec.Reader, tag byte, require bool) error {
	var (
		err  error
		have bool
	)
	st.ResetDefault()

	have, err = readBuf.SkipTo(codec.StructBegin, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require StatPropMsgHead, but not exist. tag %d", tag)
		}
		return nil
	}

	err = st.ReadFrom(readBuf)
	if err != nil {
		return err
	}

	err = readBuf.SkipToStructEnd()
	if err != nil {
		return err
	}
	_ = have
	return nil
}

// WriteTo encode struct to buffer
func (st *StatPropMsgHead) WriteTo(buf *codec.Buffer) (err error) {

	err = buf.WriteString(st.ModuleName, 0)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.Ip, 1)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.PropertyName, 2)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.SetName, 3)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.SetArea, 4)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.SetID, 5)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.SContainer, 6)
	if err != nil {
		return err
	}

	err = buf.WriteInt32(st.IPropertyVer, 7)
	if err != nil {
		return err
	}

	return err
}

// WriteBlock encode struct
func (st *StatPropMsgHead) WriteBlock(buf *codec.Buffer, tag byte) error {
	var err error
	err = buf.WriteHead(codec.StructBegin, tag)
	if err != nil {
		return err
	}

	err = st.WriteTo(buf)
	if err != nil {
		return err
	}

	err = buf.WriteHead(codec.StructEnd, 0)
	if err != nil {
		return err
	}
	return nil
}

// StatPropInfo struct implement
type StatPropInfo struct {
	Policy string `json:"policy"`
	Value  string `json:"value"`
}

func (st *StatPropInfo) ResetDefault() {
}

// ReadFrom reads  from readBuf and put into struct.
func (st *StatPropInfo) ReadFrom(readBuf *codec.Reader) error {
	var (
		err    error
		length int32
		have   bool
		ty     byte
	)
	st.ResetDefault()

	err = readBuf.ReadString(&st.Policy, 0, true)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.Value, 1, true)
	if err != nil {
		return err
	}

	_ = err
	_ = length
	_ = have
	_ = ty
	return nil
}

// ReadBlock reads struct from the given tag , require or optional.
func (st *StatPropInfo) ReadBlock(readBuf *codec.Reader, tag byte, require bool) error {
	var (
		err  error
		have bool
	)
	st.ResetDefault()

	have, err = readBuf.SkipTo(codec.StructBegin, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require StatPropInfo, but not exist. tag %d", tag)
		}
		return nil
	}

	err = st.ReadFrom(readBuf)
	if err != nil {
		return err
	}

	err = readBuf.SkipToStructEnd()
	if err != nil {
		return err
	}
	_ = have
	return nil
}

// WriteTo encode struct to buffer
func (st *StatPropInfo) WriteTo(buf *codec.Buffer) (err error) {

	err = buf.WriteString(st.Policy, 0)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.Value, 1)
	if err != nil {
		return err
	}

	return err
}

// WriteBlock encode struct
func (st *StatPropInfo) WriteBlock(buf *codec.Buffer, tag byte) error {
	var err error
	err = buf.WriteHead(codec.StructBegin, tag)
	if err != nil {
		return err
	}

	err = st.WriteTo(buf)
	if err != nil {
		return err
	}

	err = buf.WriteHead(codec.StructEnd, 0)
	if err != nil {
		return err
	}
	return nil
}

// StatPropMsgBody struct implement
type StatPropMsgBody struct {
	VInfo []StatPropInfo `json:"vInfo"`
}

func (st *StatPropMsgBody) ResetDefault() {
}

// ReadFrom reads  from readBuf and put into struct.
func (st *StatPropMsgBody) ReadFrom(readBuf *codec.Reader) error {
	var (
		err    error
		length int32
		have   bool
		ty     byte
	)
	st.ResetDefault()

	_, ty, err = readBuf.SkipToNoCheck(0, true)
	if err != nil {
		return err
	}

	if ty == codec.LIST {
		err = readBuf.ReadInt32(&length, 0, true)
		if err != nil {
			return err
		}

		st.VInfo = make([]StatPropInfo, length)
		for i0, e0 := int32(0), length; i0 < e0; i0++ {

			err = st.VInfo[i0].ReadBlock(readBuf, 0, false)
			if err != nil {
				return err
			}

		}
	} else if ty == codec.SimpleList {
		err = fmt.Errorf("not support SimpleList type")
		if err != nil {
			return err
		}

	} else {
		err = fmt.Errorf("require vector, but not")
		if err != nil {
			return err
		}

	}

	_ = err
	_ = length
	_ = have
	_ = ty
	return nil
}

// ReadBlock reads struct from the given tag , require or optional.
func (st *StatPropMsgBody) ReadBlock(readBuf *codec.Reader, tag byte, require bool) error {
	var (
		err  error
		have bool
	)
	st.ResetDefault()

	have, err = readBuf.SkipTo(codec.StructBegin, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require StatPropMsgBody, but not exist. tag %d", tag)
		}
		return nil
	}

	err = st.ReadFrom(readBuf)
	if err != nil {
		return err
	}

	err = readBuf.SkipToStructEnd()
	if err != nil {
		return err
	}
	_ = have
	return nil
}

// WriteTo encode struct to buffer
func (st *StatPropMsgBody) WriteTo(buf *codec.Buffer) (err error) {

	err = buf.WriteHead(codec.LIST, 0)
	if err != nil {
		return err
	}

	err = buf.WriteInt32(int32(len(st.VInfo)), 0)
	if err != nil {
		return err
	}

	for _, v := range st.VInfo {

		err = v.WriteBlock(buf, 0)
		if err != nil {
			return err
		}

	}

	return err
}

// WriteBlock encode struct
func (st *StatPropMsgBody) WriteBlock(buf *codec.Buffer, tag byte) error {
	var err error
	err = buf.WriteHead(codec.StructBegin, tag)
	if err != nil {
		return err
	}

	err = st.WriteTo(buf)
	if err != nil {
		return err
	}

	err = buf.WriteHead(codec.StructEnd, 0)
	if err != nil {
		return err
	}
	return nil
}