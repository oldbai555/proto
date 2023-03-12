package proto

import (
	"fmt"
	"github.com/oldbai555/lbtool/log"
	"github.com/oldbai555/proto/gen"
	"github.com/oldbai555/proto/gendemo"
	"os"
	"testing"
)

func TestGenProtoField(t *testing.T) {
	gen.ProtoField("lbaccount", "../github.com/oldbai555/bgg/client")
	gen.ProtoField("lbcustomer", "../github.com/oldbai555/bgg/client")
	gen.ProtoField("lbim", "../github.com/oldbai555/bgg/client")
}

func TestGenProtoModelTableName(t *testing.T) {
	gen.ModelTableName("lbaccount", "../github.com/oldbai555/bgg/client")
	gen.ModelTableName("lbcustomer", "../github.com/oldbai555/bgg/client")
	gen.ModelTableName("lbim", "../github.com/oldbai555/bgg/client")
}

func TestGenRpc(t *testing.T) {
	gen.RpcCode("lbwebsocket")
}

func TestAddRpc(t *testing.T) {
	gen.AddRpc("lbwebsocket", "HandleWs")
}

func TestInitProto(t *testing.T) {
	gen.InitProto("lbaccount")
}

func TestGenError(t *testing.T) {
	gen.ProtoError("lbaccount", "../github.com/oldbai555/bgg/client")
	gen.ProtoError("lbcustomer", "../github.com/oldbai555/bgg/client")
	gen.ProtoError("lbim", "../github.com/oldbai555/bgg/client")
}

func TestGenPro(t *testing.T) {
	curDir, err := os.Getwd()
	if err != nil {
		log.Errorf("err is %v", err)
		return
	}
	log.Infof("cur dir is %s", curDir)
	gendemo.Exec(curDir, fmt.Sprintf("./gen.sh go %s.proto", "lbaccount"))
	gendemo.Exec(curDir, fmt.Sprintf("./gen.sh go %s.proto", "lbcustomer"))
	gendemo.Exec(curDir, fmt.Sprintf("./gen.sh go %s.proto", "lbim"))
}
