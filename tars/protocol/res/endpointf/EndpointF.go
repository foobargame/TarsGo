// Package endpointf comment
// This file was generated by tars2go 1.1.4
// Generated from EndpointF.tars
package endpointf

import (
	"fmt"

	"github.com/TarsCloud/TarsGo/tars/protocol/codec"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = fmt.Errorf
var _ = codec.FromInt8

// EndpointF struct implement
type EndpointF struct {
	Host        string `json:"host"`
	Port        int32  `json:"port"`
	Timeout     int32  `json:"timeout"`
	Istcp       int32  `json:"istcp"`
	Grid        int32  `json:"grid"`
	Groupworkid int32  `json:"groupworkid"`
	Grouprealid int32  `json:"grouprealid"`
	SetId       string `json:"setId"`
	Qos         int32  `json:"qos"`
	BakFlag     int32  `json:"bakFlag"`
	Weight      int32  `json:"weight"`
	WeightType  int32  `json:"weightType"`
	AuthType    int32  `json:"authType"`
	Subset      string `json:"subset"`
}

func (st *EndpointF) ResetDefault() {
}

//ReadFrom reads  from _is and put into struct.
func (st *EndpointF) ReadFrom(_is *codec.Reader) error {
	var err error
	var length int32
	var have bool
	var ty byte
	st.ResetDefault()

	err = _is.Read_string(&st.Host, 0, true)
	if err != nil {
		return err
	}

	err = _is.Read_int32(&st.Port, 1, true)
	if err != nil {
		return err
	}

	err = _is.Read_int32(&st.Timeout, 2, true)
	if err != nil {
		return err
	}

	err = _is.Read_int32(&st.Istcp, 3, true)
	if err != nil {
		return err
	}

	err = _is.Read_int32(&st.Grid, 4, true)
	if err != nil {
		return err
	}

	err = _is.Read_int32(&st.Groupworkid, 5, false)
	if err != nil {
		return err
	}

	err = _is.Read_int32(&st.Grouprealid, 6, false)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.SetId, 7, false)
	if err != nil {
		return err
	}

	err = _is.Read_int32(&st.Qos, 8, false)
	if err != nil {
		return err
	}

	err = _is.Read_int32(&st.BakFlag, 9, false)
	if err != nil {
		return err
	}

	err = _is.Read_int32(&st.Weight, 11, false)
	if err != nil {
		return err
	}

	err = _is.Read_int32(&st.WeightType, 12, false)
	if err != nil {
		return err
	}

	err = _is.Read_int32(&st.AuthType, 13, false)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.Subset, 14, false)
	if err != nil {
		return err
	}

	_ = err
	_ = length
	_ = have
	_ = ty
	return nil
}

//ReadBlock reads struct from the given tag , require or optional.
func (st *EndpointF) ReadBlock(_is *codec.Reader, tag byte, require bool) error {
	var err error
	var have bool
	st.ResetDefault()

	err, have = _is.SkipTo(codec.STRUCT_BEGIN, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require EndpointF, but not exist. tag %d", tag)
		}
		return nil
	}

	err = st.ReadFrom(_is)
	if err != nil {
		return err
	}

	err = _is.SkipToStructEnd()
	if err != nil {
		return err
	}
	_ = have
	return nil
}

//WriteTo encode struct to buffer
func (st *EndpointF) WriteTo(_os *codec.Buffer) error {
	var err error

	err = _os.Write_string(st.Host, 0)
	if err != nil {
		return err
	}

	err = _os.Write_int32(st.Port, 1)
	if err != nil {
		return err
	}

	err = _os.Write_int32(st.Timeout, 2)
	if err != nil {
		return err
	}

	err = _os.Write_int32(st.Istcp, 3)
	if err != nil {
		return err
	}

	err = _os.Write_int32(st.Grid, 4)
	if err != nil {
		return err
	}

	err = _os.Write_int32(st.Groupworkid, 5)
	if err != nil {
		return err
	}

	err = _os.Write_int32(st.Grouprealid, 6)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.SetId, 7)
	if err != nil {
		return err
	}

	err = _os.Write_int32(st.Qos, 8)
	if err != nil {
		return err
	}

	err = _os.Write_int32(st.BakFlag, 9)
	if err != nil {
		return err
	}

	err = _os.Write_int32(st.Weight, 11)
	if err != nil {
		return err
	}

	err = _os.Write_int32(st.WeightType, 12)
	if err != nil {
		return err
	}

	err = _os.Write_int32(st.AuthType, 13)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.Subset, 14)
	if err != nil {
		return err
	}

	_ = err

	return nil
}

//WriteBlock encode struct
func (st *EndpointF) WriteBlock(_os *codec.Buffer, tag byte) error {
	var err error
	err = _os.WriteHead(codec.STRUCT_BEGIN, tag)
	if err != nil {
		return err
	}

	err = st.WriteTo(_os)
	if err != nil {
		return err
	}

	err = _os.WriteHead(codec.STRUCT_END, 0)
	if err != nil {
		return err
	}
	return nil
}

// SubsetConf struct implement
type SubsetConf struct {
	Enable   bool   `json:"enable"`
	RuleType string `json:"ruleType"`
	RuteData string `json:"ruteData"`
}

func (st *SubsetConf) ResetDefault() {
}

//ReadFrom reads  from _is and put into struct.
func (st *SubsetConf) ReadFrom(_is *codec.Reader) error {
	var err error
	var length int32
	var have bool
	var ty byte
	st.ResetDefault()

	err = _is.Read_bool(&st.Enable, 0, true)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.RuleType, 1, true)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.RuteData, 2, true)
	if err != nil {
		return err
	}

	_ = err
	_ = length
	_ = have
	_ = ty
	return nil
}

//ReadBlock reads struct from the given tag , require or optional.
func (st *SubsetConf) ReadBlock(_is *codec.Reader, tag byte, require bool) error {
	var err error
	var have bool
	st.ResetDefault()

	err, have = _is.SkipTo(codec.STRUCT_BEGIN, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require SubsetConf, but not exist. tag %d", tag)
		}
		return nil
	}

	err = st.ReadFrom(_is)
	if err != nil {
		return err
	}

	err = _is.SkipToStructEnd()
	if err != nil {
		return err
	}
	_ = have
	return nil
}

//WriteTo encode struct to buffer
func (st *SubsetConf) WriteTo(_os *codec.Buffer) error {
	var err error

	err = _os.Write_bool(st.Enable, 0)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.RuleType, 1)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.RuteData, 2)
	if err != nil {
		return err
	}

	_ = err

	return nil
}

//WriteBlock encode struct
func (st *SubsetConf) WriteBlock(_os *codec.Buffer, tag byte) error {
	var err error
	err = _os.WriteHead(codec.STRUCT_BEGIN, tag)
	if err != nil {
		return err
	}

	err = st.WriteTo(_os)
	if err != nil {
		return err
	}

	err = _os.WriteHead(codec.STRUCT_END, 0)
	if err != nil {
		return err
	}
	return nil
}
