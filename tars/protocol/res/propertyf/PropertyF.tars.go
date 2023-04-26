// Package propertyf comment
// This file was generated by tars2go 1.2.1
// Generated from PropertyF.tars
package propertyf

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	m "github.com/TarsCloud/TarsGo/tars/model"
	"github.com/TarsCloud/TarsGo/tars/protocol/codec"
	"github.com/TarsCloud/TarsGo/tars/protocol/res/basef"
	"github.com/TarsCloud/TarsGo/tars/protocol/res/requestf"
	"github.com/TarsCloud/TarsGo/tars/protocol/tup"
	"github.com/TarsCloud/TarsGo/tars/util/current"
	"github.com/TarsCloud/TarsGo/tars/util/endpoint"
	"github.com/TarsCloud/TarsGo/tars/util/tools"
	"unsafe"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = fmt.Errorf
	_ = codec.FromInt8
	_ = unsafe.Pointer(nil)
	_ = bytes.ErrTooLarge
)

// PropertyF struct
type PropertyF struct {
	servant m.Servant
}

// SetServant sets servant for the service.
func (obj *PropertyF) SetServant(servant m.Servant) {
	obj.servant = servant
}

// TarsSetTimeout sets the timeout for the servant which is in ms.
func (obj *PropertyF) TarsSetTimeout(timeout int) {
	obj.servant.TarsSetTimeout(timeout)
}

// TarsSetProtocol sets the protocol for the servant.
func (obj *PropertyF) TarsSetProtocol(p m.Protocol) {
	obj.servant.TarsSetProtocol(p)
}

// Endpoints returns all active endpoint.Endpoint
func (obj *PropertyF) Endpoints() []*endpoint.Endpoint {
	return obj.servant.Endpoints()
}

// ReportPropMsg is the proxy function for the method defined in the tars file, with the context
func (obj *PropertyF) ReportPropMsg(statmsg map[StatPropMsgHead]StatPropMsgBody, opts ...map[string]string) (int32, error) {
	return obj.ReportPropMsgWithContext(context.Background(), statmsg, opts...)
}

// ReportPropMsgWithContext is the proxy function for the method defined in the tars file, with the context
func (obj *PropertyF) ReportPropMsgWithContext(tarsCtx context.Context, statmsg map[StatPropMsgHead]StatPropMsgBody, opts ...map[string]string) (ret int32, err error) {
	var (
		length int32
		have   bool
		ty     byte
	)
	buf := codec.NewBuffer()
	err = buf.WriteHead(codec.MAP, 1)
	if err != nil {
		return ret, err
	}

	err = buf.WriteInt32(int32(len(statmsg)), 0)
	if err != nil {
		return ret, err
	}

	for k1, v1 := range statmsg {

		err = k1.WriteBlock(buf, 0)
		if err != nil {
			return ret, err
		}

		err = v1.WriteBlock(buf, 1)
		if err != nil {
			return ret, err
		}

	}

	var statusMap map[string]string
	var contextMap map[string]string
	if len(opts) == 1 {
		contextMap = opts[0]
	} else if len(opts) == 2 {
		contextMap = opts[0]
		statusMap = opts[1]
	}

	tarsResp := new(requestf.ResponsePacket)
	err = obj.servant.TarsInvoke(tarsCtx, 0, "reportPropMsg", buf.ToBytes(), statusMap, contextMap, tarsResp)
	if err != nil {
		return ret, err
	}

	readBuf := codec.NewReader(tools.Int8ToByte(tarsResp.SBuffer))
	err = readBuf.ReadInt32(&ret, 0, true)
	if err != nil {
		return ret, err
	}

	if len(opts) == 1 {
		for k := range contextMap {
			delete(contextMap, k)
		}
		for k, v := range tarsResp.Context {
			contextMap[k] = v
		}
	} else if len(opts) == 2 {
		for k := range contextMap {
			delete(contextMap, k)
		}
		for k, v := range tarsResp.Context {
			contextMap[k] = v
		}
		for k := range statusMap {
			delete(statusMap, k)
		}
		for k, v := range tarsResp.Status {
			statusMap[k] = v
		}
	}
	_ = length
	_ = have
	_ = ty
	return ret, nil
}

// ReportPropMsgOneWayWithContext is the proxy function for the method defined in the tars file, with the context
func (obj *PropertyF) ReportPropMsgOneWayWithContext(tarsCtx context.Context, statmsg map[StatPropMsgHead]StatPropMsgBody, opts ...map[string]string) (ret int32, err error) {
	var (
		length int32
		have   bool
		ty     byte
	)
	buf := codec.NewBuffer()
	err = buf.WriteHead(codec.MAP, 1)
	if err != nil {
		return ret, err
	}

	err = buf.WriteInt32(int32(len(statmsg)), 0)
	if err != nil {
		return ret, err
	}

	for k2, v2 := range statmsg {

		err = k2.WriteBlock(buf, 0)
		if err != nil {
			return ret, err
		}

		err = v2.WriteBlock(buf, 1)
		if err != nil {
			return ret, err
		}

	}

	var statusMap map[string]string
	var contextMap map[string]string
	if len(opts) == 1 {
		contextMap = opts[0]
	} else if len(opts) == 2 {
		contextMap = opts[0]
		statusMap = opts[1]
	}

	tarsResp := new(requestf.ResponsePacket)
	err = obj.servant.TarsInvoke(tarsCtx, 1, "reportPropMsg", buf.ToBytes(), statusMap, contextMap, tarsResp)
	if err != nil {
		return ret, err
	}

	_ = length
	_ = have
	_ = ty
	return ret, nil
}

type PropertyFServant interface {
	ReportPropMsg(statmsg map[StatPropMsgHead]StatPropMsgBody) (ret int32, err error)
}
type PropertyFServantWithContext interface {
	ReportPropMsg(tarsCtx context.Context, statmsg map[StatPropMsgHead]StatPropMsgBody) (ret int32, err error)
}

// Dispatch is used to call the server side implement for the method defined in the tars file. withContext shows using context or not.
func (obj *PropertyF) Dispatch(tarsCtx context.Context, val interface{}, tarsReq *requestf.RequestPacket, tarsResp *requestf.ResponsePacket, withContext bool) (err error) {
	var (
		length int32
		have   bool
		ty     byte
	)
	readBuf := codec.NewReader(tools.Int8ToByte(tarsReq.SBuffer))
	buf := codec.NewBuffer()
	switch tarsReq.SFuncName {
	case "reportPropMsg":
		var statmsg map[StatPropMsgHead]StatPropMsgBody
		statmsg = make(map[StatPropMsgHead]StatPropMsgBody)

		if tarsReq.IVersion == basef.TARSVERSION {

			_, err = readBuf.SkipTo(codec.MAP, 1, true)
			if err != nil {
				return err
			}

			err = readBuf.ReadInt32(&length, 0, true)
			if err != nil {
				return err
			}

			statmsg = make(map[StatPropMsgHead]StatPropMsgBody)
			for i3, e3 := int32(0), length; i3 < e3; i3++ {
				var k3 StatPropMsgHead
				var v3 StatPropMsgBody

				err = k3.ReadBlock(readBuf, 0, false)
				if err != nil {
					return err
				}

				err = v3.ReadBlock(readBuf, 1, false)
				if err != nil {
					return err
				}

				statmsg[k3] = v3
			}
		} else if tarsReq.IVersion == basef.TUPVERSION {
			reqTup := tup.NewUniAttribute()
			reqTup.Decode(readBuf)

			var tupBuffer []byte

			reqTup.GetBuffer("statmsg", &tupBuffer)
			readBuf.Reset(tupBuffer)
			_, err = readBuf.SkipTo(codec.MAP, 0, true)
			if err != nil {
				return err
			}

			err = readBuf.ReadInt32(&length, 0, true)
			if err != nil {
				return err
			}

			statmsg = make(map[StatPropMsgHead]StatPropMsgBody)
			for i4, e4 := int32(0), length; i4 < e4; i4++ {
				var k4 StatPropMsgHead
				var v4 StatPropMsgBody

				err = k4.ReadBlock(readBuf, 0, false)
				if err != nil {
					return err
				}

				err = v4.ReadBlock(readBuf, 1, false)
				if err != nil {
					return err
				}

				statmsg[k4] = v4
			}
		} else if tarsReq.IVersion == basef.JSONVERSION {
			var jsonData map[string]interface{}
			decoder := json.NewDecoder(bytes.NewReader(readBuf.ToBytes()))
			decoder.UseNumber()
			err = decoder.Decode(&jsonData)
			if err != nil {
				return fmt.Errorf("decode reqpacket failed, error: %+v", err)
			}
			{
				jsonStr, _ := json.Marshal(jsonData["statmsg"])
				if err = json.Unmarshal(jsonStr, &statmsg); err != nil {
					return err
				}
			}

		} else {
			err = fmt.Errorf("decode reqpacket fail, error version: %d", tarsReq.IVersion)
			return err
		}

		var funRet int32
		if !withContext {
			imp := val.(PropertyFServant)
			funRet, err = imp.ReportPropMsg(statmsg)
		} else {
			imp := val.(PropertyFServantWithContext)
			funRet, err = imp.ReportPropMsg(tarsCtx, statmsg)
		}

		if err != nil {
			return err
		}

		if tarsReq.IVersion == basef.TARSVERSION {
			buf.Reset()

			err = buf.WriteInt32(funRet, 0)
			if err != nil {
				return err
			}

		} else if tarsReq.IVersion == basef.TUPVERSION {
			rspTup := tup.NewUniAttribute()

			err = buf.WriteInt32(funRet, 0)
			if err != nil {
				return err
			}

			rspTup.PutBuffer("", buf.ToBytes())
			rspTup.PutBuffer("tars_ret", buf.ToBytes())

			buf.Reset()
			err = rspTup.Encode(buf)
			if err != nil {
				return err
			}
		} else if tarsReq.IVersion == basef.JSONVERSION {
			rspJson := map[string]interface{}{}
			rspJson["tars_ret"] = funRet

			var rspByte []byte
			if rspByte, err = json.Marshal(rspJson); err != nil {
				return err
			}

			buf.Reset()
			err = buf.WriteSliceUint8(rspByte)
			if err != nil {
				return err
			}
		}

	default:
		return fmt.Errorf("func mismatch")
	}
	var statusMap map[string]string
	if status, ok := current.GetResponseStatus(tarsCtx); ok && status != nil {
		statusMap = status
	}
	var contextMap map[string]string
	if ctx, ok := current.GetResponseContext(tarsCtx); ok && ctx != nil {
		contextMap = ctx
	}
	*tarsResp = requestf.ResponsePacket{
		IVersion:     tarsReq.IVersion,
		CPacketType:  0,
		IRequestId:   tarsReq.IRequestId,
		IMessageType: 0,
		IRet:         0,
		SBuffer:      tools.ByteToInt8(buf.ToBytes()),
		Status:       statusMap,
		SResultDesc:  "",
		Context:      contextMap,
	}

	_ = readBuf
	_ = buf
	_ = length
	_ = have
	_ = ty
	return nil
}
