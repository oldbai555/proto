package proto

import (
	"github.com/oldbai555/proto/gen"
	"testing"
)

func TestGenProtoField(t *testing.T) {
	gen.ProtoField("lbblog")
	gen.ProtoField("lbuser")
	gen.ProtoField("lbconst")
}

func TestGenProtoModelTableName(t *testing.T) {
	gen.ModelTableName("lbblog")
	gen.ModelTableName("lbuser")
	gen.ModelTableName("lbconst")
}

func TestGenRpc(t *testing.T) {
	gen.RpcCode("lbblog")
}

func TestAddRpc(t *testing.T) {
	gen.AddRpc("lbblog", "TestHello1")
}
