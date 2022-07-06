package tars

import (
    "context"
    "encoding/binary"
    "errors"
    "fmt"
    "github.com/TarsCloud/TarsGo/tars/protocol/codec"
    "strings"
    "sync/atomic"
    "time"

    "github.com/TarsCloud/TarsGo/tars/model"
    "github.com/TarsCloud/TarsGo/tars/protocol"
    "github.com/TarsCloud/TarsGo/tars/protocol/res/basef"
    "github.com/TarsCloud/TarsGo/tars/protocol/res/requestf"
    "github.com/TarsCloud/TarsGo/tars/util/current"
    "github.com/TarsCloud/TarsGo/tars/util/endpoint"
    "github.com/TarsCloud/TarsGo/tars/util/rtimer"
    "github.com/TarsCloud/TarsGo/tars/util/tools"
)

var (
    maxInt32 int32 = 1<<31 - 1
    msgID    int32
)

const (
    StatSuccess = iota
    StatFailed
)

// ServantProxy tars servant proxy instance
type ServantProxy struct {
    name     string
    comm     *Communicator
    manager  EndpointManager
    timeout  int
    version  int16
    proto    model.Protocol
    queueLen int32

    pushCallback func([]byte)
}

// NewServantProxy creates and initializes a servant proxy
func NewServantProxy(comm *Communicator, objName string) *ServantProxy {
    return newServantProxy(comm, objName)
}

func newServantProxy(comm *Communicator, objName string) *ServantProxy {
    s := &ServantProxy{}
    pos := strings.Index(objName, "@")
    if pos > 0 {
        s.name = objName[0:pos]
    } else {
        s.name = objName
    }
    pos = strings.Index(s.name, "://")
    if pos > 0 {
        s.name = s.name[pos+3:]
    }

    // init manager
    s.manager = GetManager(comm, objName)
    s.comm = comm
    s.proto = &protocol.TarsProtocol{}
    s.timeout = s.comm.Client.AsyncInvokeTimeout
    s.version = basef.TARSVERSION
    return s
}

// Name is obj name
func (s *ServantProxy) Name() string {
    return s.name
}

// TarsSetTimeout sets the timeout for client calling the server , which is in ms.
func (s *ServantProxy) TarsSetTimeout(t int) {
    s.timeout = t
}

// TarsSetVersion set tars version
func (s *ServantProxy) TarsSetVersion(iVersion int16) {
    s.version = iVersion
}

// TarsSetProtocol tars set model protocol
func (s *ServantProxy) TarsSetProtocol(proto model.Protocol) {
    s.proto = proto
}

// 生成请求 ID
func (s *ServantProxy) genRequestID() int32 {
    // 尽力防止溢出
    atomic.CompareAndSwapInt32(&msgID, maxInt32, 1)
    for {
        // 0比较特殊,用于表示 server 端推送消息给 client 端进行主动 close()
        // 溢出后回转成负数
        if v := atomic.AddInt32(&msgID, 1); v != 0 {
            return v
        }
    }
}

// SetPushCallback set callback function for pushing
func (s *ServantProxy) SetPushCallback(callback func([]byte)) {
    s.pushCallback = callback
}

// TarsInvoke is used for client invoking server.
func (s *ServantProxy) TarsInvoke(ctx context.Context, cType byte,
    sFuncName string,
    buf []byte,
    status map[string]string,
    reqContext map[string]string,
    resp *requestf.ResponsePacket) error {
    defer CheckPanic()

    // 将ctx中的dyeing信息传入到request中
    var msgType int32
    dyeingKey, ok := current.GetDyeingKey(ctx)
    if ok {
        TLOG.Debug("dyeing debug: find dyeing key:", dyeingKey)
        if status == nil {
            status = make(map[string]string)
        }
        status[current.StatusDyedKey] = dyeingKey
        msgType |= basef.TARSMESSAGETYPEDYED
    }

    // 将ctx中的trace信息传入到request中
    traceData, ok := current.GetTraceData(ctx)
    if ok && traceData.TraceCall {
        traceKey := traceData.GetTraceKeyFull(false)
        TLOG.Debug("trace debug: find trace key:", traceKey)
        if status == nil {
            status = make(map[string]string)
        }
        status[current.StatusTraceKey] = traceKey
        msgType |= basef.TARSMESSAGETYPETRACE
    }

    req := requestf.RequestPacket{
        IVersion:     s.version,
        CPacketType:  int8(cType),
        IRequestId:   s.genRequestID(),
        SServantName: s.name,
        SFuncName:    sFuncName,
        SBuffer:      tools.ByteToInt8(buf),
        // ITimeout:     s.comm.Client.ReqDefaultTimeout,
        ITimeout:     int32(s.timeout),
        Context:      reqContext,
        Status:       status,
        IMessageType: msgType,
    }
    msg := &Message{Req: &req, Ser: s, Resp: resp}
    msg.Init()

    timeout := time.Duration(s.timeout) * time.Millisecond
    ok, hashType, hashCode, isHash := current.GetClientHash(ctx)
    if ok {
        msg.isHash = isHash
        msg.hashType = HashType(hashType)
        msg.hashCode = hashCode
    }
    ok, to, isTimeout := current.GetClientTimeout(ctx)
    if ok && isTimeout {
        timeout = time.Duration(to) * time.Millisecond
    }

    var err error
    s.manager.preInvoke()
    if allFilters.cf != nil {
        err = allFilters.cf(ctx, msg, s.doInvoke, timeout)
    } else if cf := getMiddlewareClientFilter(); cf != nil {
        err = cf(ctx, msg, s.doInvoke, timeout)
    } else {
        // execute pre client filters
        for i, v := range allFilters.preCfs {
            err = v(ctx, msg, s.doInvoke, timeout)
            if err != nil {
                TLOG.Errorf("Pre filter error, no: %v, err: %v", i, err.Error())
            }
        }
        // execute rpc
        err = s.doInvoke(ctx, msg, timeout)
        // execute post client filters
        for i, v := range allFilters.postCfs {
            filterErr := v(ctx, msg, s.doInvoke, timeout)
            if filterErr != nil {
                TLOG.Errorf("Post filter error, no: %v, err: %v", i, filterErr.Error())
            }
        }
    }
    s.manager.postInvoke()

    if err != nil {
        msg.End()
        TLOG.Infof("Invoke error: %s, %s, %v, cost:%d", s.name, sFuncName, err.Error(), msg.Cost())
        if msg.Resp == nil {
            ReportStat(msg, StatSuccess, StatSuccess, StatFailed)
        } else if msg.Status == basef.TARSINVOKETIMEOUT {
            ReportStat(msg, StatSuccess, StatFailed, StatSuccess)
        } else {
            ReportStat(msg, StatSuccess, StatSuccess, StatFailed)
        }
        return err
    }
    msg.End()
    *resp = *msg.Resp
    ReportStat(msg, StatFailed, StatSuccess, StatSuccess)
    return err
}

func (s *ServantProxy) doInvoke(ctx context.Context, msg *Message, timeout time.Duration) error {
    adp, needCheck := s.manager.SelectAdapterProxy(msg)
    if adp == nil {
        return errors.New("no adapter Proxy selected:" + msg.Req.SServantName)
    }
    if s.queueLen > adp.comm.Client.ObjQueueMax {
        return errors.New("invoke queue is full:" + msg.Req.SServantName)
    }
    ep := adp.GetPoint()
    current.SetServerIPWithContext(ctx, ep.Host)
    current.SetServerPortWithContext(ctx, fmt.Sprintf("%v", ep.Port))
    msg.Adp = adp
    adp.obj = s

    if s.pushCallback != nil {
        // auto keep alive for push client
        go adp.onceKeepAlive.Do(adp.autoKeepAlive)
        adp.pushCallback = s.pushCallback
    }

    atomic.AddInt32(&s.queueLen, 1)
    readCh := make(chan *requestf.ResponsePacket)
    adp.resp.Store(msg.Req.IRequestId, readCh)
    defer func() {
        CheckPanic()
        atomic.AddInt32(&s.queueLen, -1)
        adp.resp.Delete(msg.Req.IRequestId)
    }()
    if err := adp.Send(msg.Req); err != nil {
        adp.failAdd()
        return err
    }
    if msg.Req.CPacketType == basef.TARSONEWAY {
        adp.successAdd()
        return nil
    }
    select {
    case <-rtimer.After(timeout):
        msg.Status = basef.TARSINVOKETIMEOUT
        adp.failAdd()
        msg.End()
        return fmt.Errorf("request timeout, begin time:%d, cost:%d, obj:%s, func:%s, addr:(%s:%d), reqid:%d",
            msg.BeginTime, msg.Cost(), msg.Req.SServantName, msg.Req.SFuncName, adp.point.Host, adp.point.Port, msg.Req.IRequestId)
    case msg.Resp = <-readCh:
        if needCheck {
            go func() {
                adp.reset()
                ep := endpoint.Tars2endpoint(*msg.Adp.point)
                s.manager.addAliveEp(ep)
            }()
        }
        adp.successAdd()
        if msg.Resp != nil {
            if msg.Status != basef.TARSSERVERSUCCESS || msg.Resp.IRet != 0 {
                if msg.Resp.SResultDesc == "" {
                    return fmt.Errorf("basef error code %d", msg.Resp.IRet)
                }
                if msg.Resp.IRet != 0 && msg.Resp.IRet != 1 {
                    return &Error{Code: msg.Resp.IRet, Message: msg.Resp.SResultDesc}
                }
                return errors.New(msg.Resp.SResultDesc)
            }
        } else {
            TLOG.Debug("recv nil Resp, close of the readCh?")
        }
        TLOG.Debug("recv msg success ", msg.Req.IRequestId)
    }
    return nil
}

// RequestPack is used for push
func (s *ServantProxy) RequestPack(ctx context.Context, ctype byte,
    sFuncName string,
    buf []byte,
    status map[string]string,
    reqContext map[string]string) ([]byte, error) {
    // 将ctx中的dyeinglog信息传入到request中
    var msgType int32
    dyeingKey, ok := current.GetDyeingKey(ctx)
    if ok {
        TLOG.Debug("dyeing debug: find dyeing key:", dyeingKey)
        if status == nil {
            status = make(map[string]string)
        }
        status[current.StatusDyedKey] = dyeingKey
        msgType = basef.TARSMESSAGETYPEDYED
    }

    req := requestf.RequestPacket{
        IVersion:     s.version,
        CPacketType:  int8(ctype),
        IRequestId:   s.genRequestID(),
        SServantName: s.name,
        SFuncName:    sFuncName,
        SBuffer:      tools.ByteToInt8(buf),
        //ITimeout:     s.comm.Client.ReqDefaultTimeout,
        ITimeout:     int32(s.timeout),
        Context:      reqContext,
        Status:       status,
        IMessageType: msgType,
    }
    return doRequestPack(&req)
}
func doRequestPack(req *requestf.RequestPacket) ([]byte, error) {
    os := codec.NewBuffer()
    err := os.WriteSliceInt8(make([]int8, 4))
    if err != nil {
        return nil, err
    }
    err = req.WriteTo(os)
    if err != nil {
        return nil, err
    }
    bs := os.ToBytes()
    l := len(bs)
    binary.BigEndian.PutUint32(bs, uint32(l))
    return bs, nil

}
