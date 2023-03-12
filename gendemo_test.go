package proto

import (
	"fmt"
	"github.com/oldbai555/lbtool/log"
	"github.com/oldbai555/proto/gen"
	"github.com/oldbai555/proto/gendemo"
	"github.com/oldbai555/proto/parse"
	"os"
	"path"
	"testing"
)

func TestGenDir(t *testing.T) {
	var serverName = "lbim"
	var rootDir = "../github.com/oldbai555/bgg"
	var pathClientDir = path.Join(rootDir, "client")
	gendemo.GenDir(serverName)
	gen.ProtoField(serverName, pathClientDir)
	gen.ModelTableName(serverName, pathClientDir)
	gen.ProtoError(serverName, pathClientDir)

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

	curDir, err := os.Getwd()
	if err != nil {
		log.Errorf("err is %v", err)
		return
	}
	log.Infof("cur dir is %s", curDir)
	gendemo.Exec(curDir, fmt.Sprintf("./gen.sh go %s.go", serverName))

	gendemo.GoFmt(path.Join(rootDir, "client"))
	gendemo.GoFmt(path.Join(rootDir, "server"))
	gendemo.GoFmt(path.Join(rootDir, "lbserver"))
	gendemo.GoFmt(path.Join(rootDir, "pkg"))
}
