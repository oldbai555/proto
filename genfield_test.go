package proto

import (
	"github.com/oldbai555/proto/gen"
	"testing"
)

func TestGenProtoField(t *testing.T) {
	gen.ProtoField("lbuser")
}

func TestGenProtoModelTableName(t *testing.T) {
	gen.ModelTableName("lbuser")
}

func TestGenRpc(t *testing.T) {
	gen.RpcCode("lbuser")
}

func TestAddRpc(t *testing.T) {
	gen.AddRpc("lbuser", "UpdateUserNameWithRole")
	gen.AddRpc("lbuser", "ResetPassword")
}
