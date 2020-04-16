package replicator

import (
	"io"
	"net"
	"net/rpc"
	"reflect"
	"strings"

	hcodec "github.com/hashicorp/go-msgpack/codec"
	msgpackrpc "github.com/hashicorp/net-rpc-msgpackrpc"
)

// HashiMsgpackHandle is some magic.
var HashiMsgpackHandle = func() *hcodec.MsgpackHandle {
	h := &hcodec.MsgpackHandle{}
	h.MapType = reflect.TypeOf(map[string]interface{}(nil))
	return h
}()

// NewServerCodec returns a new rpc.ServerCodec to be used by the Replicator
// Server to process RPC requests.
func NewServerCodec(conn io.ReadWriteCloser) rpc.ServerCodec {
	return msgpackrpc.NewCodecFromHandle(true, true, conn, HashiMsgpackHandle)
}

func (s *Server) listen() {
	for {
		conn, err := s.rpcListener.Accept()
		if err != nil {
			if s.shutdown {
				return
			}
			continue
		}
		go s.handleConn(conn)
	}
}

func (s *Server) handleConn(conn net.Conn) {
	defer conn.Close()
	rpcCodec := NewServerCodec(conn)
	for {
		select {
		case <-s.shutdownChan:
			return
		default:
		}

		if err := s.rpcServer.ServeRequest(rpcCodec); err != nil {
			if err != io.EOF && !strings.Contains(err.Error(), "closed") {
			}
			return
		}
	}
}
