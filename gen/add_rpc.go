package gen

import (
	"fmt"
	"github.com/oldbai555/lbtool/log"
	"github.com/oldbai555/lbtool/utils"
	"github.com/oldbai555/proto/gen/temp"
	"github.com/oldbai555/proto/gen/temp/gtpl"
	"github.com/oldbai555/proto/parse"
	"os"
	"strings"
)

func AddRpc(protoFileName string, rpc string) {

	protoCtx, err := parse.ProtoCode(protoFileName)
	if err != nil {
		log.Errorf("err is %v", err)
		return
	}

	// 加入rpc 方法
	rpc = utils.UpperFirst(rpc)
	_, ok := protoCtx.RpcMap[rpc]
	if ok {
		return
	}

	var reqType = rpc + "Req"
	var rspType = rpc + "Rsp"
	template, err := temp.GenCodeByTemplate(gtpl.AddRpc, &temp.AddRpc{
		RpcName: rpc,
		ReqType: reqType,
		RspType: rspType,
	})
	if err != nil {
		log.Errorf("err is %v", err)
		return
	}
	err = protoCtx.InsertRpcBlock(template)
	if err != nil {
		log.Errorf("err is %v", err)
		return
	}

	var MsgTemp = "message %s{}"
	var lines strings.Builder
	lines.WriteString("\n\n")
	lines.WriteString(fmt.Sprintf(MsgTemp, reqType))
	lines.WriteString("\n\n")
	lines.WriteString(fmt.Sprintf(MsgTemp, rspType))

	protoCtx.Content = fmt.Sprintf("%s%s", protoCtx.Content, lines.String())

	// 检查文件名字
	if !strings.HasSuffix(parse.ProtoFileNameSuffix, protoFileName) {
		protoFileName = protoFileName + parse.ProtoFileNameSuffix
	}
	_ = os.WriteFile(protoFileName, []byte(protoCtx.Content), 0666)
}
