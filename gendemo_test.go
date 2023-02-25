package proto

import (
	"fmt"
	"github.com/oldbai555/lbtool/log"
	"github.com/oldbai555/proto/gen"
	"github.com/oldbai555/proto/gendemo"
	"github.com/oldbai555/proto/parse"
	"path"
	"testing"
)

func TestGenDir(t *testing.T) {
	var serverName = "lbwebsocket"
	var pathClientDir = "../github.com/oldbai555/bgg/client"
	gendemo.GenDir(serverName)
	gen.ProtoField(serverName, pathClientDir)
	gen.ModelTableName(serverName, pathClientDir)

	var rpcCodeDir = path.Join("../github.com/oldbai555/bgg/server/", serverName, "impl")

	// 生成服务端实现类代码
	protoCtx, err := parse.ProtoCode(serverName)
	if err != nil {
		log.Errorf("err is %v", err)
		return
	}
	code, err := parse.ParserGoCode(rpcCodeDir, fmt.Sprintf("%s_server.go", parse.TrimProtoFileNameSuffix(parse.TrimProtoFileNamePrefixWithLb(serverName))))
	if err != nil {
		log.Errorf("err is %v", err)
		return
	}
	err = gen.GenServerFuncCode(serverName, code, protoCtx)
	if err != nil {
		log.Errorf("err is %v", err)
		return
	}
}
