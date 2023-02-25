package gen

import (
	"fmt"
	"github.com/oldbai555/lbtool/log"
	"github.com/oldbai555/proto/parse"
)

const rpcCodeDir = "../github.com/oldbai555/bgg/lbserver/impl"

func RpcCode(protoFileName string) {
	// step 1: 解析proto文件
	protoCtx, err := parse.ProtoCode(protoFileName)
	if err != nil {
		log.Errorf("err is %v", err)
		return
	}

	// step 2: 开始生成代码
	// step 2.1: 生成server 代码
	code, err := parse.ParserGoCode(rpcCodeDir, fmt.Sprintf("%s_server.go", parse.TrimProtoFileNameSuffix(parse.TrimProtoFileNamePrefixWithLb(protoFileName))))
	if err != nil {
		log.Errorf("err is %v", err)
		return
	}
	err = GenServerFuncCode(protoFileName, code, protoCtx)
	if err != nil {
		log.Errorf("err is %v", err)
		return
	}

	// step 2.2: 生成 api 端代码
	code, err = parse.ParserGoCode(rpcCodeDir, fmt.Sprintf("%s_api.go", parse.TrimProtoFileNameSuffix(parse.TrimProtoFileNamePrefixWithLb(protoFileName))))
	if err != nil {
		log.Errorf("err is %v", err)
		return
	}
	err = genApiCode(protoFileName, code, protoCtx)
	if err != nil {
		log.Errorf("err is %v", err)
		return
	}

}
