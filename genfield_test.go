package proto

import (
	"github.com/oldbai555/proto/gen"
	"testing"
)

func TestGenProtoField(t *testing.T) {
	gen.ProtoField("lbwebsocket", "../github.com/oldbai555/bgg")
}

func TestGenProtoModelTableName(t *testing.T) {
	gen.ModelTableName("lbwebsocket", "../github.com/oldbai555/bgg")
}

func TestGenRpc(t *testing.T) {
	gen.RpcCode("lbwebsocket")
}

func TestAddRpc(t *testing.T) {
	gen.AddRpc("lbwebsocket", "HandleWs")
}

func TestInitProto(t *testing.T) {
	gen.InitProto("lbwebsocket")
}

func TestGenError(t *testing.T) {
	gen.ProtoError("lbuser", "../github.com/oldbai555/bgg/client")
}
