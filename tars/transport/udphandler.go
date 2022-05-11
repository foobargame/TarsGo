package transport

import (
	"context"
	"errors"
	"net"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/TarsCloud/TarsGo/tars/protocol/res/basef"
	"github.com/TarsCloud/TarsGo/tars/util/current"
	"github.com/TarsCloud/TarsGo/tars/util/grace"
)

type udpHandler struct {
	conf *TarsServerConf
	ts   *TarsServer

	conn *net.UDPConn
}

func (h *udpHandler) Listen() (err error) {
	cfg := h.conf
	h.conn, err = grace.CreateUDPConn(cfg.Address)
	if err != nil {
		return err
	}
	TLOG.Info("UDP listen", h.conn.LocalAddr())
	return nil
}

func (h *udpHandler) Handle() error {
	atomic.AddInt32(&h.ts.numConn, 1)
	// wait invoke done
	defer func() {
		tick := time.NewTicker(time.Second)
		defer tick.Stop()
		for atomic.LoadInt32(&h.ts.numInvoke) > 0 {
			<-tick.C
		}
		atomic.AddInt32(&h.ts.numConn, -1)
	}()
	buffer := make([]byte, 65535)
	for {
		if atomic.LoadInt32(&h.ts.isClosed) == 1 {
			return nil
		}
		n, udpAddr, err := h.conn.ReadFromUDP(buffer)
		if err != nil {
			if atomic.LoadInt32(&h.ts.isClosed) == 1 {
				return nil
			}
			if isNoDataError(err) {
				continue
			} else {
				TLOG.Errorf("Close connection %s: %v", h.conf.Address, err)
				return err // TODO: check if necessary
			}
		}
		pkg := make([]byte, n)
		copy(pkg, buffer[0:n])
		go func() {
			ctx := current.ContextWithTarsCurrent(context.Background())
			current.SetClientIPWithContext(ctx, udpAddr.IP.String())
			current.SetClientPortWithContext(ctx, strconv.Itoa(udpAddr.Port))
			current.SetRecvPkgTsFromContext(ctx, time.Now().UnixNano()/1e6)
			current.SetRawConnWithContext(ctx, h.conn, udpAddr)

			atomic.AddInt32(&h.ts.numInvoke, 1)
			rsp := h.ts.invoke(ctx, pkg) // no need to check package

			cPacketType, ok := current.GetPacketTypeFromContext(ctx)
			if !ok {
				TLOG.Error("Failed to GetPacketTypeFromContext")
			}

			if cPacketType == basef.TARSONEWAY {
				atomic.AddInt32(&h.ts.numInvoke, -1)
				return
			}

			if _, err := h.conn.WriteToUDP(rsp, udpAddr); err != nil {
				TLOG.Errorf("send pkg to %v failed %v", udpAddr, err)
			}
			atomic.AddInt32(&h.ts.numInvoke, -1)
		}()
	}
}

func (h *udpHandler) OnShutdown() {
}

func (h *udpHandler) CloseIdles(n int64) bool {
	if h.ts.numInvoke == 0 {
		h.conn.Close()
		return true
	}
	return false
}

func (h *udpHandler) SendData(ctx context.Context, data []byte) error {

	clientIp, ok := current.GetClientIPFromContext(ctx)
	if !ok {
		return errors.New("can't get client IP")
	}
	clientPort, ok := current.GetClientPortFromContext(ctx)
	if !ok {
		return errors.New("can't get Cient port")
	}

	addrStr := clientIp + ":" + clientPort
	udpAddr, err := net.ResolveUDPAddr("udp4", addrStr)
	if err != nil {
		return errors.New("udpaddr error:" + err.Error())
	}

	if _, err := h.conn.WriteToUDP(data, udpAddr); err != nil {
		TLOG.Errorf("send pkg to %v failed %v", udpAddr, err)
	}

	TLOG.Infof("send data to %v", udpAddr)

	return nil
}
